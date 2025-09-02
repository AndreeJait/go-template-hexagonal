package auth

import (
	"context"
	"errors"
	"github.com/AndreeJait/go-template-hexagonal/internal/adapters/outbound/email"
	"github.com/AndreeJait/go-template-hexagonal/internal/constant"
	"github.com/AndreeJait/go-template-hexagonal/internal/domain"
	"github.com/AndreeJait/go-template-hexagonal/internal/infrastructure/config"
	"github.com/AndreeJait/go-template-hexagonal/internal/infrastructure/db"
	"github.com/AndreeJait/go-template-hexagonal/internal/utils"

	"github.com/AndreeJait/go-utility/errow"
	"github.com/AndreeJait/go-utility/jwt"
	"github.com/AndreeJait/go-utility/loggerw"
	"github.com/AndreeJait/go-utility/password"
	"github.com/AndreeJait/go-utility/tracer"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v5"
	"time"
)

type useCase struct {
	userRepo       UserRepo
	roleRepo       RoleRepo
	permissionRepo PermissionRepo
	txManager      *db.TxManager
	email          email.Email
	log            loggerw.Logger
	cfg            *config.Config
}

func (u *useCase) CreatePin(ctx context.Context, req CreateUserPinRequest) error {
	span, ctx := tracer.StartSpan(ctx, tracer.GetFuncName(u.CreatePin))
	defer span.End()
	user, err := u.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return err
	}
	timeNow := utils.TimeNowJkrt()
	if user.ActivationExpiredAt.Before(timeNow) {
		return errow.ErrInvalidToken
	}
	isValid := password.ComparePasswords(user.ActivationToken, []byte(req.Token))
	if !isValid {
		return errow.ErrInvalidToken
	}
	err = u.txManager.Run(ctx, func(tx pgx.Tx) error {
		pin, err := password.HashAndSalt([]byte(req.Pin))
		if err != nil {
			return err
		}
		err = u.userRepo.UpdateUserPasswordPinTokenActivationStatusByUserIDParams(ctx, domain.User{
			ID:                  user.ID,
			ActivationToken:     "_",
			Pin:                 pin,
			ActivationExpiredAt: utils.TimeNowJkrt().Add(-10 * time.Minute),
		}, tx)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

const TokenLength = 16

func (u *useCase) CreatePassword(ctx context.Context, req CreateUserPasswordRequest) (resp CreateUserPasswordResponse, err error) {
	span, ctx := tracer.StartSpan(ctx, tracer.GetFuncName(u.CreatePassword))
	defer span.End()

	user, err := u.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return resp, err
	}

	timeNow := utils.TimeNowJkrt()

	if user.ActivationExpiredAt.Before(timeNow) {
		return resp, errow.ErrInvalidToken
	}

	isValid := password.ComparePasswords(user.ActivationToken, []byte(req.Token))
	if !isValid {
		return resp, errow.ErrInvalidToken
	}
	// new token
	token, err := utils.GenerateRandomToken(TokenLength)
	if err != nil {
		return
	}
	tokenHash, err := password.HashAndSalt([]byte(token))
	if err != nil {
		return
	}

	err = u.txManager.Run(ctx, func(tx pgx.Tx) error {
		pass, err := password.HashAndSalt([]byte(req.Password))
		if err != nil {
			return err
		}
		tokenActivationExpiredAt := utils.TimeNowJkrt().Add(time.Duration(u.cfg.Setting.ActivationTokenInMinute) * time.Minute)
		err = u.userRepo.UpdateUserPasswordPinTokenActivationStatusByUserIDParams(ctx, domain.User{
			ID:                  user.ID,
			ActivationToken:     tokenHash,
			Password:            pass,
			Status:              1,
			ActivationExpiredAt: tokenActivationExpiredAt,
		}, tx)
		if err != nil {
			return err
		}
		return nil
	})
	resp.Token = token
	return resp, err
}

func (u *useCase) createToken(user domain.User, role domain.RoleSimplified, permissions []domain.PermissionSimplified) (token string, tokenExpiredAt time.Time, err error) {
	tokenExpiredAt = time.Now().Add(time.Duration(u.cfg.Jwt.TokenExpiredInMinutes) * time.Minute)
	token, err = jwt.CreateToken(jwt.CreateTokenRequest{
		SecretToken: u.cfg.Jwt.Secret,
		Claims: jwt.MyClaims[domain.UserRolePermission]{
			Data: domain.UserRolePermission{
				ID:         user.ID,
				FullName:   user.FullName,
				Email:      user.Email,
				RoleID:     user.RoleID,
				RoleName:   role.Name,
				Permission: permissions,
			},
			Claims: jwt2.RegisteredClaims{
				Issuer:    "finansiku",
				ExpiresAt: jwt2.NewNumericDate(tokenExpiredAt),
			},
		},
	})
	if err != nil {
		return
	}
	return
}

func (u *useCase) Login(ctx context.Context, req LoginRequest) (resp LoginResponse, err error) {
	span, ctx := tracer.StartSpan(ctx, tracer.GetFuncName(u.Login))
	defer span.End()
	// get user by email : done
	user, err := u.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return
	}
	role, err := u.roleRepo.GetRoleByID(ctx, user.RoleID)
	if err != nil {
		return
	}
	permissions, err := u.permissionRepo.GetAllPermissionsByRoleID(ctx, role.ID)
	if err != nil {
		return
	}

	// validate the password
	isValid := password.ComparePasswords(user.Password, []byte(req.Password))
	if !isValid {
		return resp, constant.ErrWrongPassword
	}
	if user.Status != domain.UserStatusActive {
		return resp, constant.ErrUserIsDeactivated
	}
	token, tokenExpiredAt, err := u.createToken(user, role, permissions)
	resp = LoginResponse{
		Token:     token,
		ExpiredAt: tokenExpiredAt.Format(time.RFC3339),
		User: domain.UserSimplified{
			ID:        user.ID,
			FullName:  user.FullName,
			Email:     user.Email,
			Status:    user.Status,
			RoleID:    user.RoleID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}
	return resp, nil
}

func (u *useCase) CreateUser(ctx context.Context, req CreateUserRequest) error {
	span, ctx := tracer.StartSpan(ctx, tracer.GetFuncName(u.CreateUser))
	defer span.End()

	user, err := u.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, constant.ErrUserNotFound) {
		return err
	}

	if user.ID != 0 {
		return constant.ErrUserAlreadyExists
	}

	err = u.txManager.Run(ctx, func(tx pgx.Tx) error {
		token, err := utils.GenerateRandomToken(TokenLength)
		if err != nil {
			return err
		}
		tokenHash, err := password.HashAndSalt([]byte(token))
		if err != nil {
			return err
		}
		if config.IsDevelopment() {
			u.log.Infof(ctx, "create user token: %s", token)
		}
		tokenActivationExpiredAt := utils.TimeNowJkrt().Add(time.Duration(u.cfg.Setting.ActivationTokenInMinute) * time.Minute)

		errTx := u.userRepo.InsertUserWithoutPassword(ctx, domain.User{
			Email:               req.Email,
			FullName:            req.FullName,
			RoleID:              req.RoleID,
			ActivationToken:     tokenHash,
			ActivationExpiredAt: tokenActivationExpiredAt,
		}, tx)
		if errTx != nil {
			return errTx
		}
		go utils.CallWithErrorWrapLog(u.log, func() error {
			return u.email.SendEmailActivation(context.Background(), email.SendEmailActivationParam{
				Email: req.Email,
				Token: token,
				Name:  req.FullName,
			})
		})
		return nil
	})
	return err
}

func (u *useCase) GetUserByEmail(ctx context.Context, req GetUserByEmailRequest) (domain.User, error) {
	span, ctx := tracer.StartSpan(ctx, tracer.GetFuncName(u.GetUserByEmail))
	defer span.End()

	user, err := u.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func NewUseCase(cfg *config.Config,
	userRepo UserRepo,
	roleRepo RoleRepo,
	permissionRepo PermissionRepo,
	txManager *db.TxManager,
	email email.Email,
	log loggerw.Logger) UseCase {
	return &useCase{cfg: cfg,
		userRepo:       userRepo,
		roleRepo:       roleRepo,
		permissionRepo: permissionRepo,
		txManager:      txManager,
		email:          email, log: log}
}

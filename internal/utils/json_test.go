package utils

import (
	"finansiku-frontoffice-svc/internal/adapters/outbound/db/postgres/sqlc"
	"finansiku-frontoffice-svc/internal/domain"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"testing"
)

func TestObjectToObject(t *testing.T) {
	type args struct {
		obj2 []domain.User
		obj1 []sqlc.GetUserByEmailRow
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				obj2: make([]domain.User, 0),
				obj1: []sqlc.GetUserByEmailRow{
					{Email: "panjaitannadree@gmail.com",
						Password: pgtype.Text{String: "testing", Valid: true},
						ID:       1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ObjectToObject(&tt.args.obj1, &tt.args.obj2); (err != nil) != tt.wantErr {
				t.Errorf("ObjectToObject() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println(tt.args.obj1)
			fmt.Println(tt.args.obj2)
		})
	}
}

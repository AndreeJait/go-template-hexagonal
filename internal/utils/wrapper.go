package utils

import (
	"context"
	"github.com/AndreeJait/go-utility/loggerw"
)

func CallWithErrorWrapLog(logger loggerw.Logger, callback func() error, message ...interface{}) {
	if err := callback(); err != nil {
		logger.Error(context.Background(), err, message)
	}
}

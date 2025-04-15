package data

import (
	"github.com/WH-5/push-service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type pushRepo struct {
	data *Data
	log  *log.Helper
}

func NewPushRepo(data *Data, logger log.Logger) biz.PushRepo {
	return &pushRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

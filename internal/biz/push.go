package biz

import (
	"github.com/WH-5/push-service/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
)

type PushRepo interface {
}
type PushUsecase struct {
	repo PushRepo
	log  *log.Helper
	CF   *conf.Bizfig
}

func NewPushUsecase(cf *conf.Bizfig, repo PushRepo, logger log.Logger) *PushUsecase {
	return &PushUsecase{CF: cf, repo: repo, log: log.NewHelper(logger)}
}

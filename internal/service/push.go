package service

import (
	"context"
	v1 "github.com/WH-5/push-service/api/user/v1"
	"github.com/WH-5/push-service/internal/biz"
	"github.com/WH-5/push-service/internal/conf"
	"github.com/WH-5/push-service/internal/pkg"
	"log"

	pb "github.com/WH-5/push-service/api/push/v1"
)

type PushService struct {
	pb.UnimplementedPushServer
	UC         *biz.PushUsecase
	UserClient v1.UserClient
}

func NewPushService(c *conf.Server, usecase *biz.PushUsecase) *PushService {

	uc := pkg.UserClient(c.Registry.GetConsul())
	if uc == nil {
		log.Fatal("user client is nil — check consul config")
	}
	return &PushService{
		UC:         usecase,
		UserClient: uc,
	}
}

func (s *PushService) PushMsg(ctx context.Context, req *pb.PushMsgRequest) (*pb.PushMsgReply, error) {
	return &pb.PushMsgReply{}, nil
}

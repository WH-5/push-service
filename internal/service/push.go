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

// PushMsg 推送消息
func (s *PushService) PushMsg(ctx context.Context, req *pb.PushMsgRequest) (*pb.PushMsgReply, error) {

	err := s.UC.PushMessage(uint(req.GetUserId()), req.GetMsg())
	if err != nil {
		return nil, PushMessageError(err)
	}
	return &pb.PushMsgReply{Msg: "send success"}, nil
}

// GetOnlineStatus 获取在线状态
func (s *PushService) GetOnlineStatus(ctx context.Context, req *pb.GetOnlineStatusRequest) (*pb.GetOnlineStatusReply, error) {
	return &pb.GetOnlineStatusReply{}, nil
}

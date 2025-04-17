package service

import (
	"context"
	"errors"
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

	//获取目标id
	uid:=uint(req.GetUserId())
	//根据消息类型解析消息
	switch req.GetMsgType() {
	case 1:
		//聊天消息
		var chat pb.ChatPayload
		if err := req.Payload.UnmarshalTo(&chat);err!=nil{
			return nil,PushMessageError(err)
		}
	case 2:
		//好友消息
		var friend pb.FriendPayload
		if err := req.Payload.UnmarshalTo(&friend);err!=nil{
			return nil,PushMessageError(err)
		}
	default:
		return nil, PushMessageError(errors.New("未知消息类型"))
	}
	
	

	err := s.UC.PushMessage(, req.GetMsg())
	if err != nil {
		return nil, PushMessageError(err)
	}
	return &pb.PushMsgReply{Msg: "send success",Timestamp: req.GetTimestamp()}, nil
}

// GetOnlineStatus 获取在线状态
func (s *PushService) GetOnlineStatus(ctx context.Context, req *pb.GetOnlineStatusRequest) (*pb.GetOnlineStatusReply, error) {
	return &pb.GetOnlineStatusReply{}, nil
}

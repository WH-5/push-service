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

	//uidValue := ctx.Value("user_id")
	//sid, ok := uidValue.(float64)
	//if !ok {
	//	return nil, PushMessageError(errors.New("invalid or missing user_id in context"))
	//}
	////获取目标id
	//uid := uint(req.GetUserId())

	//不在这里处理消息
	////根据消息类型解析消息
	//switch req.GetMsgType() {
	//case 1:
	//	//聊天消息
	//	var chat pb.ChatPayload
	//	if err := req.Payload.UnmarshalTo(&chat);err!=nil{
	//		return nil,PushMessageError(err)
	//	}
	//case 2:
	//	//好友消息
	//	var friend pb.FriendPayload
	//	if err := req.Payload.UnmarshalTo(&friend);err!=nil{
	//		return nil,PushMessageError(err)
	//	}
	//default:
	//	return nil, PushMessageError(errors.New("未知消息类型"))
	//}

	sid := req.GetSelfUserId()

	// 调用用户服务获取目标 user_id（如果有）
	tid, err := s.UserClient.GetIdByUnique(ctx, &v1.GetIdByUniqueRequest{
		UniqueId: req.ToUnique,
	})
	if err != nil {
		return nil, PushMessageError(err)
	}
	//TODO从user服务获取id的unique
	idMany, err := s.UserClient.GetUniqueByIdMany(ctx, &v1.GetUniqueByIdManyRequest{
		UserId: sid,
	})
	if err != nil {
		return nil, err
	}
	//消息经过base64编码
	err = s.UC.PushMessage(uint(tid.GetUserId()), req.GetPayload(), int(req.GetMsgType()), idMany.GetUniqueId())
	if err != nil {
		return nil, PushMessageError(err)
	}
	return &pb.PushMsgReply{Msg: "send success"}, nil
}

// GetOnlineStatus 获取在线状态
func (s *PushService) GetOnlineStatus(ctx context.Context, req *pb.GetOnlineStatusRequest) (*pb.GetOnlineStatusReply, error) {
	return &pb.GetOnlineStatusReply{}, nil
}

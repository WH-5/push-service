package biz

import (
	"encoding/json"
	"github.com/WH-5/push-service/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
)

type PushRepo interface {
	Store(uid uint, msg []byte) error
	PopAll(uid uint) ([]string, error)
	Online(id uint, conn *websocket.Conn)
	Offline(id uint)
	IsOnline(id uint) bool
	GetConn(id uint) (*websocket.Conn, error)
}
type PushUsecase struct {
	repo PushRepo
	log  *log.Helper
	CF   *conf.Bizfig
}

func NewPushUsecase(cf *conf.Bizfig, repo PushRepo, logger log.Logger) *PushUsecase {
	return &PushUsecase{CF: cf, repo: repo, log: log.NewHelper(logger)}
}

// PushMessage 推送消息
func (u *PushUsecase) PushMessage(userId uint, msg []byte, m_type int) error {
	//在线
	on := u.repo.IsOnline(userId)
	if on {
		//获取连接
		conn, err := u.repo.GetConn(userId)
		if err != nil {
			return err
		}
		//发送消息
		data := map[string]interface{}{
			"type":    m_type,
			"payload": msg,
		}
		b, err := json.Marshal(data)
		err = conn.WriteMessage(websocket.TextMessage, b)
		if err == nil {
			//发送成功
			return nil
		}
		//内部err，出if后是新err
	}
	//不在线或推送失败
	err := u.repo.Store(userId, msg)
	if err != nil {
		return err
	}
	return nil
}

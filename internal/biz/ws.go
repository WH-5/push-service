package biz

import "github.com/gorilla/websocket"

//func (u *PushUsecase) SendMessage() {
//
//}

// OnConnect 上线
func (u *PushUsecase) OnConnect(id uint, conn *websocket.Conn) {

	//设置在线状态
	u.repo.Online(id, conn)

	msgs, err := u.repo.PopAll(id) // 获取所有离线消息
	if err != nil {
		// 可加日志输出错误
		u.log.Fatal("pop all err:", err)
		return
	}

	if len(msgs) == 0 {
		// 没有离线消息，不推送
		u.log.Info("online success have no push message")
		return
	}
	//推送消息
	for _, msg := range msgs {
		if conn != nil {
			err = conn.WriteMessage(websocket.TextMessage, []byte(msg)) // 推送消息
			if err != nil {
				u.log.Fatal("write offline message error:", err)
				return
			}
		}
	}

}

// OnDisconnect 下线
func (u *PushUsecase) OnDisconnect(id uint) {
	u.repo.Offline(id)
}

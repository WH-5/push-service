package biz

import "github.com/gorilla/websocket"

func (u *PushUsecase) SendMessage() {

}

// OnConnect 上线
func (u *PushUsecase) OnConnect(id uint, conn *websocket.Conn) {

}

// OnDisconnect 下线
func (u *PushUsecase) OnDisconnect(id uint) {

}

// Package data user.go
// Author: 王辉
// Created: 2025-03-30 00:29
// 缓存中键的前缀,由于都有id后缀，全部删除了，只保留用于区分的信息
// O{userID}  o后接userid 值的类型为list，内容为因不在线而没有收到的消息
package data

import (
	"errors"
	"fmt"
	"github.com/WH-5/push-service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
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
func (d *pushRepo) GetConn(id uint) (*websocket.Conn, error) {
	conn, bo := d.data.WSD.Get(id)
	if !bo {
		return nil, errors.New(fmt.Sprintf("user %d not online", id))
	}
	return conn, nil
}
func (d *pushRepo) IsOnline(id uint) bool {
	return d.data.WSD.IsUserOnline(id)
}

func (d *pushRepo) Online(id uint, conn *websocket.Conn) {
	// 加入连接池、设置在线状态
	d.data.WSD.Add(id, conn)
}

func (d *pushRepo) Offline(id uint) {
	//退出连接池、退出在线用户map
	d.data.WSD.Remove(id)
}

var _ biz.PushRepo = (*pushRepo)(nil)

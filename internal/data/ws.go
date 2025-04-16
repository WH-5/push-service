package data

import (
	"github.com/gorilla/websocket"
	"sync"
)

type WSData struct {
	//连接池
	CM *ConnManager
	//在线用户
	OnlineUsers sync.Map // map[userID] -> true
}

func NewWSData() *WSData {
	return &WSData{
		CM: &ConnManager{
			conns: sync.Map{}, // 初始化 ConnManager 的 conns 字段
		},
		OnlineUsers: sync.Map{}, // 初始化 OnlineUsers 字段
	}
}

type ConnManager struct {
	conns sync.Map // map[userID]string -> *websocket.Conn
}

// Add 注册连接，踢掉旧连接
func (ws *WSData) Add(userID uint, conn *websocket.Conn) {
	if old, ok := ws.CM.conns.Load(userID); ok {
		_ = old.(*websocket.Conn).Close() // 踢掉旧连接
	}
	ws.CM.conns.Store(userID, conn)
	ws.SetUserOnline(userID) // 通知上线
}

// Get 获取连接
func (ws *WSData) Get(userID uint) (*websocket.Conn, bool) {
	conn, ok := ws.CM.conns.Load(userID)
	return conn.(*websocket.Conn), ok
}

// Remove 移除连接
func (ws *WSData) Remove(userID uint) {
	ws.CM.conns.Delete(userID)
	ws.SetUserOffline(userID) // 通知下线
}

//设置在线状态

// SetUserOnline 设置在线
func (ws *WSData) SetUserOnline(userID uint) {
	ws.OnlineUsers.Store(userID, true)
}

// SetUserOffline 设置下线
func (ws *WSData) SetUserOffline(userID uint) {
	ws.OnlineUsers.Delete(userID)
}

// IsUserOnline 获取状态
func (ws *WSData) IsUserOnline(userID uint) bool {
	_, ok := ws.OnlineUsers.Load(userID)
	return ok
}

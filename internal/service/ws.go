package service

import (
	"fmt"
	"github.com/WH-5/push-service/internal/pkg"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// 升级器：把 HTTP 协议升级为 WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 不校验来源，方便测试
		return true
	},
}

// NewWSHandler 处理 WebSocket 连接的函数
func NewWSHandler(service *PushService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		//「

		//验证jwt
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := pkg.ParseToken(tokenStr, service.UC.CF.JWT_SECRET_KEY)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		if err != nil || !token.Valid {
			http.Error(w, fmt.Sprintf("token 无效: %v", err), http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "无法解析 Claims", http.StatusUnauthorized)
			return
		}
		//把用户ID和token放入上下文
		uid := claims["user_id"]

		//」

		// 升级为 WebSocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("升级 WebSocket 失败:", err)
			return
		}
		defer func(conn *websocket.Conn) {
			err := conn.Close()
			if err != nil {
				return
			}
		}(conn)

		//成功建立连接
		service.UC.OnConnect(uint(uid.(float64)), conn)
		fmt.Println("WebSocket 已连接")

		// 设置心跳：60秒内未收到任何消息将断开连接
		err = conn.SetReadDeadline(time.Now().Add(20 * time.Second))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
			return
		}
		conn.SetPongHandler(func(string) error {
			err := conn.SetReadDeadline(time.Now().Add(20 * time.Second))
			if err != nil {
				return err
			}
			return nil
		})
		for {
			// 读取消息
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				//断开了
				service.UC.OnDisconnect(uint(uid.(float64)))
				log.Println("读取消息失败:", err)
				break
			}
			log.Printf("收到消息: %s\n", message)

			// 向客户端原样发送消息
			reply := fmt.Sprintf("你发的是：%s", message)
			err = conn.WriteMessage(messageType, []byte(reply))
			if err != nil {
				//同样断开
				service.UC.OnDisconnect(uint(uid.(float64)))
				log.Println("发送消息失败:", err)
				break
			}
		}
	}

}

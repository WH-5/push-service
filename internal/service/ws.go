package service

import (
	"context"
	"encoding/json"
	"fmt"
	v1 "github.com/WH-5/push-service/api/user/v1"
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
		// 目前不校验来源，方便测试

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
		//上线函数
		service.UC.OnConnect(uint(uid.(float64)), conn)
		fmt.Println("WebSocket 已连接")

		// 设置心跳：60秒内未收到任何消息将断开连接
		err = conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
			return
		}
		conn.SetPongHandler(func(string) error {
			err := conn.SetReadDeadline(time.Now().Add(60 * time.Second))
			if err != nil {
				return err
			}
			return nil
		})
		for {
			// 读取消息
			_, message, err := conn.ReadMessage()
			if err != nil {
				//断开了
				//下线函数
				service.UC.OnDisconnect(uint(uid.(float64)))
				log.Println("读取消息失败:", err)
				break
			}

			log.Printf("收到消息: %s\n", message)

			var msgData map[string]interface{}
			if err := json.Unmarshal(message, &msgData); err != nil {
				log.Println("无法解析消息为 JSON:", err)
				continue
			}

			mType, ok := msgData["type"].(float64) // JSON 中数字默认是 float64
			if !ok {
				log.Println("消息缺少 type 字段或类型错误")
				continue
			}

			if int(mType) == 1 {

				log.Println("收到聊天消息，待处理...")

				tidStr, ok := msgData["target_unique"].(string)
				if !ok {
					log.Println("缺少或非法 target_unique 字段")
					continue
				}
				sid := uint(uid.(float64))
				tid, err := service.UserClient.GetIdByUnique(context.Background(), &v1.GetIdByUniqueRequest{
					UniqueId: tidStr,
				})
				if err != nil {
					log.Println(err)
				}

				idMany, err := service.UserClient.GetUniqueByIdMany(context.Background(), &v1.GetUniqueByIdManyRequest{
					UserId: uint64(sid),
				})
				if err != nil {
					return
				}
				err = service.UC.PushMessage(uint(tid.UserId), message, int(mType), idMany.GetUniqueId())
				if err != nil {
					log.Printf("转发消息失败: %v", err)
				}
				log.Printf("转发消息成功")

			} else if int(mType) == 2 {
				// TODO:处理好友消息
				log.Printf("收到好友消息")
			} else if int(mType) == 0 {

				log.Printf("收到ping消息")
			} else {
				log.Printf("收到其他类型消息：%v", mType)
			}
		}
	}

}

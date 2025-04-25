package httpex

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	websocketCounter int
)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有来源
	},
}

func handleWebSocketProxy(w http.ResponseWriter, r *http.Request) {
	// 升级 HTTP 连接到 WebSocket
	clientConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}

	// 连接到后端 WebSocket 服务器
	backendConn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8081/ws", nil)
	if err != nil {
		log.Println("Error connecting to backend WebSocket server:", err)
		return
	}

	websocketCounter++
	log.Printf("new websocket connection(total:%d): client=%s->%s backend=%s->%s", websocketCounter,
		clientConn.RemoteAddr().String(), clientConn.LocalAddr().String(),
		backendConn.LocalAddr().String(), backendConn.RemoteAddr().String())
	defer func() {
		websocketCounter--
		log.Printf("websocket proxy finish (total:%d)", websocketCounter)
	}()

	// 使用通道来通知 goroutine 关闭
	// 创建一个上下文和取消函数
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(2)
	// 启动 goroutine 从客户端读取消息并转发到后端
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				log.Println("proxy backend disconnected")
				return
			default:
				msgType, msg, err := clientConn.ReadMessage()
				if err != nil {
					log.Println("Error reading message from client:", err)
					cancel()
					return
				}
				log.Printf("proxy Received message from client %s->%s: %s", clientConn.RemoteAddr().String(),
					clientConn.LocalAddr().String(), msg)
				// 将消息转发到后端 WebSocket 服务器
				err = backendConn.WriteMessage(msgType, msg)
				if err != nil {
					log.Println("Error writing message to backend:", err)
					cancel()
					return
				}
				log.Printf("proxy Sent message to backend %s->%s: %s", backendConn.LocalAddr().String(),
					backendConn.RemoteAddr().String(), msg)
			}
		}
	}()

	// 从后端读取消息并转发到客户端
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				log.Println("proxy Client disconnected")
				return
			default:
				msgType, msg, err := backendConn.ReadMessage()
				if err != nil {
					log.Println("Error reading message from backend:", err)
					cancel()
					return
				}
				log.Printf("proxy Received message from backend %s->%s: %s", backendConn.RemoteAddr().String(),
					backendConn.LocalAddr().String(), msg)
				// 将消息转发到客户端
				err = clientConn.WriteMessage(msgType, msg)
				if err != nil {
					log.Println("Error writing message to client:", err)
					cancel()
					return
				}
				log.Printf("proxy Sent message to client %s->%s: %s", clientConn.LocalAddr().String(),
					clientConn.RemoteAddr().String(), msg)
			}
		}
	}()
	<-ctx.Done()
	backendConn.Close()
	clientConn.Close()
	wg.Wait()

}

// WebSocketProxy WebSocket 代理
func WebSocketProxy() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", handleWebSocketProxy)
	fmt.Println("WebSocket proxy server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handleWebsocketServer(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()
	defer func() {
		log.Println("Client disconnected")
	}()

	log.Println("Client connected")

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("server Error reading message:", err)
			break
		}
		log.Printf("server Received message from %s : %s", conn.RemoteAddr().String(), msg)

		// Echo the message back to the client
		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			log.Println("server Error writing message:", err)
			break
		}
		log.Printf("server Sent message to %s : %s", conn.RemoteAddr().String(), msg)
	}
}

func WebSocketServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", handleWebsocketServer)
	fmt.Println("WebSocket server is running on :8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}

func WebSocketDaemon() {
	go func() {
		WebSocketProxy()
	}()
	WebSocketServer()
}

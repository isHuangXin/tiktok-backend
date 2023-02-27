package test

import (
	"encoding/json"
	"fmt"
	"github.com/isHuangXin/tiktok-backend/api"
	"io"
	"net"
	"testing"
	"time"
)

func TestMessageServer(t *testing.T) {
	e := newExpect(t)
	userIdA, _ := getTestUserToken(testUserA, e)
	userIdB, _ := getTestUserToken(testUserB, e)

	connA, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("Connect server failed: #{err}\n")
		return
	}
	connB, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("Connect server failed: #{err}\n")
		return
	}

	createChat(userIdA, connA, userIdB, connB)

	go readMessage(connB)
	sendMessage(userIdA, userIdB, connA)
}

func readMessage(conn net.Conn) {
	defer conn.Close()

	var buf [256]byte
	for {
		n, err := conn.Read(buf[:])
		if n == 0 {
			if err == io.EOF {
				break
			}
			fmt.Println("Read message failed: #{err}\n")
			continue
		}

		var event = api.MessagePushEvent{}
		_ = json.Unmarshal(buf[:n], &event)
		fmt.Println("Read message: #{event}\n")
	}
}

func sendMessage(fromUserId int, toUserId int, fromConn net.Conn) {
	defer fromConn.Close()

	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		sendEvent := api.MessageSendEvent{
			UserId:     int64(fromUserId),
			ToUserId:   int64(toUserId),
			MsgContent: "Test Content",
		}
		data, _ := json.Marshal(sendEvent)
		_, err := fromConn.Write(data)
		if err != nil {
			fmt.Printf("Send message failed: %v\n", err)
			return
		}
	}
	time.Sleep(time.Second)
}

func createChat(userIdA int, connA net.Conn, userIdB int, connB net.Conn) {
	chatEventA := api.MessageSendEvent{
		UserId:   int64(userIdA),
		ToUserId: int64(userIdB),
	}
	chatEventB := api.MessageSendEvent{
		UserId:   int64(userIdB),
		ToUserId: int64(userIdA),
	}
	eventA, _ := json.Marshal(chatEventA)
	eventB, _ := json.Marshal(chatEventB)
	_, err := connA.Write(eventA)
	if err != nil {
		fmt.Printf("Create chatA failed: %v\n", err)
		return
	}
	_, err = connB.Write(eventB)
	if err != nil {
		fmt.Printf("Create chatB failed: %v\n", err)
		return
	}
}

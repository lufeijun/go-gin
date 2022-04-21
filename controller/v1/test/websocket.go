package test

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	wg      sync.WaitGroup
	wgCount = 50 // 协程数量
	wgConns = 5  // 每个协程中，连接的个数
)

func SocketOne(c *gin.Context) {

	wg.Add(wgCount)

	for i := 0; i < wgCount; i++ {
		go tickWriter(i)
	}

	wg.Wait()

	c.String(http.StatusOK, "ok")
	return
}

func tickWriter(num int) {

	var i int

	dialer := websocket.Dialer{}

	// connects := make([]*websocket.Conn, wgConns)
	var connects []*websocket.Conn
	var conn *websocket.Conn
	var err error
	var realConn int

	for i = 0; i < wgConns; i++ {
		conn, _, err = dialer.Dial("ws://www.jipeng.com", nil)
		if err != nil {
			continue
		}
		realConn++
		connects = append(connects, conn)
	}

	defer wg.Done()
	for i = 0; i < realConn; i++ {
		defer connects[i].Close()
	}

	// 十次循环，给每个 connect 发送十条数据
	for i := 0; i < 10; i++ {
		//向客户端发送类型为文本的数据
		for index := 0; index < realConn; index++ {
			// aa := connects[index]
			// fmt.Println(aa)
			err = connects[0].WriteMessage(websocket.TextMessage, []byte("进程号："+strconv.Itoa(num)+" == 连接号"+strconv.Itoa(index)+" == 发送消息了 ==== 消息号："+strconv.Itoa(i)))
			if nil != err {
				fmt.Println(err)
			}
		}

		//休息一秒
		time.Sleep(time.Second)
	}

}

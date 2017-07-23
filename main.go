package main

import (
	"encoding/json"
	"fmt"
	"github.com/ami-go/baduk-engine"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request, board *libaduk.AbstractBoard) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(string(msg))

		move := libaduk.Move{}

		json.Unmarshal(msg, &move)

		move.Color = 1
		board.PlayMove(move)

		fmt.Printf("%v", board.ToString())
		conn.WriteMessage(t, msg)
	}
}

func main() {
	var err error
	var board *libaduk.AbstractBoard

	if board, err = libaduk.NewBoard(19); err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Static("static", "static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request, board)
	})

	r.Run(":8080")

}

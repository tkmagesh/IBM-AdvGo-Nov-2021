package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var wsConections []*websocket.Conn

func reader(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))
		broadcast(string(p))
	}
}

func broadcast(msg string) {
	for _, wsConn := range wsConections {
		if er := wsConn.WriteMessage(1, []byte(msg)); er != nil {
			log.Println(er)
			return
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndPoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	wsConections = append(wsConections, ws)
	fmt.Println("client count : ", len(wsConections))
	log.Println("Client connected")
	err = ws.WriteMessage(1, []byte("Hello Client"))
	if err != nil {
		log.Println(err)
		return
	}
	go reader(ws)
}

func wsTime(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Client connected")
	timer := time.Tick(5 * time.Second)
	for {
		select {
		case t := <-timer:
			err := ws.WriteMessage(1, []byte(t.String()))
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/ws", wsEndPoint)
	http.HandleFunc("/wsTime", wsTime)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

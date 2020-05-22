package UserModel

import (
	"awesomeProject.example.com/Structures"
	"golang.org/x/net/websocket"
	"log"
)
//recive message from clients and auto send message from server
func Socket(ws *websocket.Conn) {

	for {
		// allocate our container struct
		var m Structures.MessageStruct
		// receive a message using the codec
		if err := websocket.JSON.Receive(ws, &m); err != nil {
			log.Println(err)
			break
		}
		log.Println("Received message:", m.Message)
		//--------- add it in database--------------
		AddMessage(m)
		// send a response
		m2 := Structures.MessageStruct{Message:"Thanks for the message!",
			From:"server"}
		if err := websocket.JSON.Send(ws, m2); err != nil {
			log.Println(err)
			break
		}
	}
}


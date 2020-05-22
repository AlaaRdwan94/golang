package UserModel

import (
	"awesomeProject.example.com/Structures"
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"time"
)

func AddMessage(m Structures.MessageStruct){
	m.MessID = getCurrentMsgID()
	m.Time = time.Now().UTC()
	fmt.Println("message", m)
	fmt.Println("all", GetAllMessages())
	value, _ := json.Marshal(m)
	MsgOpendatabase()
	err := MDB.Put([]byte(m.MessID), []byte(value), nil)
	if err != nil {
		fmt.Println("message not saved in database",err)
	}
}


func getCurrentMsgID() string {

	allIDs := GetAllMsgsID()
	var CurrentID string
	if allIDs == nil {
		CurrentID,_ = ConvertIntToFixedLengthString(0,20)
		return CurrentID
	}else{
		lastID := ConvertFixedLengthStringtoInt(allIDs[len(allIDs)-1])+1
		CurrentID,_ = ConvertIntToFixedLengthString(lastID,20)
	}
	return CurrentID
}
//get all users from database
func GetAllMessages() (values []Structures.MessageStruct) {
	MsgOpendatabase()
	iter := MDB.NewIterator(nil, nil)
	for iter.Next() {

		value := iter.Value()
		var newdata Structures.MessageStruct
		json.Unmarshal(value, &newdata)
		values = append(values, newdata)
	}
	return values
}
//get all IDs from database
func GetAllMsgsID() (ID []string) {

	MsgOpendatabase()
	iter := MDB.NewIterator(nil, nil)
	for iter.Next() {

		value := iter.Value()
		var newdata Structures.MessageStruct
		json.Unmarshal(value, &newdata)
		ID = append(ID, newdata.MessID)
	}
	fmt.Println("I reach here >>>>>>>>>>>>>>>>>>>>>>>")
	return ID
}


//---------API for Admin messages room-----------

func GetAllMessagesRooms(ws *websocket.Conn) {
	fmt.Println("GetAllMessages()",GetAllMessages())
		Rooms := []Structures.Room{}
		for _ , client := range GetAllUsers() {
			clientmsgLst := getAllclientmsgs(client.Login.UserName)
			if  clientmsgLst != nil {
				Rooms = append(Rooms,Structures.Room{
					Client:client.Login.UserName,
					Messages: clientmsgLst})
			}
		}
		if err := websocket.JSON.Send(ws, Rooms); err != nil {
			log.Println(err)

		}
}
func getAllclientmsgs(name string) (m []string) {

	for _ , msg := range GetAllMessages(){
		if msg.From == name{
			m = append(m,msg.Message)
		}
	}
	return m
}

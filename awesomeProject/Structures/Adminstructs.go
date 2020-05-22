package Structures

type AdminStruct struct {
	Admindata UserStruct
	Adminrule bool
	chatrooms []ChatRoom
}

type ChatRoom struct {
	RoomID string
	ClientID string
	messages []MessageStruct
}
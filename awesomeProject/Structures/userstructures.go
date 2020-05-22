package Structures

import "time"

type LoginStruct struct {
	UserName string
	PassWord string
	Lastlogintime time.Time
}

type UserStruct struct {
	ID string
	Login LoginStruct
	RegisterTime time.Time
	Active bool
	Online bool
	Message  MessageStruct
}


package Structures

import "time"

type MessageStruct struct {
	MessID string
	Message string
	From string
	Time time.Time
}

type Room struct {
	Client string
	Messages []string
}
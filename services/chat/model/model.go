package model

type Message struct {
	MessageID      string `json:"message_id"`
	SenderID       string `json:"sender_id"`
	RoomID         string `json:"room_id"`
	MessageContent string `json:"message_content"`
	TimeStamp      string `json:"time_stamp"`
}

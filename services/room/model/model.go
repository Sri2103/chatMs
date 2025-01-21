package model

type Room struct {
	RoomId    string   `json:"room_id"`
	RoomName  string   `json:"room_name"`
	CreatedBy string   `json:"created_by"`
	Users     []string `json:"users"`
}

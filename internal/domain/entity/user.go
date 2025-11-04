package entity

import "time"

// User represents a user in the system
type User struct {
	Username  string    `json:"username" bson:"username"`
	Password  string    `json:"password" bson:"password"`
	FirstName string    `json:"firstname" bson:"firstname"`
	LastName  string    `json:"lastname" bson:"lastname"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

// Group represents a user group in the system
type Group struct {
	GID       string   `json:"gid" bson:"gid"`
	GroupName string   `json:"groupname" bson:"groupname"`
	Members   []string `json:"members" bson:"members"`
}

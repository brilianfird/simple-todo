package entity

import "time"

type ToDo struct {
	Timestamp 	time.Time
	Title 		string
	Description string
	Cleared		bool
}
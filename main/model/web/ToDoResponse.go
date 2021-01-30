package web

import "time"

type ToDoRestWeb struct {
	Timestamp 	time.Time
	Title 		string
	Description string
	Cleared		bool
}

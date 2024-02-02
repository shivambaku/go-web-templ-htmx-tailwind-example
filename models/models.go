package model

import "time"

type User struct {
	Email string
}

type CreateAnimeRecord struct {
	Title    string
	Started  time.Time
	Finished time.Time
	Rating   float32
	Memo     string
}

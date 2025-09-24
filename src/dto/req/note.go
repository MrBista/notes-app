package req

import "time"

type NoteRequestCreate struct {
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	IsFavorite bool      `json:"is_favorite"`
	ReminderAt time.Time `json:"reminder_at"`
}

type NoteRequestUpdate struct {
	Title      string    `json:"title"`
	Status     int       `json:"status"`
	Content    string    `json:"content"`
	IsFavorite bool      `json:"is_favorite"`
	ReminderAt time.Time `json:"reminder_at"`
}

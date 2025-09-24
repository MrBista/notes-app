package res

import (
	"notes-golang/src/models"
	"time"
)

type NoteResponse struct {
	Id         int       `json:"id"`
	Status     int       `json:"status"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	IsFavorite bool      `json:"is_favorite"`
	ReminderAt time.Time `json:"reminder_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func ToNoteResponse(note models.Note) NoteResponse {
	noteResponse := NoteResponse{
		Id:         note.Id,
		Status:     note.Status,
		Title:      note.Title,
		Content:    note.Content,
		IsFavorite: note.IsFavorite,
		ReminderAt: note.ReminderAt,
		CreatedAt:  note.CreatedAt,
		UpdatedAt:  note.UpdatedAt,
	}

	return noteResponse
}

func ToNoteResponses(notes []models.Note) []NoteResponse {
	var noteResponses []NoteResponse

	for _, note := range notes {
		noteResponse := ToNoteResponse(note)
		noteResponses = append(noteResponses, noteResponse)
	}

	return noteResponses
}

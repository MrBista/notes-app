package service

import (
	"errors"
	"notes-golang/src/dto/req"
	"notes-golang/src/dto/res"
	"notes-golang/src/models"
	"notes-golang/src/repository"

	"gorm.io/gorm"
)

type NoteService interface {
	FindAllNote() ([]res.NoteResponse, error)
	FindNoteById(id int) (res.NoteResponse, error)
	UpdateNoteById(id int, note req.NoteRequestUpdate) error
	CreateNote(note req.NoteRequestCreate) error
	DeleteNoteById(id int) error
}

type NoteServiceImpl struct {
	DB             *gorm.DB
	NoteRepository repository.NoteRepository
}

func NewNoteService(Db *gorm.DB, noteRepository repository.NoteRepository) NoteService {
	return &NoteServiceImpl{
		DB:             Db,
		NoteRepository: noteRepository,
	}
}

func (n *NoteServiceImpl) FindAllNote() ([]res.NoteResponse, error) {

	notes, err := n.NoteRepository.FindAllNote()

	return res.ToNoteResponses(notes), err

}

func (n *NoteServiceImpl) FindNoteById(id int) (res.NoteResponse, error) {
	note, err := n.NoteRepository.FindNoteById(id)

	if errors.Is(gorm.ErrRecordNotFound, err) {
		return res.NoteResponse{}, errors.New("note not found")
	}

	return res.ToNoteResponse(note), err
}
func (n *NoteServiceImpl) UpdateNoteById(id int, note req.NoteRequestUpdate) error {

	findNote, err := n.NoteRepository.FindNoteById(id)

	if err != nil {
		return errors.New("note not found")
	}

	noteBody := models.Note{
		Id:         findNote.Id,
		Status:     findNote.Status,
		Title:      findNote.Title,
		Content:    findNote.Content,
		IsFavorite: findNote.IsFavorite,
		ReminderAt: findNote.ReminderAt,
	}

	err = n.NoteRepository.UpdateNote(noteBody)

	if err != nil {
		return errors.New("failed update note")
	}

	return nil
}

func (n *NoteServiceImpl) CreateNote(note req.NoteRequestCreate) error {
	statusActive := 1
	noteBody := models.Note{
		Status:     statusActive,
		Title:      note.Title,
		Content:    note.Content,
		IsFavorite: note.IsFavorite,
		ReminderAt: note.ReminderAt,
	}
	err := n.NoteRepository.CreateNote(noteBody)

	return err
}

func (n *NoteServiceImpl) DeleteNoteById(id int) error {
	_, err := n.NoteRepository.FindNoteById(id)

	if err != nil {
		return errors.New("note not found")
	}

	err = n.NoteRepository.DeleteNoteById(id)

	return err
}

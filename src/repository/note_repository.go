package repository

import (
	"notes-golang/src/models"

	"gorm.io/gorm"
)

type NoteRepository interface {
	CreateNote(note models.Note) error
	UpdateNote(note models.Note) error
	DeleteNoteById(id int) error
	FindAllNote() ([]models.Note, error)
	FindNoteById(idNote int) (models.Note, error)
}

type NoteRepositoryImpl struct {
	DB *gorm.DB
}

func NewNoteRepository(DB *gorm.DB) NoteRepository {
	return &NoteRepositoryImpl{
		DB: DB,
	}
}

func (n *NoteRepositoryImpl) CreateNote(note models.Note) error {
	err := n.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&note).Error
		return err
	})

	return err
}

func (n *NoteRepositoryImpl) UpdateNote(note models.Note) error {
	err := n.DB.Where("id = ?", note.Id).Updates(note).Error

	return err
}

func (n *NoteRepositoryImpl) DeleteNoteById(id int) error {
	return n.DB.Delete(&models.Note{}, "id = ?", id).Error
}

func (n *NoteRepositoryImpl) FindAllNote() ([]models.Note, error) {
	var notes []models.Note

	result := n.DB.Find(&notes)

	return notes, result.Error
}

func (n *NoteRepositoryImpl) FindNoteById(idNote int) (models.Note, error) {
	var note models.Note

	result := n.DB.Take(&note, "id = ?", idNote)

	return note, result.Error

}

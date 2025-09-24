package controller

import (
	"encoding/json"
	"net/http"
	"notes-golang/src/dto/req"
	"notes-golang/src/dto/res"
	"notes-golang/src/handler"
	"notes-golang/src/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type NoteController interface {
	FindAllNote(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindNoteById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	CreateNote(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateNote(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	DeleteNoteById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type NoteControllerImpl struct {
	NoteService service.NoteService
}

func NewNoteController(noteService service.NoteService) NoteController {
	return &NoteControllerImpl{
		NoteService: noteService,
	}
}

func (n *NoteControllerImpl) FindAllNote(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	noteResult, err := n.NoteService.FindAllNote()

	if err != nil {
		handler.HandleError(w, err)
		return
	}

	response := res.NewCommonResponseSuccess(noteResult, "Sucessfully find all note", http.StatusOK)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	encoder.Encode(&response)

}

func (n *NoteControllerImpl) FindNoteById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	idNoteParam := params.ByName("id")

	idNote, err := strconv.Atoi(idNoteParam)

	if err != nil {
		handler.HandleError(w, err)
		return
	}

	noteResult, err := n.NoteService.FindNoteById(idNote)

	if err != nil {
		handler.HandleError(w, err)
		return
	}

	response := res.NewCommonResponseSuccess(noteResult, "Sucessfully find all note", http.StatusOK)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	encoder.Encode(&response)

}

func (n *NoteControllerImpl) CreateNote(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var noteBody req.NoteRequestCreate

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&noteBody)

	if err != nil {
		handler.HandleError(w, err)
		return
	}

	err = n.NoteService.CreateNote(noteBody)

	if err != nil {
		handler.HandleError(w, err)
		return
	}

	noteResult := true

	response := res.NewCommonResponseSuccess(noteResult, "Sucessfully create note", http.StatusCreated)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}

func (n *NoteControllerImpl) UpdateNote(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var noteBody req.NoteRequestUpdate

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&noteBody)

	if err != nil {
		handler.HandleError(w, err)
		return
	}

	idNoteParam := params.ByName("id")

	idNote, err := strconv.Atoi(idNoteParam)

	if err != nil {
		handler.HandleError(w, err)
		return
	}

	err = n.NoteService.UpdateNoteById(idNote, noteBody)

	if err != nil {
		handler.HandleError(w, err)
		return
	}

	noteResult := true

	response := res.NewCommonResponseSuccess(noteResult, "Sucessfully update note", http.StatusOK)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}

func (n *NoteControllerImpl) DeleteNoteById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	idNoteParam := params.ByName("id")

	idNote, err := strconv.Atoi(idNoteParam)

	if err != nil {
		handler.HandleError(w, err)
		return
	}

	err = n.NoteService.DeleteNoteById(idNote)

	if err != nil {
		handler.HandleError(w, err)
		return
	}

	noteResult := true

	response := res.NewCommonResponseSuccess(noteResult, "Sucessfully delete note", http.StatusOK)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}

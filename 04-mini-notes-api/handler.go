package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func addNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Post istegi olmali", http.StatusMethodNotAllowed)
		return
	}

	var newNote Note
	if err := json.NewDecoder(r.Body).Decode(&newNote); err != nil {
		http.Error(w, "Gecersin JSON", http.StatusBadRequest)
		return
	}

	newNote.ID = nextID
	nextID++
	notes = append(notes, newNote)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newNote)

}

func getNotesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)

}

func deleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Yalnıca Delete Method olmali", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Eksik id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Gecersiz id", http.StatusBadRequest)
		return
	}

	index := -1
	for i, Note := range notes {
		if Note.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "Not bulunamadi", http.StatusNotFound)
		return
	}

	notes = append(notes[:index], notes[index+1:]...)

	w.WriteHeader(http.StatusNoContent)
}

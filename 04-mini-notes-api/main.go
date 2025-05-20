package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/addNote", addNoteHandler)
	http.HandleFunc("/getNotes", getNotesHandler)
	http.HandleFunc("/deleteNote", deleteNoteHandler)

	log.Println("Sunucu localhost:8082 adresinde baslatildi")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}
}

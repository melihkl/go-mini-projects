package main

type Note struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

var notes []Note
var nextID int = 1

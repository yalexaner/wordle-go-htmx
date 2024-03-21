package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var word = [5]string{"w", "o", "r", "d", "e"}

type Letter struct {
	Value    string
	Position string
}

type Word struct {
	Letters [5]Letter
}

type Board struct {
	Words   [6]Word
	Current int
}

var board = Board{
	Current: 0,
	Words:   [6]Word{},
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/check", handleCheck)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//goland:noinspection GoUnusedParameter
func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html", "template/board.html"))
	if err := tmpl.Execute(w, board); err != nil {
		log.Printf("error executing template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func handleCheck(w http.ResponseWriter, r *http.Request) {
	// read the entered word
	var enteredWord Word
	for i := range 5 {
		value := r.FormValue(strconv.Itoa(i))
		enteredWord.Letters[i] = Letter{Value: value}
	}

	// check for placed and present letters in the entered word
	for i, wordLetter := range word {
		if enteredWord.Letters[i].Value == wordLetter {
			enteredWord.Letters[i].Position = "placed"
			continue
		}

		for j, letter := range enteredWord.Letters {
			if letter.Value == wordLetter && i != j {
				enteredWord.Letters[j].Position = "present"
				break
			}
		}
	}

	board.Words[board.Current] = enteredWord
	board.Current++

	fmt.Println(board)

	tmpl := template.Must(template.ParseFiles("template/board.html"))
	if err := tmpl.ExecuteTemplate(w, "board", board); err != nil {
		log.Printf("error executing template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

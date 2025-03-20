package main

import (
	"bufio"
	"example/note/note"
	"fmt"
	"os"
	"strings"
)

func getUserInput(text string) string {
	fmt.Printf("%v ", text)
	reader := bufio.NewReader(os.Stdin) // command line'ı dinleyen bi reader oluşturuyoruz.

	text, err := reader.ReadString('\n') //okumayı bırakması gereken yer

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content) //kullanıcı girdisi olarak alınan ikili, buraya parametre olarak girildi
	if err != nil {
		fmt.Println(err)
		return
	}

	note.Output(userNote)
	err = note.Save(userNote)

	if err != nil {
		fmt.Println("Saving the note is failed")
		return
	} else {
		fmt.Println("Saving the note succeeded!")
	}
}

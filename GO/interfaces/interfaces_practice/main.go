package main

import (
	"bufio"
	"example/interfaces/note"
	"example/interfaces/todo"
	"fmt"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type displayer interface {
	Output()
	saver // embadding yaptık
}

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

func saveData(x saver) error {
	err := x.Save()
	if err != nil {
		fmt.Println("Saving the todo is failed")
		return err
	}

	fmt.Println("Saving the todo succeeded!")
	return nil
}

func display(x displayer) error {
	x.Output()
	return saveData(x)
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	todoTxt, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content) //kullanıcı girdisi olarak alınan ikili, buraya parametre olarak girildi
	if err != nil {
		fmt.Println(err)
		return
	}

	err = display(todoTxt)

	if err != nil {
		return
	}

	err = display(userNote)

	if err != nil {
		return
	}
}

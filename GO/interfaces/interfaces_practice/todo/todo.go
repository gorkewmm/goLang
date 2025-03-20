package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(todo string) (Todo, error) { // constructor function
	if todo == "" {
		return Todo{}, errors.New("invalid input") //kullanıcı girdileri boşsa boş Note verir.
	}
	return Todo{
		Text: todo,
	}, nil
}

func (x Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(x) //Structumuzu json formatında texte çevirir

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func (x Todo) Output() {
	fmt.Println(x.Text)
}

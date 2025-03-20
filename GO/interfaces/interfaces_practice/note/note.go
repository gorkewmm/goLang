package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"` //struct tags, json paketinin çıktılarını ve işlevlerini değiştirir
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) { // constructor function
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input") //kullanıcı girdileri boşsa boş Note verir.
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (x Note) Save() error {
	fileName := strings.ReplaceAll(x.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(x) //Structumuzu json formatında texte çevirir

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func (x Note) Output() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", x.Title, x.Content)
}

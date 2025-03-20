package fileops

import (
	"errors"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	// strbalance := fmt.Sprintf("%2.f", balance)
	// os.WriteFile("balance.txt", []byte(strbalance), 0644)

	b := []byte(strconv.FormatFloat(value, 'f', 3, 64))
	os.WriteFile(fileName, b, 0644)
}

func ReadFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("dosya bulunamadı")
	}
	var str = string(data)
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 1000, errors.New("değer parse edilemedi")
	}
	return value, nil

}

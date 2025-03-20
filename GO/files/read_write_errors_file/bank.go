package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func readFromFile() (float64, error) {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		return 1000, errors.New("dosya bulunamadı!")
	}
	var str = string(data)
	flt, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 1000, errors.New("bakiye degeri parse edilemedi!")
	}
	return flt, nil

}

func writeToFile(balance float64) {
	// strbalance := fmt.Sprintf("%2.f", balance)
	// os.WriteFile("balance.txt", []byte(strbalance), 0644)

	b := []byte(strconv.FormatFloat(balance, 'f', 3, 64))
	os.WriteFile("balance.txt", b, 0644)
}

func main() {
	accountBalance, err := readFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
	}

	fmt.Println("Go bankacılıgına hoşgeldiniz")
	for {
		fmt.Println("Hesap bakiyeni öğren: 1")
		fmt.Println("Para yatır: 2")
		fmt.Println("Para çek: 3")
		fmt.Println("Çıkış yap: 4")
		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Hesap bakiyeniz:", accountBalance)
			readFromFile()
		} else if choice == 2 {
			var depmon float64
			fmt.Println("Ne kadar para yatırmak istersiniz?")
			fmt.Scan(&depmon)
			if depmon <= 0 {
				fmt.Println("Geçersiz miktar")
				continue
			}
			accountBalance = accountBalance + depmon
			fmt.Println("Güncel bakiyeniz: ", accountBalance)
			writeToFile(accountBalance)
		} else if choice == 3 {
			var drawmon float64
			fmt.Println("How much money do you want to withdraw")
			fmt.Scan(&drawmon)
			if drawmon <= 0 {
				fmt.Println("Geçersiz miktar")
				continue
			}
			if drawmon > accountBalance {
				fmt.Println("Hesabınızda yeterli bakiye yok!")
				continue
			}
			accountBalance = accountBalance - drawmon
			fmt.Println("Güncel bakiyeniz: ", accountBalance)
			writeToFile(accountBalance)
		} else if choice == 4 {
			println("Görüşürüz")
			return
		} else {
			fmt.Println("Geçersiz miktar")
		}
	}

}

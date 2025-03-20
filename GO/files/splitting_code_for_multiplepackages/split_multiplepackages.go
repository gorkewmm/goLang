package main

import (
	// "errors"
	"fmt"
	// "os"
	// "strconv"
	"split_multiple/fileops"

	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.ReadFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
	}

	fmt.Println("Go bankacılıgına hoşgeldiniz")
	fmt.Println(randomdata.Address())
	for {
		fmt.Println("Hesap bakiyeni öğren: 1")
		fmt.Println("Para yatır: 2")
		fmt.Println("Para çek: 3")
		fmt.Println("Çıkış yap: 4")

		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Hesap bakiyeniz:", accountBalance)
			fileops.ReadFromFile(accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		} else if choice == 4 {
			println("Görüşürüz")
			return
		} else {
			fmt.Println("Geçersiz miktar")
		}
	}
}

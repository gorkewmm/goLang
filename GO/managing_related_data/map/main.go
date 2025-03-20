package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	websites := map[string]string{
		"Google":              "https://google.com", //key-value ilişkisi
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites["Amazon Web Services"])
	fmt.Println(websites)

	websites["LinkedIn"] = "https://linkedin.com" //dinamik olarak ekleme yaptık
	fmt.Println(websites)

	delete(websites, "LinkedIn")
	fmt.Println(websites)

	var courseRates floatMap
	courseRates = make(floatMap, 3)

	courseRates["go"] = 4.6
	courseRates["react"] = 4.8
	courseRates["python"] = 4.5

	courseRates.output()

}

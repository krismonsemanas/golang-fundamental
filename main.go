package main

import "fmt"

func main() {
	// deklarasi variabel
	var nama string = "Krismon"
	umur := 22

	if umur >= 20 {
		fmt.Println("Bentar lagi nikah woi....")
	}

	kalimat := "Cepat wisuda ye"

	for idx, huruf := range kalimat {
		if idx%2 == 0 {
			fmt.Println("Huruf index genap ", string(huruf))
		}
	}

	for idx, huruf := range kalimat {
		hurufVokal := string(huruf)

		// cara ke 1
		/* if hurufVokal == "a" || hurufVokal == "i" || hurufVokal == "u" || hurufVokal == "e" || hurufVokal == "o" {
			fmt.Println("Huruf vokal pada index ke-", idx)
		} */

		switch hurufVokal {
		case "a", "i", "u", "e", "o":
			fmt.Println("Index :", idx, "Huruf : ", hurufVokal)
		}
	}

	fmt.Println("Hello World, nama saya ", nama)
	fmt.Println(umur)
}

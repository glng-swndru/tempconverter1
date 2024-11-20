package main

import "fmt"

func main() {
	fmt.Println("Masukkan kalkulator suhu yang ingin dipakai")
	fmt.Println("1. Celcius ke Fahrenheit")
	fmt.Println("2. Celcius ke Kelvin")
	fmt.Println("3. Fahrenheit ke Celcius")
	fmt.Println("4. Fahrenheit ke Kelvin")
	fmt.Println("5. Kelvin ke Celcius")
	fmt.Println("6. Kelvin ke Fahrenheit")
	fmt.Println("Masukkan pilihan anda:")

	var pilihan int
	fmt.Scanf("%d\n", &pilihan)
	for pilihan < 1 || pilihan > 6 {
		fmt.Println("Kalkulator tidak tersedia, masukkan kalkulator pilihan anda: ")
		fmt.Scanf("%d\n", &pilihan)
	}

	fmt.Println("Masukkan suhu yang ingin dikonversi:")
	var suhu float64
	fmt.Scanf("%f\n", &suhu)

	var suhuAkhir float64
	if pilihan == 1 {
		suhuAkhir = CelciusToFahrenheit(suhu)
	} else if pilihan == 2 {
		suhuAkhir = CelciusToKelvin(suhu)
	} else if pilihan == 3 {
		suhuAkhir = FahrenheitToCelcius(suhu)
	} else if pilihan == 4 {
		suhuAkhir = FahrenheitToKelvin(suhu)
	} else if pilihan == 5 {
		suhuAkhir = KelvinToCelcius(suhu)
	} else {
		suhuAkhir = KelvinToFahrenheit(suhu)
	}

	fmt.Printf("Suhu akhir hasil konversi adalah: %.2f\n", suhuAkhir)
}

func CelciusToFahrenheit(suhuCelcius float64) float64 {
	return (9.0/5.0)*suhuCelcius + 32
}

func CelciusToKelvin(suhuCelcius float64) float64 {
	return suhuCelcius + 273.15
}

func FahrenheitToCelcius(suhuFahrenheit float64) float64 {
	return (suhuFahrenheit - 32) * 5.0 / 9.0
}

func FahrenheitToKelvin(suhuFahrenheit float64) float64 {
	return (suhuFahrenheit + 459.67) * (5.0 / 9.0)
}

func KelvinToCelcius(suhuKelvin float64) float64 {
	return suhuKelvin - 273.15
}

func KelvinToFahrenheit(suhuKelvin float64) float64 {
	return (suhuKelvin * (9.0 / 5.0)) - 459.67
}

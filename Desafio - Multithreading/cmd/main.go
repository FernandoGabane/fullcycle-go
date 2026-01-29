package main

import (
	"fmt"
	"fullcycle/go/exercise-2/internal/service"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <CEP>")
		return
	}

	cep := os.Args[1]
	svc := service.NewCepService()

	result, err := svc.GetFastest(cep)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("API:", result.Source)
	fmt.Println("CEP:", result.CEP)
	fmt.Println("Street:", result.Street)
	fmt.Println("City:", result.City)
	fmt.Println("State:", result.State)
}

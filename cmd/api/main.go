package main

import (
	"fmt"
	"log"

	"github.com/Modular-Company/modular-bar/internal/infra"
)

func main() {
	fmt.Println("Olá Mundo")
	if err := infra.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

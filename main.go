package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	PhotoApi = "https://api.pexels.com/v1"
	VideoApi = "https://api.pexels.com/videos"
)

func main() {
	godotenv.Load()

	token := os.Getenv("PEXELS_TOKEN")
	fmt.Println("Hello World!")

	c := NewClient(token)
}

func NewClient() {

}

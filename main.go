package main

import (
	"log"

	"netease-music-mygo/internal/app"
)

func main() {
	useBundledLibs()
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

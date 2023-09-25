package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	openCreateFile("./my_cards")

}

// Opens or creates a file and defers the close
// till the function finishes but before it returns
func openCreateFile(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Println("File open error:", err.Error())
	} else {
		log.Printf("\nFile %s opened.", f.Name())
		defer f.Close()
	}

	c, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("File could not be created:", err.Error())
		//Without defer to f.Close(), the file handler f will remain open
	} else {
		slog.Info("File created.", "name", c.Name())
		defer c.Close()
	}

}

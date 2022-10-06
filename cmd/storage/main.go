package main

import (
	"fmt"
	"github.com/istrel/storage/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("main package", st)
	file, err := st.Upload("test.txt", []byte("Hello"))

	if err != nil {
		log.Fatal(err)
	}

	file1, err1 := st.GetByID(file.ID)

	if err != nil {
		log.Fatal(err1)
	}

	fmt.Println("Uploaded", file1)
}

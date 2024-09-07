package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const PORT = "6900"
const csvDir = "csv_data"
const todoCSV = "todo.csv"

// read from csv
// write to csv

// gonna create a server using the basic net/http with a get post put and delete
// routes, for the database I'm going to store everything on a csv

type Todo struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Todos []Todo

func initDirectory(s string) error {
	dir, err := os.Stat(s)
	if os.IsNotExist(err) {
		return os.MkdirAll(s, 0700)
	}
	if err != nil {
		return err
	}
	if !dir.IsDir() {
		return fmt.Errorf("%s already exists but is not a directory", s)
	}
	return nil
}

func initCSV(s string) error {
	csvPath := csvDir + "/" + todoCSV

	

	currCSV, err := os.ReadFile(csvPath)
	err != nil{
		// TODO CHECK THE FILE EXIST
		return os.WriteFile(csvPath,[], 0600 )

	}


	return nil

}

func readFromCSV() {

}

func writeToCSV() {

}

func getTodos(w http.ResponseWriter, r *http.Request) {

	// simply return an array
	fmt.Fprintf(w, "hello world")
}

func main() {
	err := initDirectory(csvDir)
	if err != nil {
		log.Println("Cannot init the dir", err)
		return
	}

	http.HandleFunc("/", getTodos)
	log.Println("Running on port ", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

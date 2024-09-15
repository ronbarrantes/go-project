package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"ronb.co/project/database"
	"ronb.co/project/utils"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	utils.LoadEnvs()
	database.InitDB()

	PORT := os.Getenv("PORT")
	http.HandleFunc("/", HelloWorld)

	http.HandleFunc("/qr", utils.PrintQRCode)

	log.Println("Running on port", PORT)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+PORT, nil))
}

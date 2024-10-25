package main

func main() {
	PORT := "8000"
	server := Server(PORT)

	server.Run()
}

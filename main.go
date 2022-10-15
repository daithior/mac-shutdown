package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

// TODO Add a logger
// write the logs to a file
// Where ?

func executeShellCommand() string {
	cmd := exec.Command("uname")
	//cmd := exec.Command("shutdown", "-h", "now")
	stout, err := cmd.Output()

	if err != nil {
		fmt.Println("Error running shell command")
	}
	log.Print("Shell command ran successfully")
	return string(stout)
}

func handleShutdown(writer http.ResponseWriter, request *http.Request) {
	log.Print("/shutdown was called. Executing shell")
	// What happens if this is run in a Goroutine ?
	executeShellCommand()
	log.Print("Writing HTTP response")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "shutdown complete")
}

func main() {
	log.Print("Starting")
	http.HandleFunc("/shutdown", handleShutdown)
	http.ListenAndServe(":8080", nil)
}

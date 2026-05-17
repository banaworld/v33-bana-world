package main

import (
	"fmt"
	"net/http"
	"strings"
	"os"
)

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/webhook", webhook)

	port := os.Getenv("PORT")
	if port == "" {
		port = "7860"
	}

	fmt.Println("V33 running on :", port)
	http.ListenAndServe(":"+port, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "V33 ACTIVE - Bana World System Running")
}

func webhook(w http.ResponseWriter, r *http.Request) {

	msg := r.URL.Query().Get("msg")
	phone := r.URL.Query().Get("phone")

	intent := detect(msg)

	result := assign(intent)

	fmt.Fprintf(w, "Phone: %s | Task: %s", phone, result)
}

func detect(msg string) string {
	msg = strings.ToLower(msg)

	if strings.Contains(msg, "plumber") {
		return "plumber"
	}
	if strings.Contains(msg, "electrician") {
		return "electrician"
	}
	if strings.Contains(msg, "delivery") {
		return "delivery"
	}
	return "general"
}

func assign(intent string) string {
	switch intent {
	case "plumber":
		return "Worker-PL-001"
	case "electrician":
		return "Worker-EL-002"
	case "delivery":
		return "Worker-DL-003"
	default:
		return "AI-FALLBACK"
	}
}

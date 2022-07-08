package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func subscribeUser(w http.ResponseWriter, r *http.Request) {
	// Fetch Buttondown key
	key := os.Getenv("BUTTONDOWN_KEY")
	if key == "" {
		http.ServeFile(w, r, "web/failed.html")
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("[Error] Missing Buttondown API key")
		return
	}
	apiKey := []string{"Token",key}

	// Fetch email from form data
	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		log.Printf("[Error] Didn't receive proper data")
		http.ServeFile(w, r, "web/failed.html")
		return
	}
	r.ParseForm()
	email := r.FormValue("email")

	// Send email via POST request
	reqUrl := []string{
		"https://api.buttondown.email/v1/subscribers?email=",
		email,
	}
	req, err := http.NewRequest("POST", strings.Join(reqUrl, ""), nil)
	if err != nil {
		http.ServeFile(w, r, "web/failed.html")
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("[Error] Failed to make request to Buttondown", err)
		return
	}
	req.Header.Set("Authorization", strings.Join(apiKey, " "))
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	res, err := client.Do(req)
	if err != nil {
		http.ServeFile(w, r, "web/failed.html")
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("[Error] Failed to make request to Buttondown", err)
		return
	}
	res.Body.Close()
	http.Redirect(w, r, "/subscribed.html", http.StatusFound)
}

func main() {
	// Replace the below web/index.html, web/unsub.html, and web/subscribed.html
	// with other files in this directory. You can also comment out these lines
	// if you do not want to or need to serve them.
	http.Handle("/", http.FileServer(http.Dir("web/")))

	// This handles POST requests containing multipart/form-data. For an
	// example, see the included web/ directory.
	http.HandleFunc("/subscribe", func(w http.ResponseWriter, r *http.Request) { subscribeUser(w, r) })

	log.Fatal(http.ListenAndServe(":8080", nil))
}

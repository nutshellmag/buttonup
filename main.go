package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func cors(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Permission-Policy", "interest-cohort=(), browsing-topics=(), join-ad-interest-group=(), run-ad-auction=()")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("Vary", "Origin")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000;includeSubDomains")
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		// Comment the line below if being used in production outside of Nutshell
		w.Header().Set("Content-Security-Policy", "default-src 'none'; style-src 'self'; img-src 'self' https://cdn.nutshellmag.com; font-src https://cdn.nutshellmag.com; form-action 'self'; upgrade-insecure-requests; block-all-mixed-content; object-src 'none'; connect-src 'none'; base-uri 'self'; frame-ancestors 'none'")

		fs.ServeHTTP(w, r)
	}
}

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

	// Set headers for security reason
	w.Header().Add("Permission-Policy", "interest-cohort=(), browsing-topics=(), join-ad-interest-group=(), run-ad-auction=()")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("Referrer-Policy", "no-referrer")
	w.Header().Set("Vary", "Origin")
	w.Header().Set("Strict-Transport-Security", "max-age=31536000;includeSubDomains")
	//w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
	//w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
	// Comment the line below if being used in production outside of Nutshell
	//w.Header().Set("Content-Security-Policy", "default-src 'none'; style-src 'self'; img-src 'self' https://cdn.nutshellmag.com; font-src https://cdn.nutshellmag.com; form-action 'self'; upgrade-insecure-requests; block-all-mixed-content; object-src 'none'; connect-src 'none'; base-uri 'self'; frame-ancestors 'none'")

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
	http.Redirect(w, r, "/subscribed", http.StatusFound)
}

func main() {
	// Replace the below web/index.html, web/unsub.html, and web/subscribed.html
	// with other files in this directory. You can also comment out these lines
	// if you do not want to or need to serve them.
	http.Handle("/", cors(http.FileServer(http.Dir("web/"))))

	// This handles POST requests containing multipart/form-data. For an
	// example, see the included web/ directory.
	http.HandleFunc("/subscribe", func(w http.ResponseWriter, r *http.Request) { subscribeUser(w, r) })

	log.Fatal(http.ListenAndServe(":8080", nil))
}

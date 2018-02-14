package main

import (
	"log"
	"net/http"
	"time"

	"github.com/goller/traces"
)

func Tier2() {
	http.HandleFunc("/tier2", func(w http.ResponseWriter, r *http.Request) {
		span := traces.NewSpanFromRequest(r, "tier2")
		defer span.Finish()

		time.Sleep(20 * time.Millisecond)
		req, _ := http.NewRequest("GET", "http://localhost:8081/tier3", nil)
		traces.SendSpan(req, span)
		http.DefaultClient.Do(req)

	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	Tier2()
}

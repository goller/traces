package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/goller/traces"
)

// LongSleep sleeps between 1 and 1000 milliseconds
func LongSleep() {
	r := rand.Intn(1000)
	time.Sleep(time.Duration(r) * time.Millisecond)
}

// Tier3 waits between 1 and 1000 milliseconds before responding
func Tier3() {
	http.HandleFunc("/tier3", func(w http.ResponseWriter, r *http.Request) {
		span := traces.NewSpanFromRequest(r, "tier3")
		defer span.Finish()

		LongSleep()
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	Tier3()
}

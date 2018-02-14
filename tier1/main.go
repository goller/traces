package main

import (
	"net/http"
	"time"

	"github.com/goller/traces"
)

// Tier1 calls server on tier2
func Tier1() {
	span := traces.NewSpan("tier1")
	defer span.Finish()

	req, _ := http.NewRequest("GET", "http://localhost:8080/tier2", nil)
	traces.SendSpan(req, span)

	http.DefaultClient.Do(req)
}

func main() {
	// Create 3 traces through the system
	for i := 0; i < 3; i++ {
		Tier1()
		time.Sleep(5 * time.Second)
	}
}

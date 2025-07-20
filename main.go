package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/starfederation/datastar-go/datastar"
)

//go:embed index.html
var indexHTML []byte

const port = 1337

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(indexHTML)
	})

	r.Get("/forever", func(w http.ResponseWriter, r *http.Request) {
		sse := datastar.NewSSE(w, r)

		var counter = 0
		for {
			counter += 1

			time.Sleep(500 * time.Millisecond)
			new_signals, _ := json.Marshal(map[string]int{"counter": counter})
			sse.PatchSignals(new_signals)

			time.Sleep(500 * time.Millisecond)
			sse.PatchElements(`<div id="target-element-direct" class="large-number" data-text="$counter"></div>`)

			time.Sleep(500 * time.Millisecond)
			sse.PatchElements(`<div id="target-element-nested"><span data-text="$counter" class="large-number">$counter NOT RESOLVED</span></div>`)

		}
	})

	log.Printf("Starting server on http://localhost:%d", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"net/http"
	"time"
)

func singleChunkHandler(w http.ResponseWriter, r *http.Request) {
	flusher, _ := w.(http.Flusher)

	w.Header().Set("Transfer-Encoding", "chunked")
	fmt.Fprintf(w, "This is a single chunk response\n")
	flusher.Flush()
}

func multipleChunksHandler(w http.ResponseWriter, r *http.Request) {
	flusher, _ := w.(http.Flusher)

	w.Header().Set("Transfer-Encoding", "chunked")
	for i := 1; i <= 5; i++ {
		_, err := fmt.Fprintf(w, "Chunk %d\n", i)
		if err != nil {
			fmt.Errorf("failed writing to ResponseWriter: %v", err)
		}
		flusher.Flush()
		time.Sleep(6 * time.Second)
	}
}

func main() {
	http.HandleFunc("/single", singleChunkHandler)
	http.HandleFunc("/multiple", multipleChunksHandler)

	fmt.Println("Server is starting...")
	http.ListenAndServe(":8080", nil)
}

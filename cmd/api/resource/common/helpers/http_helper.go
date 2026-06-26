package helpers

import (
	"log"
	"net/http"
)

func SafeWrite(w http.ResponseWriter, bytes []byte) {
	_, err := w.Write(bytes)
	if err != nil {
		log.Println("Write failed:", err)
	}
}

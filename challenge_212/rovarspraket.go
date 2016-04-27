package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type input struct {
	Message string
}

func main() {
	http.HandleFunc("/rovarspraket", func(w http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		var i input
		err := decoder.Decode(&i)
		if err != nil {
			fmt.Fprintf(w, "Parse error!\n")
		} else {
			fmt.Fprintf(w, encodeText(i.Message)+"\n")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func encodeText(input string) string {
	consonants := "bcdfghjklmnpqrstvwxz"
	encoded := ""
	for _, c := range input {
		if strings.Contains(consonants, strings.ToLower(string(c))) {
			encoded += string(c) + "o" + strings.ToLower(string(c))
		} else {
			encoded += string(c)
		}
	}

	return encoded
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Resume struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Location string `json:"location"`
	Email    string `json:"email"`
	Summary  string `json:"summary"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resume := Resume{
			Name:     "Eric Thomas",
			Title:    "Software Developer",
			Location: "Corner Brook, NL, CA",
			Email:    "eric@ericthomas.ca",
			Summary:  "I am a dedicated Software Developer with expertise in designing, installing, testing, and maintaining software systems. My strong analytical skills have been honed through diverse work experiences. Proficient in various platforms, languages, systems analysis, and security, I am well-versed in the latest cutting-edge development tools and procedures. Whether self-managing independent projects or collaborating in a team, I consistently deliver high-quality results.",

		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(resume)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
	})

	fmt.Println("Server listening on :7777...")
	http.ListenAndServe(":7777", nil)
}

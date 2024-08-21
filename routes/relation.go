package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	ctypes "groupie-tracker/types"

	"groupie-tracker/requests"

)

func Relation(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/relation/")
    if id == "" {
        http.Error(w, "Page not found", http.StatusNotFound)
        return
    }

	relationUrl := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%s", id)
	var relation ctypes.Relation

	jsonRelationData, status, found := requests.Get(relationUrl)
	if found && status == "200 OK" {
		err := json.Unmarshal([]byte(jsonRelationData), &relation)
		if err != nil {
			fmt.Println("Error unmarshalling location JSON:", err)
			return
		}

		data := struct {
			Title    string
			Relation map[string][]string
		}{
			Title:    "Relation",
			Relation: relation.Concerts, // Updated to use Concerts
		}

		templ, err := template.ParseFiles("templates/relation.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	t := template.Must(templ, err)
		t.Execute(w, data)
	} else {
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}
package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"../model"
	"../views"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// POST -> Create data
		if r.Method == http.MethodPost {

			data := views.PostRequest{}
			// read json data from request body and decode to data object
			err := decodeJSONBody(w, r, &data)
			if err != nil {
				var mr *malformedRequest
				if errors.As(err, &mr) {
					http.Error(w, mr.msg, mr.status)
				} else {
					log.Println(err.Error())
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
				return
			}

			// fmt.Fprintf(w, "TASK: %+v", data)
			fmt.Println("Nome:", data.Name, "Tarefa:", data.Task)

			err = model.InsertTask(data.Name, data.Task) // insert into database
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}

			w.WriteHeader(http.StatusCreated)
			// write
			json.NewEncoder(w).Encode(data)
			fmt.Println("Tarefa inserida com sucesso!")
		}
	}
}

// GET -> List data by name
func listByName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			data, err := model.ReadByName(name)

			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}

// List all data
func list() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data, err := model.ReadAll()
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}

func delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodDelete {
			// fmt.Println(r.URL.Path[8:])
			name := r.URL.Path[8:]

			if err := model.DeleteTask(name); err != nil {
				w.Write([]byte(err.Error()))
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct{
			Status string `json:status`
		}{"Item deletado"})
	}
}
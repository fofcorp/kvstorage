package httpserv

import (
	"encoding/json"
	"net/http"

	"github.com/fofcorp/kvstorage/src/storage"
)

// GetHandler ...
func GetHandler(store storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		key := r.URL.Query().Get("key")
		result, err := store.Get(key)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}
}

// PutHandler ...
func PutHandler(store storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var body map[string]string
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		err := decoder.Decode(&body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		key, ok := body["key"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		value, ok := body["value"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = store.Put(key, value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}

// DeleteHandler ...
func DeleteHandler(store storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		key := r.URL.Query().Get("key")
		err := store.Delete(key)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}

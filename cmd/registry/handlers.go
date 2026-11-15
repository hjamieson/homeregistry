package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
	"uuid"

	"github.com/hjamieson/homeregistry/internal/model"
	"github.com/hjamieson/homeregistry/internal/stores"
)

func (a *app) handlePost(w http.ResponseWriter, r *http.Request) {
	maxBytes := int64(2048)
	readCloser := http.MaxBytesReader(w, r.Body, maxBytes)

	decoder := json.NewDecoder(readCloser)
	device := model.Device{}
	err := decoder.Decode(&device)
	if err != nil {
		log.Printf("decode failed: %v", err)
		var tooBig *http.MaxBytesError
		switch {
		case errors.As(err, &tooBig):
			http.Error(w, "Request Too Large", http.StatusRequestEntityTooLarge)
		default:
			http.Error(w, "Malformed Input", http.StatusBadRequest)
		}
		return
	}
	// it is an error if Id is present; we must assign it here!
	if device.Id != "" {
		msg := fmt.Sprintf("Id is not allowed on POST: %q", device.Id)
		log.Print(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// set fields we manage:
	device.Id = uuid.NewV7().String()
	device.LastInfo = time.Now()

	// store it:
	err = a.devices.Put(device)
	if err != nil {
		log.Printf("PUT error: %v", err)
		http.Error(w, "other error occurred", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(device)
	if err != nil {
		log.Printf("error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", "/devices/"+device.Id)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
	log.Printf("POSTED device: %#v", device)
}

func (a *app) handleList(w http.ResponseWriter, r *http.Request) {
	devices, err := a.devices.List()
	if err != nil {
		log.Printf("error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	result, err := json.Marshal(devices)
	if err != nil {
		log.Printf("error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(result)
}

func (a *app) handleGet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	device, err := a.devices.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, stores.ErrNotFound):
			log.Printf("Id %q not found", id)
			http.NotFound(w, r)
		default:
			log.Printf("GET id=%q failed: %v", id, err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}
	response, err := json.Marshal(device)
	if err != nil {
		log.Printf("data for %q is corrupt: %v", id, err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(response)
}

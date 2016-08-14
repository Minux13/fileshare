package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Archive struct {
	SavePath string
	Name     string
	Key      string
	Expire   int
}

var ArchiveStore map[string]Archive

// Given an string id searches for a file and delivers it to the client
func DownloadArchiveHandler(w http.ResponseWriter, r *http.Request) {
}

// Updates the file name and expire date
func UpdateArchiveHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := CheckAuth(r); !ok {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "{\"error\": \"Unauthorized\"}")
		return
	}
}

// Deletes a file from the server
func DeleteArchiveHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := CheckAuth(r); !ok {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "{\"error\": \"Unauthorized\"}")
		return
	}
}

// Lists all the files that the server has available to download
func ListArchiveHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := CheckAuth(r); !ok {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "{\"error\": \"Unauthorized\"}")
		return
	}

	resp := map[string][]Archive{
		"archives": make([]Archive, 0),
	}

	for _, archive := range ArchiveStore {
		resp["archives"] = append(resp["archives"], archive)
	}

	json.NewEncoder(w).Encode(resp)
}

// Uploads a file to the server, returns the status and expire date
func AddArchiveHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := CheckAuth(r); !ok {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "{\"error\": \"Unauthorized\"}")
		return
	}
}
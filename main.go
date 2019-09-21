package main

import (
	"golang-inter/handler"
	"golang-inter/model"
	"fmt"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

var username, password, host, namaDB, defaultDB string
var db *sql.DB
var err error

func init() {
	username = "root"
	password = ""
	host = "localhost"
	namaDB = "gunadarma"
	defaultDB = "mysql"
}

func main() {
	db, err = model.Connect(username, password, host, namaDB)
	if err != nil {
		return
	}
	defer db.Close()

	http.HandleFunc("/api/", handler.Api)

	http.HandleFunc("/mahasiswa", mahasiswa)
	http.HandleFunc("/mahasiswas", mahasiswas)
	
	http.HandleFunc("/matakuliah", matakuliah)
	http.HandleFunc("/matakuliahs", matakuliahs)
	
	http.HandleFunc("/nilai", nilai)
	http.HandleFunc("/nilais", nilais)	

	http.HandleFunc("/login", login)

	log.Println("localhost : 8000")
	http.ListenAndServe(":8000", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		fmt.Fprint(w, "Benar")
	default:
		fmt.Fprint(w, "Salah")
	}
}

func nilai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	npm := r.URL.Query()["npm"][0]
	data, err := model.GetNilai(db, npm)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonData, _ := json.Marshal(data)
	w.Write(jsonData)
}

func nilais(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := model.GetAllNilai(db)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonData, _ := json.Marshal(data)
	w.Write(jsonData)
}

func mahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	npm := r.URL.Query()["npm"][0]
	data, err := model.GetMahasiswa(db, npm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonData, _ := json.Marshal(data)
	w.Write(jsonData)
}

func mahasiswas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := model.GetAllMahasiswa(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonData, _ := json.Marshal(data)
	w.Write(jsonData)
}

func matakuliah(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	kode := r.URL.Query()["kode"][0]
	data, err := model.GetMatkul(db, kode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonData, _ := json.Marshal(data)
	w.Write(jsonData)
}

func matakuliahs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := model.GetAllMatkul(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonData, _ := json.Marshal(data)
	w.Write(jsonData)
}

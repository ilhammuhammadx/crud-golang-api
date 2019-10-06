package handler

import (
	"crud-golang-api/model"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
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

func API(w http.ResponseWriter, r *http.Request) {
	db, err = model.Connect(username, password, host, namaDB)
	if err != nil {
		return
	}
	defer db.Close()

	w.Header().Set("Content-Type", "text-html; charset=utf-8; application/json")
	dataURL := strings.Split(fmt.Sprintf("%s", r.URL.Path), "/")
	switch dataURL[2] {
	case "mahasiswa":
		switch r.Method {
		case http.MethodGet:
			HandlerMahasiswaGet(w, r)
		case http.MethodPost:
			HandlerMahasiswaPost(w, r)
		case http.MethodPut:
			HandlerMahasiswaPut(w, r)
		case http.MethodDelete:
			HandlerMahasiswaDelete(w, r)
		default:
			w.Write([]byte("method tidak ditemukan"))
		}
	case "matkul":
		switch r.Method {
		case "GET":
			HandlerMatkulGet(w, r)
		case "POST":
			HandlerMatkulPost(w, r)
		case "PUT":
			HandlerMatkulPut(w, r)
		case "DELETE":
			HandlerMatkulDelete(w, r)
		default:
			w.Write([]byte("Method tidak ditemukan"))
		}
	case "nilai":
		switch r.Method {
		case "GET":
			HandlerNilaiGet(w, r)
		case "POST":
			HandlerNilaiPost(w, r)
		case "PUT":
			HandlerNilaiPut(w, r)
		case "DELETE":
			HandlerNilaiDelete(w, r)
		default:
			w.Write([]byte("Method tidak ditemukan"))
		}
	default:
		w.Write([]byte("request tidak ditemukan"))
	}
}
package handler

import (
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

func LastIndex(r *http.Request) string {
	dataURL := strings.Split(fmt.Sprintf("%s", r.URL.Path), "/")
	lastIndex := dataURL[len(dataURL)-1]
	return lastIndex
}

func Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text-htmll; charset=utf-8; application/json")
	dataURL := strings.Split(fmt.Sprintf("%s", r.URL.Path), "/")
	switch dataURL[2] {
	case "mahasiswa":
		switch r.Method {
		case "GET":
			HandlerMahasiswaGet(w, r)
		case "POST":
			HandlerMahasiswaPost(w, r)
		// case http.MethodPut:
		// 	HandlerMahasiswaPut(w, r)
		// case http.MethodDelete:
		// 	HandlerMahasiswaDelete(w, r)
		default:
			w.Write([]byte("method tidak ditemukan"))
		}
	default:
		w.Write([]byte("request tidak ditemukan"))
	}
}
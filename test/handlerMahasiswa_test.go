package test

import (
// 	"golang-inter/handler"
// 	"golang-inter/model"
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/http/httptest"
// 	"net/url"
// 	"testing"
)

// func TestMahasiswaHandler(t *testing.T) {
// 	db, err := initDatabase()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	defer db.Close()

// 	var dataInsertMhs = []model.Mahasiswa {
// 		model.Mahasiswa{
// 			NPM:   "12345678",
// 			Nama:  "Budi Doremi",
// 			Kelas: "3KA20",
// 		},
// 		model.Mahasiswa{
// 			NPM:   "19283746",
// 			Nama:  "Doremi Budi",
// 			Kelas: "4KA20",
// 		},
// 		model.Mahasiswa{
// 			NPM:   "44444444",
// 			Nama:  "DoBud",
// 			Kelas: "4KA21",
// 		},
// 	}

// 	webHandler := http.HandlerFunc(handler.Api)
	
// }
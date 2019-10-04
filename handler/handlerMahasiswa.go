package handler

import (
	"golang-inter/model"
	"encoding/json"	
	// "io/ioutil"	
	"net/http"
)

func HandlerMahasiswaGet(w http.ResponseWriter, r *http.Request) {
	lastIndex := LastIndex(r)

	if lastIndex == "mahasiswa" {
		npm := r.URL.Query()["npm"][0]
		data, err := model.GetMahasiswa(db, npm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		jsonData, _ := json.Marshal(data)
		w.Write(jsonData)		
	}
}

func HandlerMahasiswaPost(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	w.Write([]byte("Halo " + username))
}
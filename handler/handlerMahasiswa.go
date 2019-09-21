package handler

import (
	"golang-inter/model"
	"encoding/json"	
	"io/ioutil"
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
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var data model.Mahasiswa
	if err = json.Unmarshal(body, &data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = data.Insert(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(jsonData)
}
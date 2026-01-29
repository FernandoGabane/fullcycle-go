package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type Request struct {
	CEP string `json:"cep"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	var req Request
	json.NewDecoder(r.Body).Decode(&req)
	if len(req.CEP) != 8 {
		http.Error(w, "invalid zipcode", 422)
		return
	}
	body, _ := json.Marshal(req)
	resp, _ := http.Post(os.Getenv("SERVICE_B_URL")+"/weather", "application/json", bytes.NewBuffer(body))
	defer resp.Body.Close()
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

package handler

import (
	"encoding/json"
	"net/http"
	"service-b/internal/service"
)

type Req struct {
	CEP string `json:"cep"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	var req Req
	json.NewDecoder(r.Body).Decode(&req)
	if len(req.CEP) != 8 {
		http.Error(w, "invalid zipcode", 422)
		return
	}
	res, err := service.GetWeather(r.Context(), req.CEP)
	if err != nil {
		http.Error(w, "can not find zipcode", 404)
		return
	}
	json.NewEncoder(w).Encode(res)
}

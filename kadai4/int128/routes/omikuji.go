package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gopherdojo/dojo2/kadai4/int128/omikuji"
)

// OmikujiHandler is a HTTP handler for GET omikuji
type OmikujiHandler struct {
	service omikuji.Service
}

// OmikujiGetResponse is a GET response of OmikujiHandler.
type OmikujiGetResponse struct {
	Description string `json:"description"`
	Value       int    `json:"value"`
}

func (h *OmikujiHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	switch req.Method {
	case "GET":
		w.Header().Add("Content-Type", "application/json")
		o := h.service.Hiku()
		e := json.NewEncoder(w)
		res := OmikujiGetResponse{o.String(), int(o)}
		if err := e.Encode(res); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

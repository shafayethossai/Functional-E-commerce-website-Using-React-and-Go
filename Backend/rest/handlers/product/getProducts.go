package product

import (
	"first-program/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	// HandleCors(w)

	// //HandlePreFlightReq(w, r)

	// if r.Method == "OPTIONS" {
	// 	w.WriteHeader(200)
	// 	return
	// }

	// if r.Method != "GET" { // r.method = post, put, match, delete
	// 	http.Error(w, "Please use GET method", 400)
	// 	return
	// }
	productList, err := h.svc.List()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
	}
	util.SendData(w, http.StatusOK, productList)
}

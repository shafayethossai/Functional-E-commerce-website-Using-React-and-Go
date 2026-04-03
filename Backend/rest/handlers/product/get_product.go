package product

import (
	"first-program/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)

	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid req body")
		return
	}
	product, err := h.svc.Get(pID)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Invalid req body")
		return
	}
	if product == nil {
		util.SendError(w, http.StatusNotFound, "Invalid req body")
		return
	}

	util.SendData(w, http.StatusOK, product)
}

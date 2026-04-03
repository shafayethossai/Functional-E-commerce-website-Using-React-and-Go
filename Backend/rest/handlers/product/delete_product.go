package product

import (
	"first-program/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProducts(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)

	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	err = h.svc.Delete(pID)
	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	util.SendData(w, http.StatusOK, "Successfully Deleted product")
}

package product

import (
	"encoding/json"
	"first-program/domain"
	"first-program/util"
	"fmt"
	"net/http"
	"strconv"
)

type ReqUpdateProducts struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h *Handler) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)

	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var req ReqUpdateProducts
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json", 400)
		util.SendError(w, http.StatusBadRequest, "Invalid req body")
		return
	}
	_, err = h.svc.Update(domain.Product{
		ID:          pID,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      r.RequestURI,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server error")
		return
	}

	util.SendData(w, http.StatusOK, "Successfully updated product")
}

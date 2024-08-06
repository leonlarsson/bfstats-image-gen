package handlers

import (
	"encoding/json"
	"image/png"
	"net/http"

	"github.com/leonlarsson/bfstats-image-gen/canvas"
	"github.com/leonlarsson/bfstats-image-gen/create"
	"github.com/leonlarsson/bfstats-image-gen/structs"
)

func BF2042Handler(w http.ResponseWriter, r *http.Request) {

	var data structs.BaseData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c, _ := create.CreateBF2042Image(data)
	w.Header().Set("Content-Type", "image/png")
	img := canvas.CanvasToImage(c)
	png.Encode(w, img)
}

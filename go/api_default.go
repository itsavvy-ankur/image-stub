package swagger

import (
	"encoding/json"
	"log"
	"net/http"
)

func Getimage(w http.ResponseWriter, r *http.Request) {
	requestId := r.Header.Get("RequestId")
	imageUrl := r.Header.Get("ImageUrl")
	itemBarCode := r.Header.Get("ItemBarCode")

	imageBytes, err := IC.GetImageFromConfig(requestId)
	if err != nil {

		sendError(w, requestId, imageUrl, itemBarCode, err)

	} else {
		w.Header().Set("Content-Type", "image/jpeg")
		w.WriteHeader(http.StatusOK)

		w.Write(imageBytes)
	}

}

type ErrorDescription struct {
	Description string `json:"description"`
}

func sendError(w http.ResponseWriter, requestId string, imageUrl string, itemBarCode string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	modelError := ModelError{}
	modelError.RequestId = requestId
	modelError.ImageUrl = imageUrl
	modelError.ItemBarcode = itemBarCode
	errorDescription := ErrorDescription{
		Description: err.Error(),
	}
	modelError.Error_ = errorDescription
	//modelError.Error_ = fmt.Sprintf("{\"description\":\"%s\"}", err.Error())

	errorBytes, err := json.Marshal(modelError)
	if err != nil {
		log.Panic("Error marshalling json")
	}

	w.Write(errorBytes)

}

package swagger

type ModelError struct {

	// The URL for the image
	ImageUrl string `json:"imageUrl,omitempty"`

	// The Request Identifier for correlation.
	RequestId string `json:"requestId,omitempty"`

	// The Item Barcode for correlation.
	ItemBarcode string `json:"itemBarcode,omitempty"`

	// Error message.
	Error_ interface{} `json:"error"`
}

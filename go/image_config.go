package swagger

type ImageConfig []struct {
	RequestID   string `yaml:"requestId"`
	ImageURL    string `yaml:"imageUrl"`
	ItemBarcode string `yaml:"itemBarcode"`
}

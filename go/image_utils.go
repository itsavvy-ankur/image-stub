package swagger

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (ic *ImageConfig) GetImageFromConfig(requestId string) ([]byte, error) {
	fmt.Printf("ic: %v\n", *ic)

	for _, x := range *ic {
		if x.RequestID == requestId {
			if strings.HasPrefix(x.ImageURL, "http") {
				// don't worry about errors
				response, e := http.Get(x.ImageURL)
				if e != nil {
					return nil, fmt.Errorf("image not found")
				}
				httpBytes, err := ioutil.ReadAll(response.Body)
				if err != nil {
					return nil, fmt.Errorf("image not found")
				}
				return httpBytes, nil

			} else {
				fileBytes, err := ioutil.ReadFile(x.ImageURL)
				if err != nil {
					return nil, fmt.Errorf("image not found")
				}
				return fileBytes, nil
			}

		}
	}
	//fmt.Printf("RequestId : %s | ImageUrl : %s | Item Barcode : %s", "", ic.ImageUrl, ic.ItemBarcode)
	return nil, fmt.Errorf("Image Not found")
}

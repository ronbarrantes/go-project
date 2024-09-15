package utils

import (
	"fmt"
	"net/http"

	"github.com/skip2/go-qrcode"
)

func PrintQRCode(w http.ResponseWriter, r *http.Request) {
	var png []byte
	png, err := qrcode.Encode("https://ronb.co", qrcode.Medium, 256)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(png)
}

package QRHandler

import qrcode "github.com/skip2/go-qrcode"

func Generate_QR(text string) ([]byte, error) {
	qr, err := qrcode.Encode(text, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return qr, nil
}

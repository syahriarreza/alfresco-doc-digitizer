package engine

import (
	"fmt"
	"image"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/novalagung/gubrak"
	"github.com/spf13/viper"
)

// PDFExtract result of pdf extraction
type PDFExtract struct {
	Pages  []image.Image
	QRCode Qrcodes
}

// ExtractPDF ExtractPDF
func ExtractPDF(filename string) ([]PDFExtract, error) {
	allPDFs := []PDFExtract{}
	onePDF := PDFExtract{}
	qrCodeIDs := []string{}
	qrDecoder := qrcode.NewQRCodeReader()

	// scanning pdf and extract images
	doc, e := fitz.New(filepath.Join(viper.GetString("folder.scan"), filename))
	if e != nil {
		return nil, fmt.Errorf("error get pdf file '%s' : %s", filename, e.Error())
	}
	defer doc.Close()

	for i := 0; i < doc.NumPage(); i++ {
		img, e := doc.Image(i)
		if e != nil {
			return nil, fmt.Errorf("get image: %s", e.Error())
		}

		bmp, e := gozxing.NewBinaryBitmapFromImage(img)
		if e != nil {
			return nil, fmt.Errorf("create new binary bitmap: %s", e.Error())
		}

		qrCodeRes, _ := qrDecoder.Decode(bmp, nil)

		// qrcode in first page
		if i == 0 && qrCodeRes != nil {
			onePDF.QRCode.ID = qrCodeRes.String()
			qrCodeIDs = append(qrCodeIDs, qrCodeRes.String())
			continue
		}

		if qrCodeRes == nil {
			// regular page
			onePDF.Pages = append(onePDF.Pages, img)
		} else {
			// qrcode exist
			allPDFs = append(allPDFs, onePDF)
			onePDF = PDFExtract{}
			onePDF.QRCode.ID = qrCodeRes.String()
			qrCodeIDs = append(qrCodeIDs, qrCodeRes.String())
		}
	}
	allPDFs = append(allPDFs, onePDF) // add last pdf

	//get QR Codes
	qrCodes := []Qrcodes{}
	if len(qrCodeIDs) > 0 {
		if res := DB.Where("id IN ?", qrCodeIDs).Find(&qrCodes); res.Error != nil {
			return nil, res.Error
		}
	}

	qrCodeM := map[string]Qrcodes{}
	for _, qr := range qrCodes {
		qrCodeM[qr.ID] = qr
	}

	// map each pdf extract to qr code
	allPDFsI, e := gubrak.Map(allPDFs, func(each PDFExtract) PDFExtract {
		each.QRCode = qrCodeM[each.QRCode.ID]
		return each
	})
	if e != nil {
		return nil, e
	}
	if v, ok := allPDFsI.([]PDFExtract); ok {
		allPDFs = v
	}

	return allPDFs, nil
}

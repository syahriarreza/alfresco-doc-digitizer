package main

import (
	"fmt"
	"net/http"

	"github.com/gen2brain/go-fitz"
	"github.com/gin-gonic/gin"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/syahriarreza/alfresco-doc-digitizer/engine"
)

func main() {
	router := gin.Default()

	router.GET("/account", engine.AccountEngine)

	router.GET("/pentol", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	router.POST("/user/:name/*action", func(c *gin.Context) {
		if c.FullPath() == "/user/:name/*action" {
			c.String(http.StatusOK, "masuuukk")
		}
	})

	router.Run(":8080")
}

func readQRCodeInPDF(filePath string) error {
	qrDecoder := qrcode.NewQRCodeReader()
	doc, e := fitz.New(filePath)
	if e != nil {
		panic(e)
	}
	defer doc.Close()

	// Extract pages as images
	for n := 0; n < doc.NumPage(); n++ {
		img, e := doc.Image(n)
		if e != nil {
			return fmt.Errorf("get image: %s", e.Error())
		}

		bmp, e := gozxing.NewBinaryBitmapFromImage(img)
		if e != nil {
			return fmt.Errorf("create new binary bitmap: %s", e.Error())
		}

		res, _ := qrDecoder.Decode(bmp, nil)
		if res == nil {
			fmt.Printf("Page #%03d: QR Code not found\n", n+1)
		} else {
			fmt.Printf("Page #%03d: %s\n", n+1, res.String())
		}
	}

	return nil
}

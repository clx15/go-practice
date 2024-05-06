package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"net/http"
	"os"
	"strconv"
)

func main() {
	r := gin.Default()

	// Define the /generate endpoint
	r.POST("/generate", func(c *gin.Context) {
		// Get parameters from the POST request
		sizeStr := c.PostForm("size")
		content := c.PostForm("content")

		// Convert size from string to integer
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid size parameter"})
			return
		}

		// Generate the QR code image
		qrCodeData, err := qrcode.Encode(content, qrcode.Medium, size)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code"})
			return
		}

		// Save the QR code image to a file
		filename := "qrcode.png"
		file, err := os.Create(filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file"})
			return
		}
		defer file.Close()

		// Write QR code data directly to file as it's already PNG encoded
		if _, err := file.Write(qrCodeData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save QR code image"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "QR code generated and saved successfully", "filename": filename})
	})

	// Serve the HTML page with the QR code
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

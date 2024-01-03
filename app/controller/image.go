package utility

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {

	v1 := r.Group("/api/v1")
	{
		v1.POST("upload", upload)
	}
}

func upload(c *gin.Context) {
	// Retrieve the uploaded file from the form-data
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve uploaded file"})
		return
	}

	// Create the destination file on the server
	savePath := "/images/" + file.Filename
	output, err := os.Create(savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}
	defer output.Close()

	// Save the uploaded file to the destination
	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})
}

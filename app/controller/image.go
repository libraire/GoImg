package utility

import (
	"net/http"
	"os"
	"strconv"

	utils "lensman/app/utils"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {

	v1 := r.Group("/api/v1")
	{
		v1.POST("upload", upload)
		v1.GET("health", healthCheck)
	}
}

func healthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func upload(c *gin.Context) {

	email := c.Request.Context().Value("email").(*string)

	if *email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Retrieve the uploaded file from the form-data
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve uploaded file"})
		return
	}

	position := c.Request.FormValue("position")
	if position == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Position is missing"})
		return
	}

	pos, num_err := strconv.Atoi(position)

	if num_err != nil || (pos > 12 || pos < 0) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Position must between 1 to 12"})
		return
	}

	err = utils.MakeUserImageFolder(*email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	savePath := utils.GenerateImageSourcePath(*email, uint(pos))
	output, err := os.Create(savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer output.Close()

	// Save the uploaded file to the destination
	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})
}

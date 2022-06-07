package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"product-api/dtos"
	"product-api/model"
	"product-api/services"
)

func FileUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		//upload
		formfile, _, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				dtos.MediaDto{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Select a file to upload"},
				})
			return
		}

		uploadUrl, err := services.NewMediaUpload().FileUpload(model.File{File: formfile})
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				dtos.MediaDto{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Error uploading file"},
				})
			return
		}

		c.JSON(
			http.StatusOK,
			dtos.MediaDto{
				StatusCode: http.StatusOK,
				Message:    "success",
				Data:       map[string]interface{}{"data": uploadUrl},
			})
	}
}
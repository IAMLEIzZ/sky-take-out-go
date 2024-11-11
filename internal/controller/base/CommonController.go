package base

import (
	"log"
	"sky-take-out-go/internal/api/response"
	"sky-take-out-go/utils"
	"github.com/gin-gonic/gin"
)

// upload a file
// PATH: admin/common/upload
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("ERROR: " + err.Error())
		response.Response_Error(c)
		return 
	}

	// TOD：Upload file to OSS
	url, err := utils.UploadFileToOss(file)
	
	if err != nil {
		log.Println("ERROR: " + err.Error())
		response.Response_Error(c)
		return 
	}

	response.Response_Success(c, url)	
}
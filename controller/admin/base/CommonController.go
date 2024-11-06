package base

import (
	"log"
	"sky-take-out-go/controller/common"

	"github.com/gin-gonic/gin"
)

// upload a file
// PATH: admin/common/upload
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("ERROR: " + err.Error())
		common.Response_Error(c)
		return
	}

	// TOD：Upload file to OSS
	// (Temp) Upload file to local
	err = c.SaveUploadedFile(file, "./upload/"+file.Filename)
	
	if err != nil {
		log.Println("ERROR: " + err.Error())
		common.Response_Error(c)
		return
	}

	common.Response_Success(c, nil)
}
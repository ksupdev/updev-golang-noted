package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	runningDir, _ := os.Getwd()
	count := 0

	// Set a lower memory limit for multipart from (default is 32 Mib)
	router.MaxMultipartMemory = 8 << 20 //8 Mib

	router.POST("/upload", func(c *gin.Context) {
		// simple file
		file, _ := c.FormFile("file")
		fileExt := filepath.Ext(file.Filename)
		count++

		// c.SaveUploadedFile(file, fmt.Sprintf("%s/upload/%s", runningDir, file.Filename))
		c.SaveUploadedFile(file, fmt.Sprintf("%s/upload/%d%s", runningDir, count, fileExt))
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	router.Run(":85")
}

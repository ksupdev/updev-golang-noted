package upload

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type CsvHdl struct {
	FileDir string
}

func NewHdl(dir string) *CsvHdl {
	return &CsvHdl{FileDir: dir}
}

func (csvHdl *CsvHdl) UploadCsvHandle(c *gin.Context) {
	name := c.PostForm("name")
	fileHeader, err := c.FormFile("file")

	// Handler file size
	if err != nil {
		c.JSON(http.StatusBadRequest, &Resp{IsError: true, ErrorMessage: err.Error()})
		return
	}

	// Validate file type
	if fileHeader.Header.Get("Content-Type") != "text/csv" {
		c.JSON(http.StatusBadRequest, &Resp{IsError: true, ErrorMessage: "this fuction use for upload csv file only"})
		return
	}

	err = saveFile(fileHeader, csvHdl.FileDir)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Resp{IsError: true, ErrorMessage: err.Error()})
		return
	}

	//fmt.Println(fileHeader.Header.Get("Content-Type"))

	/*
		fmt.Println(fileHeader.Header) => map[Content-Disposition:[form-data; name="file"; filename="01.csv"] Content-Type:[text/csv]]
	*/
	extension := filepath.Ext(fileHeader.Filename)
	fmt.Printf("%v ,%v ,%v \n", name, fileHeader.Size, extension)
}

func saveFile(fileHeader *multipart.FileHeader, dir string) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}

	defer file.Close()
	buff := make([]byte, 512)

	_, err = file.Read(buff)
	if err != nil {
		return err
	}

	// Validate Content type of buffer
	fileType := http.DetectContentType(buff)
	// fmt.Printf("----- %v -----", fileType)
	if fileType != "text/plain; charset=utf-8" {
		return fmt.Errorf("The provided file format is not allowed. Please upload CSV file only")
	}

	// ?
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	// create directory for keep file
	err = os.MkdirAll(dir, os.ModePerm)
	// err = os.MkdirAll("upload/file_uploads", os.ModePerm) => create directory in folder upload
	if err != nil {
		return err
	}

	f, err := os.Create(fmt.Sprintf("%v/%d%s", dir, time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
	if err != nil {
		return err
	}

	defer f.Close()

	pr := &Progress{
		TotalSize: fileHeader.Size,
	}

	_, err = io.Copy(f, io.TeeReader(file, pr))
	if err != nil {
		return err
	}

	// pr.Print()

	return nil
}

/*
func (uhdl *UploadHdl) UploadMultipleHandler(c *gin.Context) {
	name := c.PostForm("name")
	// files, err := c.FormFile("file")

	files := c.Request.MultipartForm.File["file"]
	for _, file := range files {

	}

	// Handler file size
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"is_error": true,
	// 		"message":  err.Error(),
	// 	})
	// 	return
	// }

	// if file.Size > (1024 * 1024) {

	// }

	extension := filepath.Ext(file.Filename)
	fmt.Printf("%v ,%v ,%v \n", name, file.Size, extension)


	c.JSON(http.StatusOK, map[string]interface{}{
		"data":     fmt.Sprintf("%v ,%v ,%v", name, file.Size, extension),
		"is_error": false,
		"message":  "",
	})
}

*/

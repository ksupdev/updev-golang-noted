package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"updev.com/go-upload-csv/manual_synunit"
	"updev.com/go-upload-csv/upload"
)

func main() {

	const UPLOAD_DIRECTORY = "./datas/file_uploads"
	const ACHIEVE_DIRECTORY = "./datas/achieve"

	port_str := fmt.Sprintf(":%v", 8081)
	router := initWebServer()

	// router.LoadHTMLGlob("templates/**/*.html")
	router.LoadHTMLGlob("*.html")
	upHdl := upload.NewHdl(UPLOAD_DIRECTORY)
	router.POST("api/upload", upHdl.UploadCsvHandle)
	router.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	synUnit := manual_synunit.NewExeHdl(UPLOAD_DIRECTORY, ACHIEVE_DIRECTORY)

	go func() {
		for {
			err := synUnit.Execute(action)
			if err != nil {
				fmt.Println(err)
			}

			// file, err := os.Open("data/import-template.csv")
			// reader := csv.NewReader(file)
			// records, _ := reader.ReadAll()

			// fmt.Println(records)

			time.Sleep(2 * time.Second)
		}
	}()

	/*
		router.POST("api/old/upload", func(c *gin.Context) {

			name := c.PostForm("name")
			file, _ := c.FormFile("file")
			extension := filepath.Ext(file.Filename)

			c.JSON(http.StatusOK, map[string]interface{}{
				"result":        fmt.Sprintf("%v ,%v ,%v", name, file, extension),
				"response_from": "mock-serv-01",
			})

		})

		router.GET("api/data", func(c *gin.Context) {

			c.JSON(http.StatusOK, map[string]interface{}{
				"result":        "result",
				"response_from": "mock-serv-01",
			})

		})

	*/

	// router.Static("/page", "index.html")

	err := router.Run(port_str)
	if err != nil {
		panic(err.Error())
	}

}

func action(data string) {
	log.Printf("v : %v \n", data)
	row := strings.Split(data, ",")
	log.Printf("v : %v = %v \n", row[0], row[1])
	// return nil
}

func initWebServer() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))
	router.Use(bodySizeMiddleware)
	log.Printf("[Startup] Webserver start success ")
	return router
}
func bodySizeMiddleware(c *gin.Context) {
	//var maxBytes int64 = 1024 * 1024 * 5 // 5MB
	const MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB
	var w http.ResponseWriter = c.Writer
	c.Request.Body = http.MaxBytesReader(w, c.Request.Body, MAX_UPLOAD_SIZE)
	c.Next()
}

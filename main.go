package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func upload(c *gin.Context) {
	// Lấy tệp từ form với khóa "file"
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Error getting uploaded file:", err)
		c.String(http.StatusBadRequest, "Error getting uploaded file")
		return
	}

	log.Println("Uploaded file:", file.Filename)

	// Đường dẫn để lưu tệp
	uploadDir := "./upload"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		log.Println("Error creating upload directory:", err)
		c.String(http.StatusInternalServerError, "Error creating upload directory")
		return
	}

	dst := filepath.Join(uploadDir, file.Filename)

	// Lưu tệp đến đường dẫn cụ thể
	if err := c.SaveUploadedFile(file, dst); err != nil {
		log.Println("Error saving uploaded file:", err)
		c.String(http.StatusInternalServerError, "Error saving uploaded file")
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	v2 := r.Group("/v1")
	{
		v2.POST("/upload", upload)

	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)

	})

	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	r.POST("/user/:name/*action", func(c *gin.Context) {

		b := c.FullPath() == "/user/:name/*action" // true
		c.String(http.StatusOK, "%t", b)
	})

	r.GET("/user/groups", func(c *gin.Context) {
		c.String(http.StatusOK, "The available groups are [...]")
	})

	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"page":    page,
			"name":    name,
			"message": message,
		})
	})
	r.Run() // listen and serve on
}

// https://jonathanmh.com/creating-simple-markdown-blog-go-gin/

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	//    "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

type Post struct {
	Title   string
	Content template.HTML
}

func main() {

	r := gin.Default()
	r.Use(gin.Logger())
	r.Delims("{{", "}}")

	r.Static("/assets", "./assets")
	// r.Use(static.Serve("/assets", static.LocalFile("/assets", false)))
	r.LoadHTMLGlob("./templates/*.tmpl.html")

	r.GET("/", func(c *gin.Context) {
		var posts []string

		files, err := ioutil.ReadDir("./markdown/")
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			// fmt.Println(file.Name())
			posts = append(posts, file.Name())
		}

		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"posts": posts,
		})
	})

	r.GET("/:postName", func(c *gin.Context) {
		postName := c.Param("postName")

		// deprecated
		// mdfile, err := ioutil.ReadFile("./markdown/" + postName)
		mdfile, err := os.ReadFile("./markdown/" + postName)

		// if the file can not be found
		if err != nil {
			fmt.Println(err)
			c.HTML(http.StatusNotFound, "error.tmpl.html", nil)
			return
		}

		postHTML := template.HTML(blackfriday.MarkdownCommon([]byte(mdfile)))

		post := Post{Title: postName, Content: postHTML}

		c.HTML(http.StatusOK, "post.tmpl.html", gin.H{
			"Title":   post.Title,
			"Content": post.Content,
		})
	})

	r.Run()
}

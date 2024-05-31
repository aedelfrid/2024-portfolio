package main

import (

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"

	views "aedelfrid/portfolio-server/views"
)

func render (c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func main() {
	r := gin.Default()

	index := views.Index()

	r.GET("/", func (c *gin.Context)  {
		render(c, 200, index)
	})

	r.Run("0.0.0.0:8000") // listen and serve on 0.0.0.0:8000 (for windows "localhost:8080")
}

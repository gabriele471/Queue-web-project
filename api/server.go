package api

import (
	"net/http"
	"structure/data"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{ //just put the page

	})
}
func All(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Queue)
}
func Insert(c *gin.Context) {
	var newJob data.Job

	if err := c.Bind(&newJob); err != nil { //form not returned in json
		return //c.AbortWithError(400, err) need an interface for the err //need to handle the error and redirect to some page
	}

	data.Enqueue(newJob)
	//c.IndentedJSON(http.StatusCreated, newJob) redirect to insert with json on screen
	c.Redirect(http.StatusFound, "/")

}

func Delete(c *gin.Context) {

	data.PopQueue()
	c.Redirect(http.StatusFound, "/")
}

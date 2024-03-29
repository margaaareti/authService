package mainhttp

import (
	"Test_derictory/auth"
	"Test_derictory/mainpage"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndPoints(c *gin.RouterGroup, crd mainpage.HomePage, auth auth.UseCase) {
	cr := NewHomeHandler(crd, auth)

	c.GET("/home", cr.ShowPage)
	c.POST("/log-out", cr.LogOut)
	c.POST("/home/add", cr.NewStudent)
	c.GET("/home/getAll", cr.GetAllNotes)
	c.POST("/home/delete/:id", cr.DeleteNoteById)
	c.GET("/home/getEntry/:id", cr.GetById)
	c.PUT("/home/setEntry/:id", cr.UpdateItem)
	c.POST("/home/getEntry/:id/addImg", cr.UploadFile)

}

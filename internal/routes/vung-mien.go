package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DocVungMienJson(c *gin.Context) {
	dsVungMien, _ := models.DocVungMienCSDL()
	c.IndentedJSON(http.StatusOK, dsVungMien)
}

func ThemVungMienJson(c *gin.Context) {
	var vungMien models.VungMien

	if err := c.BindJSON(&vungMien); err != nil {
		fmt.Print(err.Error())
		return
	}

	nguyenLieu, err := models.ThemVungMienCSDL(vungMien)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, nguyenLieu)
	} else {
		c.IndentedJSON(http.StatusCreated, nguyenLieu)
	}
}

func CapNhatVungMienJson(c *gin.Context) {
	var vungMien models.VungMien

	if err := c.BindJSON(&vungMien); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.CapNhatVungMienCSDL(vungMien)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func XoaVungMienJson(c *gin.Context) {
	var Doc map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.XoaVungMienCSDL(Doc["id"])
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

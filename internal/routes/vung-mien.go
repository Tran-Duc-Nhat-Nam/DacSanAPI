package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DocVungMienAPI(c *gin.Context) {
	dsVungMien, _ := models.DocVungMien()
	c.IndentedJSON(http.StatusOK, dsVungMien)
}

func ThemVungMienAPI(c *gin.Context) {
	var vungMien models.VungMien

	if err := c.BindJSON(&vungMien); err != nil {
		fmt.Print(err.Error())
		return
	}

	nguyenLieu, err := models.ThemVungMien(vungMien)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, nguyenLieu)
	} else {
		c.IndentedJSON(http.StatusCreated, nguyenLieu)
	}
}

func CapNhatVungMienAPI(c *gin.Context) {
	var vungMien models.VungMien

	if err := c.BindJSON(&vungMien); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.CapNhatVungMien(vungMien)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func XoaVungMienAPI(c *gin.Context) {
	var Doc map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.XoaVungMien(Doc["id"])
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

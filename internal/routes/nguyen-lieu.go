package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DocNguyenLieuJson(c *gin.Context) {
	dsNguyenLieu, _ := models.DocNguyenLieuCSDL()
	c.IndentedJSON(http.StatusOK, dsNguyenLieu)
}

func ThemNguyenLieuJson(c *gin.Context) {
	var nguyenLieu models.NguyenLieu

	if err := c.BindJSON(&nguyenLieu); err != nil {
		fmt.Print(err.Error())
		return
	}

	nguyenLieu, err := models.ThemNguyenLieuCSDL(nguyenLieu)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, nguyenLieu)
	} else {
		c.IndentedJSON(http.StatusCreated, nguyenLieu)
	}
}

func CapNhatNguyenLieuJson(c *gin.Context) {
	var nguyenLieu models.NguyenLieu

	if err := c.BindJSON(&nguyenLieu); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.CapNhatNguyenLieuCSDL(nguyenLieu)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func XoaNguyenLieuJson(c *gin.Context) {
	var Doc map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.XoaNguyenLieuCSDL(Doc["id"])
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

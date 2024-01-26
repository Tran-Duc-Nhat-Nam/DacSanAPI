package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DocTinhThanhJson(c *gin.Context) {
	tinhThanh, err := models.DocTinhThanhCSDL()
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, tinhThanh)
}

func ThemTinhThanhJson(c *gin.Context) {
	var tinhThanh models.TinhThanh

	if err := c.BindJSON(&tinhThanh); err != nil {
		fmt.Print(err.Error())
		return
	}

	tinhThanh, err := models.ThemTinhThanhCSDL(tinhThanh)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, tinhThanh)
	} else {
		c.IndentedJSON(http.StatusCreated, tinhThanh)
	}
}

func CapNhatTinhThanhJson(c *gin.Context) {
	var tinhThanh models.TinhThanh

	if err := c.BindJSON(&tinhThanh); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.CapNhatTinhThanhCSDL(tinhThanh)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func XoaTinhThanhJson(c *gin.Context) {
	var Doc map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.XoaTinhThanhCSDL(Doc["id"])
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

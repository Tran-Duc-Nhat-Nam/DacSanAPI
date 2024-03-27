package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DocTinhThanhAPI(c *gin.Context) {
	tinhThanh, err := models.DocTinhThanh()
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, tinhThanh)
}

func ThemTinhThanhAPI(c *gin.Context) {
	var tinhThanh models.TinhThanh

	if err := c.BindJSON(&tinhThanh); err != nil {
		fmt.Print(err.Error())
		return
	}

	tinhThanh, err := models.ThemTinhThanh(tinhThanh)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, tinhThanh)
	} else {
		c.IndentedJSON(http.StatusCreated, tinhThanh)
	}
}

func CapNhatTinhThanhAPI(c *gin.Context) {
	var tinhThanh models.TinhThanh

	if err := c.BindJSON(&tinhThanh); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.CapNhatTinhThanh(tinhThanh)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func XoaTinhThanhAPI(c *gin.Context) {
	var Doc map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.XoaTinhThanh(Doc["id"])
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

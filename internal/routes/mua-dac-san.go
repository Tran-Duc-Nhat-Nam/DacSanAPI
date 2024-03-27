package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DocMuaDacSanAPI(c *gin.Context) {
	dsMuaDacSan, _ := models.DocMua()
	c.IndentedJSON(http.StatusOK, dsMuaDacSan)
}

func ThemMuaDacSanAPI(c *gin.Context) {
	var muaDacSan models.MuaDacSan

	if err := c.BindJSON(&muaDacSan); err != nil {
		fmt.Print(err.Error())
		return
	}

	muaDacSan, err := models.ThemMua(muaDacSan)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, muaDacSan)
	} else {
		c.IndentedJSON(http.StatusCreated, muaDacSan)
	}
}

func CapNhatMuaDacSanAPI(c *gin.Context) {
	var muaDacSan models.MuaDacSan

	if err := c.BindJSON(&muaDacSan); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.CapNhatMua(muaDacSan)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func XoaMuaDacSanAPI(c *gin.Context) {
	var Doc map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.XoaMua(Doc["id"])
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

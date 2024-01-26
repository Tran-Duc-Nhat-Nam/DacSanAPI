package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DocMuaDacSanJson(c *gin.Context) {
	dsMuaDacSan, _ := models.DocMuaCSDL()
	c.IndentedJSON(http.StatusOK, dsMuaDacSan)
}

func ThemMuaDacSanJson(c *gin.Context) {
	var muaDacSan models.MuaDacSan

	if err := c.BindJSON(&muaDacSan); err != nil {
		fmt.Print(err.Error())
		return
	}

	muaDacSan, err := models.ThemMuaCSDL(muaDacSan)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, muaDacSan)
	} else {
		c.IndentedJSON(http.StatusCreated, muaDacSan)
	}
}

func CapNhatMuaDacSanJson(c *gin.Context) {
	var muaDacSan models.MuaDacSan

	if err := c.BindJSON(&muaDacSan); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.CapNhatMuaCSDL(muaDacSan)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func XoaMuaDacSanJson(c *gin.Context) {
	var Doc map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.XoaMuaCSDL(Doc["id"])
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

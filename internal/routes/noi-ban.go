package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DocNoiBanJson(c *gin.Context) {
	dsNoiBan, _ := models.DocNoiBanCSDL()
	c.IndentedJSON(http.StatusOK, dsNoiBan)
}

func DocNoiBanTheoIdJson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	dacSan, err := models.DocNoiBanTheoIdCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func ThemNoiBanJson(c *gin.Context) {
	var noiBan models.NoiBan

	if err := c.BindJSON(&noiBan); err != nil {
		fmt.Print(err.Error())
		return
	}

	noiBan, err := models.ThemNoiBanCSDL(noiBan)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, noiBan)
	} else {
		c.IndentedJSON(http.StatusCreated, noiBan)
	}
}

func CapNhatNoiBanJson(c *gin.Context) {
	var noiBan models.NoiBan

	if err := c.BindJSON(&noiBan); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.CapNhatNoiBanCSDL(noiBan)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func XoaNoiBanJson(c *gin.Context) {
	var Doc map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.XoaNoiBanCSDL(Doc["id"])
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

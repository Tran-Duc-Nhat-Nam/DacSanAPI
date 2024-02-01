package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DocNguoiDungJson(c *gin.Context) {
	nguoiDung, err := models.DocNguoiDungCSDL()
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, nguoiDung)
}

func DocNguoiDungTheoIdJson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	nguoiDung, err := models.DocNguoiDungTheoIdCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, nguoiDung)
}

func ThemNguoiDungJson(c *gin.Context) {
	var nguoiDung models.NguoiDung

	if err := c.BindJSON(&nguoiDung); err != nil {
		fmt.Print(err.Error())
		return
	}

	nguoiDung, err := models.ThemNguoiDungCSDL(nguoiDung)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, nguoiDung)
	} else {
		c.IndentedJSON(http.StatusCreated, nguoiDung)
	}
}

func CapNhatNguoiDungJson(c *gin.Context) {
	var nguoiDung models.NguoiDung

	if err := c.BindJSON(&nguoiDung); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.CapNhatNguoiDungCSDL(nguoiDung)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func XoaNguoiDungJson(c *gin.Context) {
	var Doc map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.XoaNguoiDungCSDL(Doc["id"])
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}
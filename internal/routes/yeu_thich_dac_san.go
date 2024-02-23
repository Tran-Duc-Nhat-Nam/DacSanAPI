package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func YeuThichDacSan(c *gin.Context) {
	var yeuThich models.YeuThichDacSan

	if err := c.BindJSON(&yeuThich); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}

	if err := models.ThemYeuThichDacSanCSDL(yeuThich); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}

	c.IndentedJSON(http.StatusOK, true)
}

func BoYeuThichDacSan(c *gin.Context) {
	var yeuThich models.YeuThichDacSan

	if err := c.BindJSON(&yeuThich); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}

	if err := models.XoaYeuThichDacSanCSDL(yeuThich); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}

	c.IndentedJSON(http.StatusOK, true)
}

func DocYeuThichDacSan(c *gin.Context) {
	var yeuThich models.YeuThichDacSan

	if err := c.BindJSON(&yeuThich); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}

	if err := models.DocYeuThichDacSanCSDL(yeuThich); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}

	c.IndentedJSON(http.StatusOK, true)
}

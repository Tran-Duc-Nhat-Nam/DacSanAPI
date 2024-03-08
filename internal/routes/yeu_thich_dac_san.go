package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func YeuThichDacSan(c *gin.Context) {
	idDacSan, err := strconv.Atoi(c.Param("idDacSan"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	idNguoiDung := c.Param("idNguoiDung")

	if err := models.ThemYeuThichDacSanCSDL(models.YeuThichDacSan{IdNguoiDung: idNguoiDung, IdDacSan: idDacSan}); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

func BoYeuThichDacSan(c *gin.Context) {
	idDacSan, err := strconv.Atoi(c.Param("idDacSan"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	idNguoiDung := c.Param("idNguoiDung")

	if err := models.XoaYeuThichDacSanCSDL(models.YeuThichDacSan{IdNguoiDung: idNguoiDung, IdDacSan: idDacSan}); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

func DocYeuThichDacSan(c *gin.Context) {
	idDacSan, err := strconv.Atoi(c.Param("idDacSan"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	idNguoiDung := c.Param("idNguoiDung")

	if err := models.DocYeuThichDacSanCSDL(models.YeuThichDacSan{IdNguoiDung: idNguoiDung, IdDacSan: idDacSan}); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func YeuThichDacSanAPI(c *gin.Context) {
	idDacSan, err := strconv.Atoi(c.Param("idDacSan"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	idNguoiDung := c.Param("idNguoiDung")

	if err := models.ThemYeuThichDacSan(models.YeuThichDacSan{IdNguoiDung: idNguoiDung, IdDacSan: idDacSan}); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

func BoYeuThichDacSanAPI(c *gin.Context) {
	idDacSan, err := strconv.Atoi(c.Param("idDacSan"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	idNguoiDung := c.Param("idNguoiDung")

	if err := models.XoaYeuThichDacSan(models.YeuThichDacSan{IdNguoiDung: idNguoiDung, IdDacSan: idDacSan}); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

func DocYeuThichDacSanAPI(c *gin.Context) {
	idDacSan, err := strconv.Atoi(c.Param("idDacSan"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	idNguoiDung := c.Param("idNguoiDung")

	if err := models.DocYeuThichDacSan(models.YeuThichDacSan{IdNguoiDung: idNguoiDung, IdDacSan: idDacSan}); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

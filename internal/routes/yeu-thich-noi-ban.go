package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func YeuThichNoiBanAPI(c *gin.Context) {
	idNoiBan, err := strconv.Atoi(c.Param("idNoiBan"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	idNguoiDung := c.Param("idNguoiDung")

	if err := models.ThemYeuThichNoiBan(models.YeuThichNoiBan{IdNguoiDung: idNguoiDung, IdNoiBan: idNoiBan}); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

func BoYeuThichNoiBanAPI(c *gin.Context) {
	idNoiBan, err := strconv.Atoi(c.Param("idNoiBan"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	idNguoiDung := c.Param("idNguoiDung")

	if err := models.XoaYeuThichNoiBan(models.YeuThichNoiBan{IdNguoiDung: idNguoiDung, IdNoiBan: idNoiBan}); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

func DocYeuThichNoiBanAPI(c *gin.Context) {
	idNoiBan, err := strconv.Atoi(c.Param("idNoiBan"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}
	idNguoiDung := c.Param("idNguoiDung")

	if err := models.DocYeuThichNoiBan(models.YeuThichNoiBan{IdNguoiDung: idNguoiDung, IdNoiBan: idNoiBan}); err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

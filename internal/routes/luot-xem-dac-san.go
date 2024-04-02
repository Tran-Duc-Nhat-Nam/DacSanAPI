package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nam/dac_san_api/internal/models"
	"net/http"
)

func DocDanhSachDacSanDaXemAPI(c *gin.Context) {
	idNguoiDung := c.Param("idNguoiDung")

	ds, err := models.DocDanhSachDacSanDaXem(idNguoiDung)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, ds)
}

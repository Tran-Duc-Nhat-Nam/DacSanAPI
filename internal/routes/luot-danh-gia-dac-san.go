package routes

import (
	"nam/dac_san_api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TinhDiemDanhGiaDacSanTheoIdJson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusConflict, 0)
	}
	c.IndentedJSON(http.StatusOK, models.TinhDiemDanhGiaDacSanCSDL(id))
}

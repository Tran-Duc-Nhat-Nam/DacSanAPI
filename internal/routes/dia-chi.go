package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DocDiaChiTheoIdAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	diaChi, err := models.DocDiaChiTheoId(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, diaChi)
}

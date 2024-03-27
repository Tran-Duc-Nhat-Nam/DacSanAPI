package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DocHinhAnhTheoIdAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	hinhAnh, err := models.DocHinhAnhTheoId(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, hinhAnh)
}

func DocHinhAnhDacSanAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	hinhAnh, err := models.DocHinhAnhDacSan(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, hinhAnh)
}

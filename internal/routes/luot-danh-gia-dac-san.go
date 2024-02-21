package routes

import (
	"fmt"
	"nam/dac_san_api/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func TinhDiemDanhGiaDacSanTheoIdJson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusConflict, 0)
	}
	c.IndentedJSON(http.StatusOK, models.TinhDiemDanhGiaDacSanCSDL(id))
}

func DanhGiaDacSan(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	idNguoiDung := c.Param("idnguoidung")
	dacSan, err := models.DocDacSanTheoIdCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	err = models.ThemDanhGiaDacSanCSDL(models.LuotDanhGiaDacSan{IdNguoiDung: idNguoiDung, IdDacSan: dacSan.ID, ThoiGianDanhGia: time.Now(), DiemDanhGia: 5})
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, err.Error())
	}
	c.IndentedJSON(http.StatusOK, true)
}

package main

import (
	"nam/dac_san_api/internal/models"

	// "encoding/json"
	"strconv"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DocTinhThanhJson(c *gin.Context) {
	tinhThanh, err := models.DocTinhThanhCSDL()
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, tinhThanh)
}

func ThemTinhThanhJson(c *gin.Context) {
	var tinhThanh models.TinhThanh

	if err := c.BindJSON(&tinhThanh); err != nil {
		fmt.Print(err.Error())
		return
	}

	tinhThanh, err := models.ThemTinhThanhCSDL(tinhThanh)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, tinhThanh)
	} else {
		c.IndentedJSON(http.StatusCreated, tinhThanh)
	}
}

func CapNhatTinhThanhJson(c *gin.Context) {
	var tinhThanh models.TinhThanh

	if err := c.BindJSON(&tinhThanh); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.CapNhatTinhThanhCSDL(tinhThanh)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func XoaTinhThanhJson(c *gin.Context) {
	var Doc map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.XoaTinhThanhCSDL(Doc["id"])
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func DocQuanHuyenTheoTinhThanhJson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	quanHuyen, err := models.DocQuanHuyenTheoTinhThanhCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, quanHuyen)
}

func DocPhuongXaTheoQuanHuyenJson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	phuongXa, err := models.DocPhuongXaTheoQuanHuyenCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, phuongXa)
}

func DocDiaChiTheoIdJson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	diaChi, err := models.DocDiaChiTheoIdCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, diaChi)
}

func DocVungMienJson(c *gin.Context) {
	dsVungMien, _ := models.DocVungMienCSDL()
	c.IndentedJSON(http.StatusOK, dsVungMien)
}

func ThemVungMienJson(c *gin.Context) {
	var vungMien models.VungMien

	if err := c.BindJSON(&vungMien); err != nil {
		fmt.Print(err.Error())
		return
	}

	nguyenLieu, err := models.ThemVungMienCSDL(vungMien)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, nguyenLieu)
	} else {
		c.IndentedJSON(http.StatusCreated, nguyenLieu)
	}
}

func CapNhatVungMienJson(c *gin.Context) {
	var vungMien models.VungMien

	if err := c.BindJSON(&vungMien); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.CapNhatVungMienCSDL(vungMien)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func XoaVungMienJson(c *gin.Context) {
	var Doc map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.XoaVungMienCSDL(Doc["id"])
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func DocMuaDacSanJson(c *gin.Context) {
	dsMuaDacSan, _ := models.DocMuaCSDL()
	c.IndentedJSON(http.StatusOK, dsMuaDacSan)
}

func ThemMuaDacSanJson(c *gin.Context) {
	var muaDacSan models.MuaDacSan

	if err := c.BindJSON(&muaDacSan); err != nil {
		fmt.Print(err.Error())
		return
	}

	muaDacSan, err := models.ThemMuaCSDL(muaDacSan)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, muaDacSan)
	} else {
		c.IndentedJSON(http.StatusCreated, muaDacSan)
	}
}

func CapNhatMuaDacSanJson(c *gin.Context) {
	var muaDacSan models.MuaDacSan

	if err := c.BindJSON(&muaDacSan); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.CapNhatMuaCSDL(muaDacSan)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func XoaMuaDacSanJson(c *gin.Context) {
	var Doc map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.XoaMuaCSDL(Doc["id"])
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func DocNguyenLieuJson(c *gin.Context) {
	dsNguyenLieu, _ := models.DocNguyenLieuCSDL()
	c.IndentedJSON(http.StatusOK, dsNguyenLieu)
}

func ThemNguyenLieuJson(c *gin.Context) {
	var nguyenLieu models.NguyenLieu

	if err := c.BindJSON(&nguyenLieu); err != nil {
		fmt.Print(err.Error())
		return
	}

	nguyenLieu, err := models.ThemNguyenLieuCSDL(nguyenLieu)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, nguyenLieu)
	} else {
		c.IndentedJSON(http.StatusCreated, nguyenLieu)
	}
}

func CapNhatNguyenLieuJson(c *gin.Context) {
	var nguyenLieu models.NguyenLieu

	if err := c.BindJSON(&nguyenLieu); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.CapNhatNguyenLieuCSDL(nguyenLieu)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func XoaNguyenLieuJson(c *gin.Context) {
	var Doc map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.XoaNguyenLieuCSDL(Doc["id"])
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func DocHinhAnhTheoIdJson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	hinhAnh, err := models.DocHinhAnhTheoIdCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, hinhAnh)
}

func DocDacSanJson(c *gin.Context) {
	dsDacSan, _ := models.DocDacSanCSDL()
	c.IndentedJSON(http.StatusOK, dsDacSan)
}

func DocDacSanTheoIdJson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	dacSan, err := models.DocDacSanTheoIdCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func ThemDacSanJson(c *gin.Context) {
	var dacSan models.DacSan

	if err := c.BindJSON(&dacSan); err != nil {
		fmt.Print(err.Error())
		return
	}

	dacSan, err := models.ThemDacSanCSDL(dacSan)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, dacSan)
	} else {
		c.IndentedJSON(http.StatusCreated, dacSan)
	}
}

func CapNhatDacSanJson(c *gin.Context) {
	var dacSan models.DacSan

	if err := c.BindJSON(&dacSan); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.CapNhatDacSanCSDL(dacSan)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func XoaDacSanJson(c *gin.Context) {
	var Doc map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.XoaDacSanCSDL(Doc["id"])
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func DocNoiBanJson(c *gin.Context) {
	dsNoiBan, _ := models.DocNoiBanCSDL()
	c.IndentedJSON(http.StatusOK, dsNoiBan)
}

func DocNoiBanTheoIdJson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	dacSan, err := models.DocNoiBanTheoIdCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, dacSan)
}

func ThemNoiBanJson(c *gin.Context) {
	var noiBan models.NoiBan

	if err := c.BindJSON(&noiBan); err != nil {
		fmt.Print(err.Error())
		return
	}

	noiBan, err := models.ThemNoiBanCSDL(noiBan)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, noiBan)
	} else {
		c.IndentedJSON(http.StatusCreated, noiBan)
	}
}

func CapNhatNoiBanJson(c *gin.Context) {
	var noiBan models.NoiBan

	if err := c.BindJSON(&noiBan); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.CapNhatNoiBanCSDL(noiBan)
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

func XoaNoiBanJson(c *gin.Context) {
	var Doc map[string]int

	if err := c.BindJSON(&Doc); err != nil {
		fmt.Print(err.Error())
		return
	}

	err := models.XoaNoiBanCSDL(Doc["id"])
	if err != nil {
		fmt.Print(err.Error())
		c.IndentedJSON(http.StatusConflict, false)
	} else {
		c.IndentedJSON(http.StatusOK, true)
	}
}

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

// func ThemDiaChiWebJson() {
// 	file, err := os.ReadFile("provinces.open-api.vn.json")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	var items []map[string]interface{}
// 	err = json.Unmarshal(file, &items)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	for _, province := range items {
// 		districts := province["districts"].([]interface{})
// 		for _, district := range districts {
// 			districtInfo := district.(map[string]interface{})
// 			var quanHuyen models.QuanHuyen
// 			quanHuyen.ID = int(districtInfo["code"].(float64))
// 			quanHuyen.Ten = districtInfo["name"].(string)
// 			quanHuyen.TinhThanh.ID = int(province["code"].(float64))
// 			quanHuyen, err := models.ThemQuanHuyenCSDL(quanHuyen)
// 			if err != nil {
// 				fmt.Println("err")
// 			} else {
// 				fmt.Println(quanHuyen.Ten)
// 			}
// 			wards := districtInfo["wards"].([]interface{})
// 			for _, ward := range wards {
// 				wardInfo := ward.(map[string]interface{})
// 				var phuongXa models.PhuongXa
// 				phuongXa.ID = int(wardInfo["code"].(float64))
// 				phuongXa.Ten = wardInfo["name"].(string)
// 				phuongXa.QuanHuyen.ID = int(districtInfo["code"].(float64))
// 				phuongXa, err := models.ThemPhuongXaCSDL(phuongXa)
// 				if err != nil {
// 					fmt.Println("err")
// 				} else {
// 					fmt.Println(phuongXa.Ten)
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println("Done")
// }

func main() {
	models.ConnectDatabase()

	// ThemDiaChiWebJson()

	router := gin.Default()
	router.GET("/dacsan", DocDacSanJson)
	router.GET("/dacsan/:id", DocDacSanTheoIdJson)
	router.POST("/dacsan", ThemDacSanJson)
	router.PUT("/dacsan", CapNhatDacSanJson)
	router.DELETE("/dacsan", XoaDacSanJson)
	router.GET("/noiban", DocNoiBanJson)
	router.GET("/noiban/:id", DocNoiBanTheoIdJson)
	router.POST("/noiban", ThemNoiBanJson)
	router.PUT("/noiban", CapNhatNoiBanJson)
	router.DELETE("/noiban", XoaNoiBanJson)
	router.GET("/tinhthanh", DocTinhThanhJson)
	router.GET("/tinhthanh/:id/quanhuyen", DocQuanHuyenTheoTinhThanhJson)
	router.GET("/quanhuyen/:id/phuongxa", DocPhuongXaTheoQuanHuyenJson)
	router.POST("/tinhthanh", ThemTinhThanhJson)
	router.PUT("/tinhthanh", CapNhatTinhThanhJson)
	router.DELETE("/tinhthanh", XoaTinhThanhJson)
	router.GET("/vungmien", DocVungMienJson)
	router.POST("/vungmien", ThemVungMienJson)
	router.PUT("/vungmien", CapNhatVungMienJson)
	router.DELETE("/vungmien", XoaVungMienJson)
	router.GET("/muadacsan", DocMuaDacSanJson)
	router.POST("/muadacsan", ThemMuaDacSanJson)
	router.PUT("/muadacsan", CapNhatMuaDacSanJson)
	router.DELETE("/muadacsan", XoaMuaDacSanJson)
	router.GET("/nguyenlieu", DocNguyenLieuJson)
	router.POST("/nguyenlieu", ThemNguyenLieuJson)
	router.PUT("/nguyenlieu", CapNhatNguyenLieuJson)
	router.DELETE("/nguyenlieu", XoaNguyenLieuJson)
	router.GET("/diachi/:id", DocDiaChiTheoIdJson)
	router.GET("/hinhanh/:id", DocHinhAnhTheoIdJson)
	router.GET("/nguoidung", DocNguoiDungJson)
	router.GET("/nguoidung/:id", DocNguoiDungTheoIdJson)
	router.POST("/nguoidung", ThemNguoiDungJson)
	router.PUT("/nguoidung", CapNhatNguoiDungJson)
	router.DELETE("/nguoidung", XoaNguoiDungJson)
	router.Run("localhost:8080")
}

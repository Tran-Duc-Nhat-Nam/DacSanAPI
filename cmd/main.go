package main

import (
	"nam/dac_san_api/internal/models"
	"nam/dac_san_api/internal/routes"

	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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

	gin.SetMode(gin.ReleaseMode)
	routeConfig := cors.DefaultConfig()
	routeConfig.AllowAllOrigins = true
	routeCors := cors.New(routeConfig)
	fmt.Println(routeConfig.AllowAllOrigins)
	router := gin.Default()
	router.Use(routeCors)
	router.GET("/dacsan", routes.DocDacSanJson)
	router.GET("/dacsan/ten=:ten", routes.DocDacSanTheoTenJson)
	router.GET("/dacsan/mota=:mota", routes.DocDacSanTheoMoTaJson)
	router.GET("/dacsan/cachchebien=:cachchebien", routes.DocDacSanTheoCachCheBienJson)
	router.GET("/dacsan/:id", routes.DocDacSanTheoIdJson)
	router.GET("/dacsan/:id/nguoidung=:idnguoidung", routes.XemDacSan)
	router.POST("/dacsan", routes.ThemDacSanJson)
	router.PUT("/dacsan", routes.CapNhatDacSanJson)
	router.DELETE("/dacsan", routes.XoaDacSanJson)
	router.GET("/noiban", routes.DocNoiBanJson)
	router.GET("/noiban/:id", routes.DocNoiBanTheoIdJson)
	router.GET("/noiban/:id/dacsan", routes.DocDacSanTheoNoiBanJson)
	router.POST("/noiban", routes.ThemNoiBanJson)
	router.PUT("/noiban", routes.CapNhatNoiBanJson)
	router.DELETE("/noiban", routes.XoaNoiBanJson)
	router.GET("/tinhthanh", routes.DocTinhThanhJson)
	router.GET("/tinhthanh/:id/quanhuyen", routes.DocQuanHuyenTheoTinhThanhJson)
	router.GET("/quanhuyen/:id/phuongxa", routes.DocPhuongXaTheoQuanHuyenJson)
	router.POST("/tinhthanh", routes.ThemTinhThanhJson)
	router.PUT("/tinhthanh", routes.CapNhatTinhThanhJson)
	router.DELETE("/tinhthanh", routes.XoaTinhThanhJson)
	router.GET("/vungmien", routes.DocVungMienJson)
	router.POST("/vungmien", routes.ThemVungMienJson)
	router.PUT("/vungmien", routes.CapNhatVungMienJson)
	router.DELETE("/vungmien", routes.XoaVungMienJson)
	router.GET("/muadacsan", routes.DocMuaDacSanJson)
	router.POST("/muadacsan", routes.ThemMuaDacSanJson)
	router.PUT("/muadacsan", routes.CapNhatMuaDacSanJson)
	router.DELETE("/muadacsan", routes.XoaMuaDacSanJson)
	router.GET("/nguyenlieu", routes.DocNguyenLieuJson)
	router.POST("/nguyenlieu", routes.ThemNguyenLieuJson)
	router.PUT("/nguyenlieu", routes.CapNhatNguyenLieuJson)
	router.DELETE("/nguyenlieu", routes.XoaNguyenLieuJson)
	router.GET("/diachi/:id", routes.DocDiaChiTheoIdJson)
	router.GET("/hinhanh/:id", routes.DocHinhAnhTheoIdJson)
	router.GET("/nguoidung", routes.DocNguoiDungJson)
	router.GET("/nguoidung/:id", routes.DocNguoiDungTheoIdJson)
	router.POST("/nguoidung", routes.ThemNguoiDungJson)
	router.PUT("/nguoidung", routes.CapNhatNguoiDungJson)
	router.DELETE("/nguoidung", routes.XoaNguoiDungJson)
	router.Run()
}

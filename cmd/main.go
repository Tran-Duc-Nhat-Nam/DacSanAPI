package main

import (
	"nam/dac_san_api/internal/models"
	"nam/dac_san_api/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	gin.SetMode(gin.ReleaseMode)
	routeConfig := cors.DefaultConfig()
	routeConfig.AllowAllOrigins = true
	routeCors := cors.New(routeConfig)

	router := gin.Default()
	router.Use(routeCors)
	router.GET("/dacsan", routes.DocDacSanJson)
	router.GET("/dacsan/ten=:ten", routes.DocDacSanTheoTenJson)
	router.GET("/dacsan/mota=:mota", routes.DocDacSanTheoMoTaJson)
	router.GET("/dacsan/cachchebien=:cachchebien", routes.DocDacSanTheoCachCheBienJson)
	router.GET("/dacsan/:id", routes.DocDacSanTheoIdJson)
	router.GET("/dacsan/:id/xem", routes.XemDacSan)
	router.GET("/dacsan/size=:size/index=:index", routes.DocTrangDacSanJson)
	router.GET("/dacsan/ten=:ten/size=:size/index=:index", routes.DocTrangDacSanTheoTenJson)
	router.GET("/dacsan/vungmien=:id/size=:size/index=:index", routes.DocTrangDacSanTheoVungMienJson)
	router.GET("/dacsan/muadacsan=:id/size=:size/index=:index", routes.DocTrangDacSanTheoMuaDacSanJson)
	router.GET("/dacsan/nguyenlieu=:id/size=:size/index=:index", routes.DocTrangDacSanTheoNguyenLieuJson)
	router.PUT("/dacsan/xemdanhgia", routes.DocDiemDacSan)
	router.POST("/dacsan/danhgia", routes.DanhGiaDacSan)
	router.PUT("/dacsan/yeuthich", routes.DocYeuThichDacSan)
	router.POST("/dacsan/yeuthich", routes.YeuThichDacSan)
	router.DELETE("/dacsan//yeuthich", routes.BoYeuThichDacSan)
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

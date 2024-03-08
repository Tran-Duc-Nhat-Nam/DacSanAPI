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
	// Đặc sản
	router.GET("/dacsan", routes.DocDacSanJson)
	router.GET("/dacsan/ten=:ten", routes.DocDacSanTheoTenJson)
	router.GET("/dacsan/mota=:mota", routes.DocDacSanTheoMoTaJson)
	router.GET("/dacsan/cachchebien=:cachchebien", routes.DocDacSanTheoCachCheBienJson)
	router.GET("/dacsan/:id", routes.DocDacSanTheoIdJson)
	router.GET("/dacsan/:id/xem", routes.XemDacSan)
	router.GET("/dacsan/size=:size/index=:index", routes.DocTrangDacSanJson)
	router.GET("/dacsan/ten=:ten/size=:size/index=:index", routes.DocTrangDacSanTheoTenJson)
	router.POST("/dacsan/theovungmien/ten=:ten/size=:size/index=:index", routes.DocDacSanTheoVungMien)
	router.POST("/dacsan/theomua/ten=:ten/size=:size/index=:index", routes.DocDacSanTheoMua)
	router.POST("/dacsan/theovungmien/theomua/ten=:ten/size=:size/index=:index", routes.DocDacSanTheoThoiKhong)
	router.GET("/dacsan/vungmien=:id/size=:size/index=:index", routes.DocTrangDacSanTheoVungMienJson)
	router.GET("/dacsan/muadacsan=:id/size=:size/index=:index", routes.DocTrangDacSanTheoMuaDacSanJson)
	router.GET("/dacsan/nguyenlieu=:id/size=:size/index=:index", routes.DocTrangDacSanTheoNguyenLieuJson)
	router.POST("/dacsan", routes.ThemDacSanJson)
	router.PUT("/dacsan", routes.CapNhatDacSanJson)
	router.DELETE("/dacsan", routes.XoaDacSanJson)
	// Đánh giá đặc sản
	router.GET("/danhgia/dacsan=:idDacSan", routes.DocDanhGiaDacSanTheoDacSanJson)
	router.GET("/danhgia/dacsan=:idDacSan/nguoidung=:idNguoiDung", routes.DocDanhGiaDacSanTheoNguoiDungJson)
	router.POST("/danhgia", routes.DanhGiaDacSanJson)
	// Yêu thích đặc sản
	router.GET("/yeuthich/dacsan=:idDacSan/nguoidung=:idNguoiDung", routes.DocYeuThichDacSan)
	router.POST("/yeuthich/dacsan=:idDacSan/nguoidung=:idNguoiDung", routes.YeuThichDacSan)
	router.DELETE("/yeuthich/dacsan=:idDacSan/nguoidung=:idNguoiDung", routes.BoYeuThichDacSan)
	// Nơi bán
	router.GET("/noiban", routes.DocNoiBanJson)
	router.GET("/noiban/:id", routes.DocNoiBanTheoIdJson)
	router.GET("/noiban/:id/dacsan", routes.DocDacSanTheoNoiBanJson)
	router.POST("/noiban", routes.ThemNoiBanJson)
	router.PUT("/noiban", routes.CapNhatNoiBanJson)
	router.DELETE("/noiban", routes.XoaNoiBanJson)
	// Địa chỉ
	router.GET("/tinhthanh", routes.DocTinhThanhJson)
	router.GET("/tinhthanh/:id/quanhuyen", routes.DocQuanHuyenTheoTinhThanhJson)
	router.GET("/quanhuyen/:id/phuongxa", routes.DocPhuongXaTheoQuanHuyenJson)
	router.POST("/tinhthanh", routes.ThemTinhThanhJson)
	router.PUT("/tinhthanh", routes.CapNhatTinhThanhJson)
	router.DELETE("/tinhthanh", routes.XoaTinhThanhJson)
	// Vùng miền
	router.GET("/vungmien", routes.DocVungMienJson)
	router.POST("/vungmien", routes.ThemVungMienJson)
	router.PUT("/vungmien", routes.CapNhatVungMienJson)
	router.DELETE("/vungmien", routes.XoaVungMienJson)
	// Mùa
	router.GET("/muadacsan", routes.DocMuaDacSanJson)
	router.POST("/muadacsan", routes.ThemMuaDacSanJson)
	router.PUT("/muadacsan", routes.CapNhatMuaDacSanJson)
	router.DELETE("/muadacsan", routes.XoaMuaDacSanJson)
	// Nguyên liệu
	router.GET("/nguyenlieu", routes.DocNguyenLieuJson)
	router.POST("/nguyenlieu", routes.ThemNguyenLieuJson)
	router.PUT("/nguyenlieu", routes.CapNhatNguyenLieuJson)
	router.DELETE("/nguyenlieu", routes.XoaNguyenLieuJson)
	router.GET("/diachi/:id", routes.DocDiaChiTheoIdJson)
	// Hình ảnh
	router.GET("/hinhanh/:id", routes.DocHinhAnhTheoIdJson)
	// Người dùng
	router.GET("/nguoidung", routes.DocNguoiDungJson)
	router.GET("/nguoidung/:id", routes.DocNguoiDungTheoIdJson)
	router.POST("/nguoidung", routes.ThemNguoiDungJson)
	router.PUT("/nguoidung", routes.CapNhatNguoiDungJson)
	router.DELETE("/nguoidung", routes.XoaNguoiDungJson)

	router.Run()
}

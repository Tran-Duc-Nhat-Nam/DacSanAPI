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
	router.GET("/dacsan", routes.DocDacSanAPI)
	router.POST("/dacsan", routes.ThemDacSanAPI)
	router.PUT("/dacsan", routes.CapNhatDacSanAPI)
	router.DELETE("/dacsan", routes.XoaDacSanAPI)
	router.GET("/dacsan/:id", routes.DocDacSanTheoIdAPI)
	router.GET("/dacsan/:id/nguoidung=:idNguoiDung", routes.XemDacSan)
	router.GET("/dacsan/:id/chitiet", routes.XemDacSanKhach)
	router.GET("/dacsan/ten=:ten", routes.DocDacSanTheoTenAPI)
	router.GET("/dacsan/mota=:mota", routes.DocDacSanTheoMoTaAPI)
	router.GET("/noiban/:id/dacsan", routes.DocDacSanTheoNoiBanAPI)
	router.GET("/dacsan/cachchebien=:cachchebien", routes.DocDacSanTheoCachCheBienAPI)
	router.GET("/dacsan/size=:size/index=:index", routes.DocTrangDacSanAPI)
	router.GET("/dacsan/ten=:ten/size=:size/index=:index", routes.TimKiemDacSanAPI)
	router.GET("/dacsan/vungmien=:id/size=:size/index=:index", routes.TimKiemDacSanTheoVungMienAPI)
	router.GET("/dacsan/muadacsan=:id/size=:size/index=:index", routes.TimKiemDacSanTheoMuaAPI)
	router.GET("/dacsan/nguyenlieu=:id/size=:size/index=:index", routes.TimKiemDacSanTheoNguyenLieuAPI)
	router.POST("/dacsan/theovungmien/ten=:ten/size=:size/index=:index", routes.TimKiemDacSanTheoDanhSachVungMienAPI)
	router.POST("/dacsan/theomua/ten=:ten/size=:size/index=:index", routes.TimKiemDacSanTheoDanhSachMuaAPI)
	router.POST("/dacsan/theovungmien/theomua/ten=:ten/size=:size/index=:index", routes.TimKiemDacSanTheoMuaVungMienAPI)
	router.POST("/dacsan/theonguyenlieu/ten=:ten/size=:size/index=:index", routes.TimKiemDacSanTheoDanhSachNguyenLieuAPI)
	router.POST("/dacsan/theonguyenlieu/theovungmien/ten=:ten/size=:size/index=:index", routes.TimKiemDacSanTheoNguyenLieuVungMienAPI)
	router.POST("/dacsan/theonguyenlieu/theomua/ten=:ten/size=:size/index=:index", routes.TimKiemDacSanTheoNguyenLieuMuaAPI)
	router.POST("/dacsan/theodieukien/ten=:ten/size=:size/index=:index", routes.TimKiemDacSanTheoDieuKien)
	// Đánh giá đặc sản
	router.GET("/danhgia/dacsan=:idDacSan", routes.DocDanhSachDanhGiaDacSanAPI)
	router.GET("/danhgia/dacsan=:idDacSan/nguoidung=:idNguoiDung", routes.DocDanhGiaDacSanTheoNguoiDungAPI)
	router.POST("/danhgia/dacsan=:idDacSan", routes.DanhGiaDacSanAPI)
	router.PUT("/danhgia/dacsan=:idDacSan", routes.CapNhatDanhGiaDacSanAPI)
	router.DELETE("/danhgia/dacsan=:idDacSan/nguoidung=:idNguoiDung", routes.HuyDanhGiaDacSanAPI)
	// Yêu thích đặc sản
	router.GET("/yeuthich/dacsan/nguoidung=:idNguoiDung", routes.DocDanhSachDacSanYeuThichAPI)
	router.GET("/yeuthich/dacsan/nguoidung=:idNguoiDung/size=:size/index=:index", routes.TimKiemDanhSachDacSanYeuThichAPI)
	router.GET("/yeuthich/dacsan=:idDacSan/nguoidung=:idNguoiDung", routes.DocYeuThichDacSanAPI)
	router.POST("/yeuthich/dacsan=:idDacSan/nguoidung=:idNguoiDung", routes.YeuThichDacSanAPI)
	router.DELETE("/yeuthich/dacsan=:idDacSan/nguoidung=:idNguoiDung", routes.BoYeuThichDacSanAPI)
	// Đánh giá nơi bán
	router.GET("/danhgia/noiban=:idNoiBan", routes.DocDanhSachDanhGiaNoiBanAPI)
	router.GET("/danhgia/noiban=:idNoiBan/nguoidung=:idNguoiDung", routes.DocDanhGiaNoiBanTheoNguoiDungAPI)
	router.POST("/danhgia/noiban=:idNoiBan", routes.DanhGiaNoiBanAPI)
	router.PUT("/danhgia/dacsan=:idDacSan", routes.CapNhatDanhGiaNoiBanAPI)
	router.DELETE("/danhgia/dacsan=:idDacSan/nguoidung=:idNguoiDung", routes.HuyDanhGiaNoiBanAPI)
	// Yêu thích nơi bán
	router.GET("/yeuthich/noiban/nguoidung=:idNguoiDung", routes.DocDanhSachNoiBanYeuThichAPI)
	router.GET("/yeuthich/dacsan/nguoidung=:idNguoiDung/size=:size/index=:index", routes.TimKiemDanhSachNoiBanYeuThichAPI)
	router.GET("/yeuthich/noiban=:idNoiBan/nguoidung=:idNguoiDung", routes.DocYeuThichNoiBanAPI)
	router.POST("/yeuthich/noiban=:idNoiBan/nguoidung=:idNguoiDung", routes.YeuThichNoiBanAPI)
	router.DELETE("/yeuthich/noiban=:idNoiBan/nguoidung=:idNguoiDung", routes.BoYeuThichNoiBanAPI)
	// Nơi bán
	router.GET("/noiban", routes.DocNoiBanAPI)
	router.POST("/noiban", routes.ThemNoiBanAPI)
	router.PUT("/noiban", routes.CapNhatNoiBanAPI)
	router.DELETE("/noiban", routes.XoaNoiBanAPI)
	router.GET("/noiban/:id", routes.DocNoiBanTheoIdAPI)
	router.GET("/noiban/:id/nguoidung=:idNguoiDung", routes.XemNoiBan)
	router.GET("/noiban/:id/chitiet", routes.XemNoiBanKhach)
	router.GET("/noiban/ten=:ten/size=:size/index=:index", routes.TimKiemNoiBanAPI)
	// Địa chỉ
	router.GET("/tinhthanh", routes.DocTinhThanhAPI)
	router.GET("/tinhthanh/:id/quanhuyen", routes.DocQuanHuyenTheoTinhThanhAPI)
	router.GET("/quanhuyen/:id/phuongxa", routes.DocPhuongXaTheoQuanHuyenAPI)
	router.POST("/tinhthanh", routes.ThemTinhThanhAPI)
	router.PUT("/tinhthanh", routes.CapNhatTinhThanhAPI)
	router.DELETE("/tinhthanh", routes.XoaTinhThanhAPI)
	router.GET("/diachi/:id", routes.DocDiaChiTheoIdAPI)
	// Vùng miền
	router.GET("/vungmien", routes.DocVungMienAPI)
	router.POST("/vungmien", routes.ThemVungMienAPI)
	router.PUT("/vungmien", routes.CapNhatVungMienAPI)
	router.DELETE("/vungmien", routes.XoaVungMienAPI)
	// Mùa
	router.GET("/muadacsan", routes.DocMuaDacSanAPI)
	router.POST("/muadacsan", routes.ThemMuaDacSanAPI)
	router.PUT("/muadacsan", routes.CapNhatMuaDacSanAPI)
	router.DELETE("/muadacsan", routes.XoaMuaDacSanAPI)
	// Nguyên liệu
	router.GET("/nguyenlieu", routes.DocNguyenLieuAPI)
	router.POST("/nguyenlieu", routes.ThemNguyenLieuAPI)
	router.PUT("/nguyenlieu", routes.CapNhatNguyenLieuAPI)
	router.DELETE("/nguyenlieu", routes.XoaNguyenLieuAPI)
	router.GET("/nguyenlieu/ten=:ten/size=:size/index=:index", routes.DocNguyenLieuTheoTenAPI)
	// Hình ảnh
	router.GET("/hinhanh/:id", routes.DocHinhAnhTheoIdAPI)
	router.GET("/dacsan/:id/hinhanh", routes.DocHinhAnhDacSanAPI)
	// Người dùng
	router.GET("/nguoidung", routes.DocNguoiDungAPI)
	router.POST("/nguoidung", routes.ThemNguoiDungAPI)
	router.PUT("/nguoidung", routes.CapNhatNguoiDungAPI)
	router.DELETE("/nguoidung", routes.XoaNguoiDungAPI)
	router.GET("/nguoidung/:id", routes.DocNguoiDungTheoIdAPI)
	router.GET("/nguoidung/ten=:ten/size=:size/index=:index", routes.TimKiemNguoiDungAPI)
	// Lịch sử xem
	router.GET("/lichsu/dacsan/nguoidung=:idNguoiDung", routes.DocDanhSachDacSanDaXemAPI)
	router.GET("/lichsu/noiban/nguoidung=:idNguoiDung", routes.DocDanhSachNoiBanDaXemAPI)

	router.Run()
}

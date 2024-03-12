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
	router.POST("/dacsan", routes.ThemDacSanJson)
	router.PUT("/dacsan", routes.CapNhatDacSanJson)
	router.DELETE("/dacsan", routes.XoaDacSanJson)
	router.GET("/dacsan/:id", routes.DocDacSanTheoIdJson)
	router.GET("/dacsan/:id/xem", routes.XemDacSan)
	router.GET("/dacsan/ten=:ten", routes.DocDacSanTheoTenJson)
	router.GET("/dacsan/mota=:mota", routes.DocDacSanTheoMoTaJson)
	router.GET("/noiban/:id/dacsan", routes.DocDacSanTheoNoiBanAPI)
	router.GET("/dacsan/cachchebien=:cachchebien", routes.DocDacSanTheoCachCheBienJson)
	router.GET("/dacsan/size=:size/index=:index", routes.DocTrangDacSanJson)
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
	// Yêu thích đặc sản
	router.GET("/yeuthich/dacsan=:idDacSan/nguoidung=:idNguoiDung", routes.DocYeuThichDacSanAPI)
	router.POST("/yeuthich/dacsan=:idDacSan/nguoidung=:idNguoiDung", routes.YeuThichDacSanAPI)
	router.DELETE("/yeuthich/dacsan=:idDacSan/nguoidung=:idNguoiDung", routes.BoYeuThichDacSanAPI)
	// Đánh giá nơi bán
	router.GET("/danhgia/noiban=:idNoiBan", routes.DocDanhSachDanhGiaNoiBanAPI)
	router.GET("/danhgia/noiban=:idNoiBan/nguoidung=:idNguoiDung", routes.DocDanhGiaNoiBanTheoNguoiDungAPI)
	router.POST("/danhgia/noiban=:idNoiBan", routes.DanhGiaNoiBanAPI)
	// Yêu thích nơi bán
	router.GET("/yeuthich/noiban=:idNoiBan/nguoidung=:idNguoiDung", routes.DocYeuThichNoiBanAPI)
	router.POST("/yeuthich/noiban=:idNoiBan/nguoidung=:idNguoiDung", routes.YeuThichNoiBanAPI)
	router.DELETE("/yeuthich/noiban=:idNoiBan/nguoidung=:idNguoiDung", routes.BoYeuThichNoiBanAPI)
	// Nơi bán
	router.GET("/noiban", routes.DocNoiBanAPI)
	router.POST("/noiban", routes.ThemNoiBanAPI)
	router.PUT("/noiban", routes.CapNhatNoiBanAPI)
	router.DELETE("/noiban", routes.XoaNoiBanAPI)
	router.GET("/noiban/:id", routes.DocNoiBanTheoIdAPI)
	router.GET("/noiban/ten=:ten/size=:size/index=:index", routes.TimKiemNoiBanAPI)
	// Địa chỉ
	router.GET("/tinhthanh", routes.DocTinhThanhJson)
	router.GET("/tinhthanh/:id/quanhuyen", routes.DocQuanHuyenTheoTinhThanhJson)
	router.GET("/quanhuyen/:id/phuongxa", routes.DocPhuongXaTheoQuanHuyenJson)
	router.POST("/tinhthanh", routes.ThemTinhThanhJson)
	router.PUT("/tinhthanh", routes.CapNhatTinhThanhJson)
	router.DELETE("/tinhthanh", routes.XoaTinhThanhJson)
	router.GET("/diachi/:id", routes.DocDiaChiTheoIdJson)
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
	router.GET("/nguyenlieu/ten=:ten/size=:size/index=:index", routes.DocNguyenLieuTheoTenJson)
	// Hình ảnh
	router.GET("/hinhanh/:id", routes.DocHinhAnhTheoIdJson)
	router.GET("/dacsan/:id/hinhanh", routes.DocHinhAnhDacSanJson)
	// Người dùng
	router.GET("/nguoidung", routes.DocNguoiDungAPI)
	router.POST("/nguoidung", routes.ThemNguoiDungAPI)
	router.PUT("/nguoidung", routes.CapNhatNguoiDungAPI)
	router.DELETE("/nguoidung", routes.XoaNguoiDungAPI)
	router.GET("/nguoidung/:id", routes.DocNguoiDungTheoIdAPI)
	router.GET("/nguoidung/ten=:ten/size=:size/index=:index", routes.TimKiemNguoiDungAPI)

	router.Run()
}

package main

import (
	"context"
	"database/sql"
	"nam/dac_san_api/internal/models"
	"net"
	"os"

	// "encoding/json"
	"strconv"

	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
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
	nguoiDung, err := models.DocDacSanTheoIdCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, nguoiDung)
}

func DocNoiBanJson(c *gin.Context) {
	dsNoiBan, _ := models.DocNoiBanCSDL()
	c.IndentedJSON(http.StatusOK, dsNoiBan)
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

func DocNguoiDungTheoIdJson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err.Error())
	}
	nguoiDung, err := models.DocNguoiDungCSDL(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, nguoiDung)
}

func connectWithConnector() (*sql.DB, error) {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Fatal Error in connect_connector.go: %s environment variable not set.", k)
		}
		return v
	}
	// Note: Saving credentials in environment variables is convenient, but not
	// secure - consider a more secure solution such as
	// Cloud Secret Manager (https://cloud.google.com/secret-manager) to help
	// keep passwords and other secrets safe.
	var (
		dbUser                 = mustGetenv("root")                                          // e.g. 'my-db-user'
		dbPwd                  = mustGetenv("nhatnam2002")                                   // e.g. 'my-db-password'
		dbName                 = mustGetenv("mysql")                                         // e.g. 'my-database'
		instanceConnectionName = mustGetenv("eternal-insight-410902:asia-east1:dac-san-api") // e.g. 'project:region:instance'
		usePrivate             = os.Getenv("")
	)

	d, err := cloudsqlconn.NewDialer(context.Background())
	if err != nil {
		return nil, fmt.Errorf("cloudsqlconn.NewDialer: %w", err)
	}
	var opts []cloudsqlconn.DialOption
	if usePrivate != "" {
		opts = append(opts, cloudsqlconn.WithPrivateIP())
	}
	mysql.RegisterDialContext("cloudsqlconn",
		func(ctx context.Context, addr string) (net.Conn, error) {
			return d.Dial(ctx, instanceConnectionName, opts...)
		})

	dbURI := fmt.Sprintf("%s:%s@cloudsqlconn(localhost:3306)/%s?parseTime=true",
		dbUser, dbPwd, dbName)

	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	return dbPool, nil
}

func main() {
	models.ConnectDatabase()

	router := gin.Default()
	router.GET("/dacsan", DocDacSanJson)
	router.GET("/dacsan/:id", DocDacSanTheoIdJson)
	router.GET("/noiban", DocNoiBanJson)
	router.POST("/noiban/them", ThemNoiBanJson)
	router.POST("/noiban/capnhat", CapNhatNoiBanJson)
	router.POST("/noiban/xoa", XoaNoiBanJson)
	router.GET("/tinhthanh", DocTinhThanhJson)
	router.POST("/tinhthanh/them", ThemTinhThanhJson)
	router.POST("/tinhthanh/capnhat", CapNhatTinhThanhJson)
	router.POST("/tinhthanh/xoa", XoaTinhThanhJson)
	router.GET("/vungmien", DocVungMienJson)
	router.POST("/vungmien/them", ThemVungMienJson)
	router.POST("/vungmien/capnhat", CapNhatVungMienJson)
	router.POST("/vungmien/xoa", XoaVungMienJson)
	router.GET("/muadacsan", DocMuaDacSanJson)
	router.POST("/muadacsan/them", ThemMuaDacSanJson)
	router.POST("/muadacsan/capnhat", CapNhatMuaDacSanJson)
	router.POST("/muadacsan/xoa", XoaMuaDacSanJson)
	router.GET("/nguyenlieu", DocNguyenLieuJson)
	router.POST("/nguyenlieu/them", ThemNguyenLieuJson)
	router.POST("/nguyenlieu/capnhat", CapNhatNguyenLieuJson)
	router.POST("/nguyenlieu/xoa", XoaNguyenLieuJson)
	router.GET("/diachi/:id", DocDiaChiTheoIdJson)
	router.GET("/hinhanh/:id", DocHinhAnhTheoIdJson)
	router.GET("/nguoidung/:id", DocNguoiDungTheoIdJson)
	router.Run("localhost:8080")
}

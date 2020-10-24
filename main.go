package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" //直接的な記述が無いが、インポートしたいものに対しては"_"を頭につける決まり
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")

	dbInit()

	//一覧
	router.GET("/", func(c *gin.Context) {
		pointList := dbGetAll()
		c.HTML(200, "index.html", gin.H{"pointList": pointList})
	})

	//登録
	router.POST("/new", func(c *gin.Context) {
		var form Point
		// ここがバリデーション部分
		if err := c.Bind(&form); err != nil {
			pointList := dbGetAll()
			c.HTML(http.StatusBadRequest, "index.html", gin.H{"pointList": pointList, "err": err})
			c.Abort()
		} else {
			sa := c.PostForm("sotenA")
			sb := c.PostForm("sotenB")
			sc := c.PostForm("sotenC")
			sd := c.PostForm("sotenD")
			sotenA, err := strconv.Atoi(sa)
			sotenB, err := strconv.Atoi(sb)
			sotenC, err := strconv.Atoi(sc)
			sotenD, err := strconv.Atoi(sd)
			if err != nil {
				panic(err)
			}
			dbInsert(sotenA, sotenB, sotenC, sotenD)
			c.Redirect(302, "/")
		}
	})

	//投稿詳細
	router.GET("/detail/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		point := dbGetOne(id)
		c.HTML(200, "detail.html", gin.H{"point": point})
	})

	//更新
	router.POST("/update/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		sa := c.PostForm("sotenA")
		sb := c.PostForm("sotenB")
		sc := c.PostForm("sotenC")
		sd := c.PostForm("sotenD")
		sotenA, err := strconv.Atoi(sa)
		sotenB, err := strconv.Atoi(sb)
		sotenC, err := strconv.Atoi(sc)
		sotenD, err := strconv.Atoi(sd)
		if err != nil {
			panic(err)
		}
		dbUpdate(id, sotenA, sotenB, sotenC, sotenD)
		c.Redirect(302, "/")
	})

	//削除確認
	router.GET("/delete_check/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		point := dbGetOne(id)
		c.HTML(200, "delete.html", gin.H{"point": point})
	})

	//削除
	router.POST("/delete/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		dbDelete(id)
		c.Redirect(302, "/")

	})

	router.Run()
}

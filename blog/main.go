package main

func main() {
	// gin.SetMode(gin.ReleaseMode) // 或者 gin.DebugMode
	// db := initDB()
	//server := gin.Default()
	// server := initWebServer()
	// initUserHdl(db, server)
	server := InitWebServer()
	// server.GET("/hello", func(ctx *gin.Context) {
	// 	ctx.String(http.StatusOK, "hello，启动成功了！")
	// })
	server.Run(":8000")
}

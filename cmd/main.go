package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mygo/form"
	"mygo/pkg/service"
	"mygo/pkg/utils"
	"mygo/pkg/utils/log"
	"net/http"
	"time"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var JwtService = service.NewJwtService()
var Config = utils.NewConfigService()
var Logger *log.Log

type Server struct {
	//Config     *utils.Config
	BaseServer *gin.Engine
}

func NewServer(cfg *utils.Config) *Server {
	return &Server{
		//Config:     cfg,
		BaseServer: gin.Default(),
	}
}

func main() {
	server := NewServer(&utils.Config{})
	r := server.BaseServer

	//读取配置文件，并初始化
	if err := Config.ReadYaml("/Users/zhuly/git/duapp/mygo/config/local/config.yaml"); err != nil {
		//配置错误
	}

	Logger = log.NewLogService(Config.Logger)

	authRouter := r.Group("/", JwtService.ParseToken)
	{
		authRouter.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		authRouter.GET("checkLogin", CheckLogin)
	}

	commonRouter := r.Group("/common")
	{
		commonRouter.GET("login", LoginHandler)
		commonRouter.GET("/my", func(context *gin.Context) {
			fmt.Println(2)
		})
	}

	s := &http.Server{
		Addr:         ":8088",
		Handler:      r,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	s.ListenAndServe()
	//http.ListenAndServe(":8080", r)
	//r.Run()
}

func CheckLogin(ctx *gin.Context) {
	JwtInfo := JwtService.GetJwtInfo(ctx)
	if JwtInfo.Username == "zhulingyu" {
		ctx.JSON(http.StatusOK, Response{
			Code: 200,
			Msg:  "success",
			Data: JwtInfo.Username,
		})
	}
	Logger.Info("check---", Response{
		Code: 200,
		Msg:  "success",
		Data: JwtInfo.Username,
	})
}

func LoginHandler(ctx *gin.Context) {
	var user form.User
	if err := ctx.ShouldBindQuery(&user); err != nil {
		ctx.JSON(http.StatusOK, Response{
			Code: 500,
			Msg:  "failed",
			Data: err,
		})
		return
	}

	token := service.NewJwtService().GenToken(ctx, user)

	ctx.JSON(http.StatusOK, Response{
		Code: 200,
		Msg:  "success",
		Data: token,
	})
}

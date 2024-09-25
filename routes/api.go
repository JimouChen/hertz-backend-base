package routes

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
	"herrz-backend-base/comm"
	"herrz-backend-base/service"
)

// LogIPMiddleware 自定义中间件记录 IP 地址
func LogIPMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		comm.Logger.Info().Msgf("Client IP: %s", c.RemoteAddr().String()) // 打印 IP 地址
		c.Next(ctx)                                                       // 执行下一个中间件或处理函数
	}
}

func Init() *server.Hertz {
	cfgLoader := comm.CfgLoader
	env := cfgLoader.GetString("app.env") + ".url"
	addr := cfgLoader.GetString(env)
	hz := server.New(server.WithHostPorts(addr))

	// 配置 CORS 中间件，允许跨域访问
	hz.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 允许所有来源
		AllowMethods:     []string{"*"}, // 允许的请求方法
		AllowHeaders:     []string{"*"}, // 允许的请求头
		AllowCredentials: true,          // 允许凭证
	}))
	hz.Use(LogIPMiddleware())
	v1 := hz.Group("/v1")

	v1.GET("/demo", service.SearchOne)
	v1.GET("/fetch_all", service.SearchAll)
	v1.POST("/fetch_all", service.SearchAllWithPost)
	v1.PUT("/demo", service.Update)
	v1.POST("/demo", service.Insert)

	// 启动服务器
	return hz
}

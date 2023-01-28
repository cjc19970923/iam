package crmapiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/iam/internal/crmapiserver/controller/v1/nodify"
	"github.com/marmotedu/iam/internal/crmapiserver/middleware"
	"github.com/marmotedu/iam/internal/crmapiserver/store/mysql"
)

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) {
	g.Use(middleware.GetAppId)

}

func installController(g *gin.Engine) {

	storeIns, _ := mysql.GetMySQLFactory(nil)
	nodifyController := nodify.NewNodifyController(storeIns)

	v1 := g.Group("/v1")
	v1.POST("/nodify", nodifyController.Nodify)

}

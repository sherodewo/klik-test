package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ApiRoute(e *echo.Echo, db *gorm.DB) {
	//aGroup := e.Group("api/v1")
	//telcoScoreController := config.InjectTelcoScoreController(db)
	//checkController := config.InjectcheckController(db)
	//kmbController := config.InjectKmbcheckController(db)
	//wgController := config.InjectWgcheckController(db)
	//telcoGroup := aGroup.Group("/score")
	//{
	//	telcoGroup.POST("/credit/:phoneNumber", telcoScoreController.CreditScore)
	//	telcoGroup.POST("/credit/:phoneNumber/limit", telcoScoreController.CreditScoreLimit)
	//	telcoGroup.GET("/credit/detail/:id", telcoScoreController.Detail)
	//	telcoGroup.POST("/experian", telcoScoreController.Experian)
	//	telcoGroup.POST("/token", telcoScoreController.GetToken)
	//	telcoGroup.POST("/pickle", telcoScoreController.InternalScoring)
	//}
	//
	//{
	//	checkGroup := aGroup.Group("/check")
	//	checkGroup.POST("/idx", checkController.Scoring)
	//	checkGroup.GET("/detail/:id", checkController.Detail)
	//}
	//
	//{
	//	kmbGroup := aGroup.Group("/check/kmb")
	//	kmbGroup.POST("/idx", kmbController.Scoring)
	//}
	//
	//{
	//	wgGroup := aGroup.Group("/check/wg")
	//	wgGroup.POST("/idx", wgController.Scoring)
	//}

}

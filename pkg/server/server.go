package main

import(
	// import gin library
	"github.com/gin-gonic/gin"

	// import sample API packages
	"github.com/miraikeitai2020/backend-summer-vacation/pkg/server/controller"
)

func initializeController() controller.Controller {
	return controller.Controller{}
}

func setupRooter(ctrl controller.Controller) *gin.Engine {
	router := gin.Default()
	// API Handlers List
	router.GET("/", ctrl.HelloWorld)
	router.POST("/sayhello", ctrl.SayHello)
	// 以下， 課題の進行状況に応じてコメントアウトを外す
	router.GET("/task1", ctrl.Task1)
	// router.POST("/task2", ctrl.Task2)
	// router.POST("/task4", ctrl.SignUp)
	// router.POST("/task4", ctrl.SignIn)
	return router
}

func main() {
	ctrl := initializeController()
	router := setupRooter(ctrl)
	err := router.Run(":8080")
	if err != nil {
		panic("server Painc")
	}
}
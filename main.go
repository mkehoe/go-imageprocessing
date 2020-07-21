package main

import (
	"github.com/astaxie/beego"
	"github.com/mkehoe/go-imageprocessing/controllers"
)

func main() {
	beego.Router("/images/stats", &controllers.ImageProcessingController{}, "get:GetStats")
	beego.Router("/images/reset_stats", &controllers.ImageProcessingController{}, "post:ResetStats")
	beego.Router("/images/tile", &controllers.ImageProcessingController{}, "post:TileImage")
	beego.Router("/images/resize", &controllers.ImageProcessingController{}, "post:ResizeImage")
	beego.Run(":8080")
}

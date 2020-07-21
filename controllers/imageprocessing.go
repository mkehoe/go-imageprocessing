package controllers

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"time"

	"github.com/astaxie/beego"
	"github.com/mkehoe/go-imageprocessing/models"
	"github.com/nfnt/resize"
)

type ImageProcessingController struct {
	beego.Controller
}

func (this *ImageProcessingController) GetStats() {
	// res := struct{ Timestamps []*models.Timestamp }
	var stats = models.GetTimestamps()
	fmt.Println(stats.TileTimestamps)
	fmt.Println("Timestamp Count = ", len(stats.TileTimestamps))
	this.Data["json"] = stats
	this.ServeJSON()
}

func (this *ImageProcessingController) ResetStats() {
	models.Reset()
}

func (this *ImageProcessingController) TileImage() {
	timestamp := models.Timestamp{}
	start := time.Now()

	file, _, _ := this.GetFile("Image")

	timestamp.T1 = time.Since(start).Milliseconds()

	img, _, _ := image.Decode(file)

	timestamp.T2 = time.Since(start).Milliseconds()

	width := img.Bounds().Max.X / 2
	height := img.Bounds().Max.Y / 2
	bgImg := image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X, img.Bounds().Max.Y))

	timestamp.T3 = time.Since(start).Milliseconds()

	resizedImg := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
	draw.Draw(bgImg, image.Rect(0, 0, width, height), resizedImg, image.ZP, draw.Src)
	draw.Draw(bgImg, image.Rect(width, height, img.Bounds().Max.X, img.Bounds().Max.Y), resizedImg, image.ZP, draw.Src)
	draw.Draw(bgImg, image.Rect(0, height, width, img.Bounds().Max.Y), resizedImg, image.ZP, draw.Src)
	draw.Draw(bgImg, image.Rect(width, 0, img.Bounds().Max.X, height), resizedImg, image.ZP, draw.Src)

	timestamp.T4 = time.Since(start).Milliseconds()
	this.Ctx.ResponseWriter.Header().Set("Content-Type", "image/jpeg") // <-- set the content-type header
	jpeg.Encode(this.Ctx.ResponseWriter, bgImg, nil)

	timestamp.T5 = time.Since(start).Milliseconds()

	models.AddTileTimestamp(timestamp)

	fmt.Println("Tile", timestamp)
}

func (this *ImageProcessingController) ResizeImage() {
	timestamp := models.Timestamp{}
	start := time.Now()

	file, _, _ := this.GetFile("Image")

	timestamp.T1 = time.Since(start).Milliseconds()

	rwidth, _ := this.GetInt("ResizeWidth")
	rheight, _ := this.GetInt("ResizeHeight")
	img, _, _ := image.Decode(file)

	timestamp.T2 = time.Since(start).Milliseconds()

	timestamp.T3 = time.Since(start).Milliseconds()
	resizedImg := resize.Resize(uint(rwidth), uint(rheight), img, resize.Lanczos3)

	timestamp.T4 = time.Since(start).Milliseconds()

	this.Ctx.ResponseWriter.Header().Set("Content-Type", "image/jpeg") // <-- set the content-type header
	jpeg.Encode(this.Ctx.ResponseWriter, resizedImg, nil)

	timestamp.T5 = time.Since(start).Milliseconds()

	models.AddResizeTimestamp(timestamp)

	fmt.Println("Resize", timestamp)
}

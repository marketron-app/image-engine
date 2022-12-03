package main

import "C"
import (
	"fmt"
	"marketron-image-engine/crawler"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Start")

	seleniumCrawler := crawler.Crawler{URL: "https://www.higroup.si"}
	seleniumCrawler.GetScreenshot()

	fmt.Println("End: " + time.Since(start).String())

	//cardPath := filepath.Join("1.png")
	//// read image
	//card := gocv.IMRead(cardPath, gocv.IMReadColor)
	//if card.Empty() {
	//	fmt.Printf("Failed to read image: %s\n", cardPath)
	//	os.Exit(1)
	//}
	//
	//points := []image.Point{
	//	{X: 0, Y: 0},
	//	{X: 1125, Y: 0},
	//	{X: 0, Y: 2436},
	//	{X: 1125, Y: 2436},
	//}
	//
	//origImg := gocv.NewPointVectorFromPoints(points)
	//
	//points = []image.Point{
	//	{1955, 836},
	//	{2743, 1623},
	//	{137, 2299},
	//	{898, 3166},
	//}
	//newImg := gocv.NewPointVectorFromPoints(points)
	//
	//transform := gocv.GetPerspectiveTransform(origImg, newImg)
	//
	//perspective := gocv.NewMat()
	//width := 3456
	//height := 3456
	//gocv.WarpPerspective(card, &perspective, transform, image.Point{X: width, Y: height})
	//
	//outPath := filepath.Join("card_perspective.jpg")
	//if ok := gocv.IMWrite(outPath, perspective); !ok {
	//	fmt.Printf("Failed to write image: %s\n")
	//	os.Exit(1)
	//}
}

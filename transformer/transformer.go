package transformer

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/draw"
	"image/png"
	requestbody "marketron-image-engine/api/request-body"
	"os"
)

type Transformer struct {
	WebsiteImage      []byte
	TemplateImage     []byte
	MappedCoordinates []requestbody.TemplateCoordinate
	FileName          string
}

func (t *Transformer) Create() error {
	// read image
	fmt.Println(t.MappedCoordinates)
	screenshot, err := gocv.IMDecode(t.WebsiteImage, gocv.IMReadUnchanged)
	if screenshot.Empty() || err != nil {
		fmt.Printf("Failed to decode screenshot image")
		return err
	}

	template, err := gocv.IMDecode(t.TemplateImage, gocv.IMReadUnchanged)
	if screenshot.Empty() || err != nil {
		fmt.Printf("Failed to decode template image")
		return err
	}

	// 0 ... height
	// 1 ... width
	screenshotSize := screenshot.Size()
	templateSize := template.Size()

	points := []image.Point{
		{X: 0, Y: 0},
		{X: 0, Y: screenshotSize[0]},
		{X: screenshotSize[1], Y: screenshotSize[0]},
		{X: screenshotSize[1], Y: 0},
	}

	origImg := gocv.NewPointVectorFromPoints(points)

	fmt.Println(t.MappedCoordinates)
	points = []image.Point{
		{t.MappedCoordinates[0].X, t.MappedCoordinates[0].Y},
		{t.MappedCoordinates[1].X, t.MappedCoordinates[1].Y},
		{t.MappedCoordinates[2].X, t.MappedCoordinates[2].Y},
		{t.MappedCoordinates[3].X, t.MappedCoordinates[3].Y},
	}
	newImg := gocv.NewPointVectorFromPoints(points)

	transform := gocv.GetPerspectiveTransform(origImg, newImg)

	perspective := gocv.NewMat()
	height := templateSize[0]
	width := templateSize[1]
	gocv.WarpPerspective(screenshot, &perspective, transform, image.Point{X: width, Y: height})

	screenshotImage, _ := perspective.ToImage()
	templateImage, _ := template.ToImage()
	r := image.Rectangle{Min: image.Point{}, Max: image.Point{X: width, Y: height}}
	rgba := image.NewRGBA(r)
	draw.Draw(rgba, screenshotImage.Bounds(), screenshotImage, image.Point{}, draw.Over)
	draw.Draw(rgba, templateImage.Bounds(), templateImage, image.Point{}, draw.Over)

	out, err := os.Create("./output.png")
	if err != nil {
		fmt.Println(err)
	}
	png.Encode(out, rgba)

	out, err = os.Create("./template.png")
	if err != nil {
		fmt.Println(err)
	}
	png.Encode(out, templateImage)

	return nil
}

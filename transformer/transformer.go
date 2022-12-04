package transformer

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	requestbody "marketron-image-engine/api/request-body"
	"path/filepath"
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
	screenshot, err := gocv.IMDecode(t.WebsiteImage, gocv.IMReadColor)
	if screenshot.Empty() || err != nil {
		fmt.Printf("Failed to decode screenshot image")
		return err
	}

	template, err := gocv.IMDecode(t.TemplateImage, gocv.IMReadColor)
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

	outPath := filepath.Join("card_perspective.jpg")
	if ok := gocv.IMWrite(outPath, perspective); !ok {
		fmt.Printf("Failed to write image")
		return err
	}

	return nil
}

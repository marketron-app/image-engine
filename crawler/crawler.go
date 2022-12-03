package crawler

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"log"
	"os"
)

type Crawler struct {
	URL string
}

func (c *Crawler) GetScreenshot() {

	// create context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// capture screenshot of an element
	var buf []byte

	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx,
		chromedp.Emulate(chromedp.Device(device.IPhone7landscape)),
		chromedp.Navigate(c.URL),
		chromedp.FullScreenshot(&buf, 90),
	); err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("fullScreenshot.png", buf, 0o644); err != nil {
		log.Fatal(err)
	}

	log.Printf("wrote elementScreenshot.png and fullScreenshot.png")

	fmt.Println("WE FINISH")

}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Note: chromedp.FullScreenshot overrides the device's emulation settings. Use
// device.Reset to reset the emulation and viewport settings.
func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}

package crawler

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"log"
	"time"
)

type Crawler struct {
	URL            string
	ViewportWidth  int64
	ViewportHeight int64
	IsMobile       bool
}

func (c *Crawler) GetScreenshot() (error, []byte) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var buf []byte

	if err := chromedp.Run(ctx,
		chromedp.Emulate(c.generateDevice()),
		chromedp.Navigate(c.URL),
		chromedp.WaitReady("body"),
		chromedp.CaptureScreenshot(&buf),
	); err != nil {
		log.Println(err)
		return err, nil
	}

	return nil, buf
}

func (c *Crawler) generateDevice() chromedp.Device {
	return device.Info{
		Name:      "Custom",
		UserAgent: "marketron",
		Width:     c.ViewportWidth,
		Height:    c.ViewportHeight,
		Scale:     0,
		Landscape: c.ViewportWidth > c.ViewportHeight,
		Mobile:    c.IsMobile,
		Touch:     c.IsMobile,
	}
}

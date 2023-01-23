package crawler

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	defaultCrawlerTimeout = 10
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

	contextTimeoutSeconds, err := strconv.Atoi(os.Getenv("CRAWLER_TIMEOUT"))
	if err != nil || contextTimeoutSeconds == 0 {
		contextTimeoutSeconds = defaultCrawlerTimeout
	}

	ctx, cancel = context.WithTimeout(ctx, time.Duration(contextTimeoutSeconds)*time.Second)
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

package crawler

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

type Crawler struct {
	URL            string
	ViewportWidth  int64
	ViewportHeight int64
}

func (c *Crawler) GetScreenshot() (error, []byte) {
	start := time.Now()
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	var buf []byte

	if err := chromedp.Run(ctx,
		chromedp.EmulateViewport(c.ViewportWidth, c.ViewportHeight),
		chromedp.Navigate(c.URL),
		chromedp.CaptureScreenshot(&buf),
	); err != nil {
		log.Println(err)
		return err, nil
	}

	fmt.Println("End: " + time.Since(start).String())
	return nil, buf
}

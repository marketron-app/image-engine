package crawler

import (
	"context"
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
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
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

	return nil, buf
}

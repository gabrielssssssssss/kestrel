package website2screenshot

import (
	"context"
	"time"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/emulation"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/rpcc"
)

func NavigateToPage(url string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	devt := devtool.New("http://127.0.0.1:9222")
	pt, err := devt.Get(ctx, devtool.Page)
	if err != nil {
		pt, err = devt.Create(ctx)
		if err != nil {
			return nil, err
		}
	}

	conn, err := rpcc.DialContext(ctx, pt.WebSocketDebuggerURL)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := cdp.NewClient(conn)

	loadEvent, err := c.Page.LoadEventFired(ctx)
	if err != nil {
		return nil, err
	}
	defer loadEvent.Close()

	if err = c.Page.Enable(ctx); err != nil {
		return nil, err
	}

	_, err = c.Page.Navigate(ctx, page.NewNavigateArgs(url))
	if err != nil {
		return nil, err
	}

	_, err = loadEvent.Recv()
	if err != nil {
		return nil, err
	}

	err = c.Emulation.SetDeviceMetricsOverride(ctx, emulation.NewSetDeviceMetricsOverrideArgs(1920, 1080, 1.0, false))
	if err != nil {
		return nil, err
	}

	time.Sleep(3 * time.Second)

	screenshotArgs := page.NewCaptureScreenshotArgs().
		SetFormat("jpeg").
		SetQuality(90).
		SetFromSurface(true)
	screenshot, err := c.Page.CaptureScreenshot(ctx, screenshotArgs)
	if err != nil {
		return nil, err
	}

	return screenshot.Data, nil
}

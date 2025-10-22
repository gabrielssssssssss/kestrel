package lighthouse

import (
	"context"

	"github.com/gabrielssssssssss/kestrel/internal/config"
	"google.golang.org/api/option"
	"google.golang.org/api/pagespeedonline/v5"
)

func Lighthouse(url string) (*pagespeedonline.PagespeedApiPagespeedResponseV5, error) {
	ctx := context.Background()

	service, err := pagespeedonline.NewService(ctx, option.WithAPIKey(config.GetConfig("LIGHTHOUSE_API_KEY")))
	if err != nil {
		return nil, err
	}

	call := service.Pagespeedapi.Runpagespeed(url)
	call = call.Strategy("desktop")
	call = call.Fields("lighthouseResult.categories.performance.score")

	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	return response, nil
}

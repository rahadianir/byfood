package urlcleaner

import (
	"byfood-app/internal/core"
	"context"
)

type URLCleanerLogic struct {
	deps *core.Dependency
}

func NewURLCleanerLogic(deps *core.Dependency) *URLCleanerLogic {
	return &URLCleanerLogic{
		deps: deps,
	}
}

func (logic *URLCleanerLogic) CleanURL(ctx context.Context, url string, operation string) (string, error) {
	
	return "", nil
}

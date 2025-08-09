package urlcleaner

import "context"

type URLCleanerLogicInterface interface {
	CleanURL(ctx context.Context, link string, operation string) (string, error)
}

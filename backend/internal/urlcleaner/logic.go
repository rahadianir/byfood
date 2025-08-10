package urlcleaner

import (
	"byfood-app/internal/core"
	"byfood-app/internal/pkg/xerrors"
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"strings"
)

type URLCleanerLogic struct {
	deps *core.Dependency
}

func NewURLCleanerLogic(deps *core.Dependency) *URLCleanerLogic {
	return &URLCleanerLogic{
		deps: deps,
	}
}

func (logic *URLCleanerLogic) CleanURL(ctx context.Context, link string, operation string) (string, error) {
	u, err := url.Parse(link)
	if err != nil {
		logic.deps.Logger.ErrorContext(ctx, "failed to parse given url", slog.Any("error", err))
		return "", err
	}

	switch strings.ToLower(operation) {
	case "all":
		u, err := logic.canonicalCleanUp(ctx, u)
		if err != nil {
			logic.deps.Logger.ErrorContext(ctx, "failed to do canonical clean up to url", slog.Any("error", err))
			return "", err
		}

		u, err = logic.redirectionCleanUp(ctx, u)
		if err != nil {
			logic.deps.Logger.ErrorContext(ctx, "failed to do redirection clean up to url", slog.Any("error", err))
			return "", err
		}

		return u.String(), nil

	case "canonical":
		u, err := logic.canonicalCleanUp(ctx, u)
		if err != nil {
			logic.deps.Logger.ErrorContext(ctx, "failed to do canonical clean up to url", slog.Any("error", err))
			return "", err
		}

		return u.String(), nil

	case "redirection":
		u, err = logic.redirectionCleanUp(ctx, u)
		if err != nil {
			logic.deps.Logger.ErrorContext(ctx, "failed to do redirection clean up to url", slog.Any("error", err))
			return "", err
		}

		return u.String(), nil

	default:
		return "", xerrors.NewClientError(fmt.Errorf("invalid operation key"))
	}

}

func (logic *URLCleanerLogic) canonicalCleanUp(ctx context.Context, link *url.URL) (*url.URL, error) {
	urlString := fmt.Sprintf("%s://%s%s", link.Scheme, link.Host, link.EscapedPath())
	u, err := url.Parse(urlString)
	if err != nil {
		logic.deps.Logger.ErrorContext(ctx, "failed to parse canonical clean up url", slog.Any("error", err))
		return nil, err
	}

	return u, nil
}

func (logic *URLCleanerLogic) redirectionCleanUp(ctx context.Context, link *url.URL) (*url.URL, error) {
	// "fail" fast
	if !strings.Contains(strings.ToLower(link.Host), "byfood") {
		return nil, xerrors.NewClientError(fmt.Errorf("invalid domain url"))
	}

	link.Host = "www.byfood.com"

	urlString := strings.ToLower(link.String())
	u, err := url.Parse(urlString)
	if err != nil {
		logic.deps.Logger.ErrorContext(ctx, "failed to parse redirection clean up url", slog.Any("error", err))
		return nil, err
	}

	return u, nil
}

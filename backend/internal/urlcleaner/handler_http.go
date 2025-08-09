package urlcleaner

import (
	"byfood-app/internal/core"
	"byfood-app/internal/model"
	"byfood-app/internal/pkg/xerrors"
	"byfood-app/internal/pkg/xhttp"
	"log/slog"
	"net/http"
)

type URLCLeanerHandler struct {
	deps  *core.Dependency
	logic URLCleanerLogicInterface
}

func NewURLCleanerHandler(deps *core.Dependency, logic URLCleanerLogicInterface) *URLCLeanerHandler {
	return &URLCLeanerHandler{
		deps:  deps,
		logic: logic,
	}
}

// CleanURL godoc
// @Summary Clean up URL with either canonical, redirection, or both (all) operations
// @Tags urlCleaner
// @Produce json
// @Param data body model.URLCleanerRequest true "URL to be clean and the selected operation ('canonical', 'redirection', 'all')"
// @Success 200 {object} model.URLCleanerResponse{}
// @Router /url/cleanup [post]
func (h *URLCLeanerHandler) CleanURL(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// parse request body
	var payload model.URLCleanerRequest
	err := xhttp.BindJSONRequest(r, &payload)
	if err != nil {
		xhttp.SendJSONResponse(w, xhttp.BaseResponse{
			Error:   err.Error(),
			Message: "failed to parse request body",
		}, http.StatusBadRequest)
		return
	}

	processedURL, err := h.logic.CleanURL(ctx, payload.URL, payload.Operation)
	if err != nil {
		h.deps.Logger.ErrorContext(ctx, "failed to clean url up", slog.Any("error", err))
		xhttp.SendJSONResponse(w, xhttp.BaseResponse{
			Error:   err.Error(),
			Message: "failed to clean url up",
		}, xerrors.ParseErrorTypeToCodeInt(err))
		return
	}

	xhttp.SendJSONResponse(w, model.URLCleanerResponse{
		ProcessedURL: processedURL,
	}, http.StatusOK)
}

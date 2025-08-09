package book

import (
	"byfood-app/internal/core"
	"byfood-app/internal/model"
	"byfood-app/internal/pkg/pagination"
	"byfood-app/internal/pkg/xhttp"
	"log/slog"
	"net/http"
)

type BookHandler struct {
	deps  *core.Dependency
	logic LogicInterface
}

func NewHTTPHandler(deps *core.Dependency, logic LogicInterface) *BookHandler {
	return &BookHandler{
		deps:  deps,
		logic: logic,
	}
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	page, err := pagination.ParsePaginationRequest(r)
	if err != nil {
		h.deps.Logger.WarnContext(ctx, "failed to parse pagination params", slog.Any("error", err))
		xhttp.SendJSONResponse(w, xhttp.BaseResponse{
			Error:   err.Error(),
			Message: "failed to parse pagination params",
		}, http.StatusInternalServerError)
		return
	}

	data, meta, err := h.logic.GetBooks(ctx, model.BookSearchParams{}, page)
	if err != nil {
		h.deps.Logger.ErrorContext(ctx, "failed to get book(s)", slog.Any("error", err))
		xhttp.SendJSONResponse(w, xhttp.BaseResponse{
			Error:   err.Error(),
			Message: "failed to get book(s)",
		}, http.StatusInternalServerError)
		return
	}

	xhttp.SendJSONResponse(w, xhttp.BaseListResponse{
		Message:  "books fetched",
		Data:     data,
		Metadata: meta,
	}, http.StatusOK)
}

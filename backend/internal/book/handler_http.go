package book

import (
	"byfood-app/internal/core"
	"byfood-app/internal/model"
	"byfood-app/internal/pkg/pagination"
	"byfood-app/internal/pkg/xerrors"
	"byfood-app/internal/pkg/xhttp"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
		}, http.StatusBadRequest)
		return
	}

	data, meta, err := h.logic.GetBooks(ctx, model.BookSearchParams{
		Search: r.URL.Query().Get("search"),
	}, page)
	if err != nil && !errors.Is(err, xerrors.ErrDataNotFound) {
		h.deps.Logger.ErrorContext(ctx, "failed to get book(s)", slog.Any("error", err))
		xhttp.SendJSONResponse(w, xhttp.BaseResponse{
			Error:   err.Error(),
			Message: "failed to get book(s)",
		}, xerrors.ParseErrorTypeToCodeInt(err))
		return
	}

	xhttp.SendJSONResponse(w, xhttp.BaseListResponse{
		Message:  "books fetched",
		Data:     data,
		Metadata: meta,
	}, http.StatusOK)
}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// get and validate id param
	id := chi.URLParam(r, "id")
	if id == "" {
		xhttp.SendJSONResponse(w, xhttp.BaseResponse{
			Error:   xerrors.ErrInvalidID.Error(),
			Message: xerrors.ErrInvalidID.Error(),
		}, http.StatusBadRequest)
		return
	}
	idParam, err := strconv.Atoi(id)
	if err != nil {
		xhttp.SendJSONResponse(w, xhttp.BaseResponse{
			Error:   err.Error(),
			Message: "failed to parse id parameter",
		}, http.StatusBadRequest)
		return
	}

	data, err := h.logic.GetBookByID(ctx, int64(idParam))
	if err != nil {
		h.deps.Logger.ErrorContext(ctx, "failed to get book data by id", slog.Any("error", err))
		xhttp.SendJSONResponse(w, xhttp.BaseResponse{
			Error:   err.Error(),
			Message: "failed to get book data",
		}, xerrors.ParseErrorTypeToCodeInt(err))
		return
	}

	xhttp.SendJSONResponse(w, xhttp.BaseResponse{
		Data:    data,
		Message: "book data fetched",
	}, http.StatusOK)
}

func (h *BookHandler) StoreBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// parse request body
	var payload model.StoreBookRequest
	err := xhttp.BindJSONRequest(r, &payload)
	if err != nil {
		xhttp.SendJSONResponse(w, xhttp.BaseResponse{
			Error:   err.Error(),
			Message: "failed to parse request body",
		}, http.StatusBadRequest)
		return
	}

	data, err := h.logic.StoreBook(ctx, model.Book{
		Title:       payload.Title,
		Author:      payload.Author,
		PublishYear: payload.PublishYear,
	})
	if err != nil {
		h.deps.Logger.ErrorContext(ctx, "failed to store book data", slog.Any("error", err))
		xhttp.SendJSONResponse(w, xhttp.BaseResponse{
			Error:   err.Error(),
			Message: "failed to store book data",
		}, xerrors.ParseErrorTypeToCodeInt(err))
		return
	}

	xhttp.SendJSONResponse(w, xhttp.BaseResponse{
		Data:    data,
		Message: "book data stored",
	}, http.StatusOK)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// get and validate id param
	id := chi.URLParam(r, "id")
	if id == "" {
		xhttp.SendJSONResponse(w, xhttp.BaseResponse{
			Error:   xerrors.ErrInvalidID.Error(),
			Message: xerrors.ErrInvalidID.Error(),
		}, http.StatusBadRequest)
		return
	}
	idParam, err := strconv.Atoi(id)
	if err != nil {
		xhttp.SendJSONResponse(w, xhttp.BaseResponse{
			Error:   err.Error(),
			Message: "failed to parse id parameter",
		}, http.StatusBadRequest)
		return
	}

	// parse request body
	var payload model.StoreBookRequest
	err = xhttp.BindJSONRequest(r, &payload)
	if err != nil {
		xhttp.SendJSONResponse(w, xhttp.BaseResponse{
			Error:   err.Error(),
			Message: "failed to parse request body",
		}, http.StatusBadRequest)
		return
	}

	data, err := h.logic.UpdateBook(ctx, model.Book{
		ID:          int64(idParam),
		Title:       payload.Title,
		Author:      payload.Author,
		PublishYear: payload.PublishYear,
	})
	if err != nil {
		h.deps.Logger.ErrorContext(ctx, "failed to update book data", slog.Any("error", err))
		xhttp.SendJSONResponse(w, xhttp.BaseResponse{
			Error:   err.Error(),
			Message: "failed to update book data",
		}, xerrors.ParseErrorTypeToCodeInt(err))
		return
	}

	xhttp.SendJSONResponse(w, xhttp.BaseResponse{
		Data:    data,
		Message: "book data updated",
	}, http.StatusOK)

}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// get and validate id param
	id := chi.URLParam(r, "id")
	if id == "" {
		xhttp.SendJSONResponse(w, xhttp.BaseResponse{
			Error:   xerrors.ErrInvalidID.Error(),
			Message: xerrors.ErrInvalidID.Error(),
		}, http.StatusBadRequest)
		return
	}
	idParam, err := strconv.Atoi(id)
	if err != nil {
		xhttp.SendJSONResponse(w, xhttp.BaseResponse{
			Error:   err.Error(),
			Message: "failed to parse id parameter",
		}, http.StatusBadRequest)
		return
	}

	err = h.logic.DeleteBook(ctx, int64(idParam))
	if err != nil {
		h.deps.Logger.ErrorContext(ctx, "failed to delete book data", slog.Any("error", err))
		xhttp.SendJSONResponse(w, xhttp.BaseResponse{
			Error:   err.Error(),
			Message: "failed to delete book data",
		}, xerrors.ParseErrorTypeToCodeInt(err))
		return
	}

	xhttp.SendJSONResponse(w, xhttp.BaseResponse{
		Message: "book data deleted",
	}, http.StatusOK)
}

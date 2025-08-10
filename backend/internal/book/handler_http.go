package book

import (
	"byfood-app/internal/core"
	"byfood-app/internal/model"
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

// GetBooks godoc
// @Summary List all books with pagination and search query params
// @Tags books
// @Produce json
// @Param search query string false "search param to search by title and author"
// @Param page query integer false "page number"
// @Param size query integer false "item per page"
// @Success 200 {object} xhttp.BaseResponse{data=[]model.Book, metadata=pagination.Metadata}
// @Router /books [get]
func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// these functions works, but
	// use GetBooksNoPagination for now
	// cuz it's hard to implement
	// pagination on the frontend side
	// skill issue on me :(

	// page, err := pagination.ParsePaginationRequest(r)
	// if err != nil {
	// 	h.deps.Logger.WarnContext(ctx, "failed to parse pagination params", slog.Any("error", err))
	// 	xhttp.SendJSONResponse(w, xhttp.BaseResponse{
	// 		Error:   err.Error(),
	// 		Message: "failed to parse pagination params",
	// 	}, http.StatusBadRequest)
	// 	return
	// }

	// data, meta, err := h.logic.GetBooks(ctx, model.BookSearchParams{
	// 	Search: r.URL.Query().Get("search"),
	// }, page)
	// if err != nil && !errors.Is(err, xerrors.ErrDataNotFound) {
	// 	h.deps.Logger.ErrorContext(ctx, "failed to get book(s)", slog.Any("error", err))
	// 	xhttp.SendJSONResponse(w, xhttp.BaseResponse{
	// 		Error:   err.Error(),
	// 		Message: "failed to get book(s)",
	// 	}, xerrors.ParseErrorTypeToCodeInt(err))
	// 	return
	// }

	// xhttp.SendJSONResponse(w, xhttp.BaseListResponse{
	// 	Message:  "books fetched",
	// 	Data:     data,
	// 	Metadata: meta,
	// }, http.StatusOK)

	data, err := h.logic.GetBooksNoPagination(ctx, model.BookSearchParams{
		Search: r.URL.Query().Get("search"),
	})
	if err != nil && !errors.Is(err, xerrors.ErrDataNotFound) {
		h.deps.Logger.ErrorContext(ctx, "failed to get book(s)", slog.Any("error", err))
		xhttp.SendJSONResponse(w, xhttp.BaseResponse{
			Error:   err.Error(),
			Message: "failed to get book(s)",
		}, xerrors.ParseErrorTypeToCodeInt(err))
		return
	}
	xhttp.SendJSONResponse(w, xhttp.BaseListResponse{
		Message: "books fetched",
		Data:    data,
	}, http.StatusOK)
}

// GetBook godoc
// @Summary Get a book data by its ID
// @Tags books
// @Produce json
// @Param id path integer true "book ID"
// @Success 200 {object} xhttp.BaseResponse{data=model.Book}
// @Router /books/{id} [get]
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

// StoreBook godoc
// @Summary Store new book data, return stored data
// @Tags books
// @Produce json
// @Param data body model.StoreBookRequest true "book data"
// @Success 200 {object} xhttp.BaseResponse{data=model.Book}
// @Router /books [post]
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

// UpdateBook godoc
// @Summary Update book data by ID, return updated data
// @Tags books
// @Produce json
// @Param id path integer true "book ID"
// @Param data body model.UpdateBookRequest true "book data"
// @Success 200 {object} xhttp.BaseResponse{data=model.Book}
// @Router /books/{id} [put]
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
	var payload model.UpdateBookRequest
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

// DeleteBook godoc
// @Summary Delete book data by ID
// @Tags books
// @Produce json
// @Param id path integer true "book ID"
// @Success 200 {object} xhttp.BaseResponse{message=string}
// @Router /books/{id} [delete]
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

package book

import (
	"byfood-app/internal/core"
	"byfood-app/internal/model"
	"byfood-app/internal/pkg/pagination"
	"context"
	"log/slog"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

type testSuite struct {
	Ctrl         *gomock.Controller
	MockBookRepo *MockRepositoryInterface
}

func setupTestSuite(t *testing.T) *testSuite {
	ctrl := gomock.NewController(t)
	return &testSuite{
		Ctrl:         ctrl,
		MockBookRepo: NewMockRepositoryInterface(ctrl),
	}
}

func TestBookLogic_StoreBook(t *testing.T) {
	type fields struct {
		deps *core.Dependency
		repo RepositoryInterface
	}
	type args struct {
		ctx  context.Context
		data model.Book
	}

	ts := setupTestSuite(t)
	mockFields := fields{
		deps: &core.Dependency{
			Logger: slog.Default(),
		},
		repo: ts.MockBookRepo,
	}

	expectedResult := model.Book{}

	tests := []struct {
		name       string
		fields     fields
		args       args
		want       model.Book
		wantErr    bool
		expectFunc func()
	}{
		{
			name:   "success store book data",
			fields: mockFields,
			args: args{
				ctx: context.Background(),
				data: model.Book{
					Title:       "One Piece",
					Author:      "Eiichiro Oda",
					PublishYear: 1997,
				},
			},
			want:    expectedResult,
			wantErr: false,
			expectFunc: func() {
				ts.MockBookRepo.EXPECT().StoreBook(gomock.Any(), model.Book{
					Title:       "One Piece",
					Author:      "Eiichiro Oda",
					PublishYear: 1997,
				}).Return(
					expectedResult,
					nil,
				)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logic := &BookLogic{
				deps: tt.fields.deps,
				repo: tt.fields.repo,
			}

			tt.expectFunc()

			got, err := logic.StoreBook(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("BookLogic.StoreBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BookLogic.StoreBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBookLogic_GetBooks(t *testing.T) {
	type fields struct {
		deps *core.Dependency
		repo RepositoryInterface
	}
	type args struct {
		ctx    context.Context
		params model.BookSearchParams
		page   pagination.Page
	}

	ts := setupTestSuite(t)
	mockFields := fields{
		deps: &core.Dependency{
			Logger: slog.Default(),
		},
		repo: ts.MockBookRepo,
	}

	expectedResult := []model.Book{}

	tests := []struct {
		name       string
		fields     fields
		args       args
		want       []model.Book
		wantMeta   pagination.Metadata
		wantErr    bool
		expectFunc func()
	}{
		{
			name:   "success get books with search param",
			fields: mockFields,
			args: args{
				ctx: context.Background(),
				params: model.BookSearchParams{
					Search: "oda",
				},
				page: pagination.Page{},
			},
			want:     expectedResult,
			wantMeta: pagination.Metadata{},
			wantErr:  false,
			expectFunc: func() {
				ts.MockBookRepo.EXPECT().GetBooks(gomock.Any(), model.BookSearchParams{
					Search: "oda",
				}, pagination.Page{}).Return(
					expectedResult,
					pagination.Metadata{},
					nil,
				)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logic := &BookLogic{
				deps: tt.fields.deps,
				repo: tt.fields.repo,
			}

			tt.expectFunc()

			got, meta, err := logic.GetBooks(tt.args.ctx, tt.args.params, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("BookLogic.GetBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BookLogic.GetBooks() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(meta, tt.wantMeta) {
				t.Errorf("BookLogic.GetBooks() meta = %v, want %v", meta, tt.wantMeta)
			}
		})
	}
}

func TestBookLogic_GetBookByID(t *testing.T) {
	type fields struct {
		deps *core.Dependency
		repo RepositoryInterface
	}
	type args struct {
		ctx context.Context
		id  int64
	}

	ts := setupTestSuite(t)
	mockFields := fields{
		deps: &core.Dependency{
			Logger: slog.Default(),
		},
		repo: ts.MockBookRepo,
	}

	expectedResult := model.Book{}

	tests := []struct {
		name       string
		fields     fields
		args       args
		want       model.Book
		wantErr    bool
		expectFunc func()
	}{
		{
			name:   "success get book data by id",
			fields: mockFields,
			args: args{
				ctx: context.Background(),
				id:  int64(1),
			},
			want:    expectedResult,
			wantErr: false,
			expectFunc: func() {
				ts.MockBookRepo.EXPECT().GetBookByID(gomock.Any(), int64(1)).Return(expectedResult, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logic := &BookLogic{
				deps: tt.fields.deps,
				repo: tt.fields.repo,
			}

			tt.expectFunc()

			got, err := logic.GetBookByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("BookLogic.GetBookByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BookLogic.GetBookByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBookLogic_UpdateBook(t *testing.T) {
	type fields struct {
		deps *core.Dependency
		repo RepositoryInterface
	}
	type args struct {
		ctx  context.Context
		data model.Book
	}

	ts := setupTestSuite(t)
	mockFields := fields{
		deps: &core.Dependency{
			Logger: slog.Default(),
		},
		repo: ts.MockBookRepo,
	}

	expectedResult := model.Book{}

	tests := []struct {
		name       string
		fields     fields
		args       args
		want       model.Book
		wantErr    bool
		expectFunc func()
	}{
		{
			name:   "success update book data",
			fields: mockFields,
			args: args{
				ctx: context.Background(),
				data: model.Book{
					ID:          int64(1),
					Title:       "One Piece",
					Author:      "Eiichiro Oda",
					PublishYear: 1997,
				},
			},
			want:    expectedResult,
			wantErr: false,
			expectFunc: func() {
				ts.MockBookRepo.EXPECT().UpdateBook(gomock.Any(), model.Book{
					ID:          int64(1),
					Title:       "One Piece",
					Author:      "Eiichiro Oda",
					PublishYear: 1997,
				}).Return(
					expectedResult,
					nil,
				)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logic := &BookLogic{
				deps: tt.fields.deps,
				repo: tt.fields.repo,
			}

			tt.expectFunc()

			got, err := logic.UpdateBook(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("BookLogic.UpdateBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BookLogic.UpdateBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBookLogic_DeleteBook(t *testing.T) {
	type fields struct {
		deps *core.Dependency
		repo RepositoryInterface
	}
	type args struct {
		ctx context.Context
		id  int64
	}

	ts := setupTestSuite(t)
	mockFields := fields{
		deps: &core.Dependency{
			Logger: slog.Default(),
		},
		repo: ts.MockBookRepo,
	}

	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    bool
		expectFunc func()
	}{
		{
			name:   "success delete book data",
			fields: mockFields,
			args: args{
				ctx: context.Background(),
				id:  int64(1),
			},
			wantErr: false,
			expectFunc: func() {
				ts.MockBookRepo.EXPECT().DeleteBook(gomock.Any(), int64(1)).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logic := &BookLogic{
				deps: tt.fields.deps,
				repo: tt.fields.repo,
			}

			tt.expectFunc()

			if err := logic.DeleteBook(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("BookLogic.DeleteBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package book

import (
	"byfood-app/internal/core"
	"byfood-app/internal/model"
	"byfood-app/internal/pkg/pagination"
	"context"
	"reflect"
	"testing"
)

func TestBookRepo_GetBooks(t *testing.T) {
	type fields struct {
		deps *core.Dependency
	}
	type args struct {
		ctx    context.Context
		params model.BookSearchParams
		page   pagination.Page
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Book
		want1   pagination.Metadata
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &BookRepo{
				deps: tt.fields.deps,
			}
			got, got1, err := repo.GetBooks(tt.args.ctx, tt.args.params, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("BookRepo.GetBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BookRepo.GetBooks() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BookRepo.GetBooks() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

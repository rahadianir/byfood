package urlcleaner

import (
	"byfood-app/internal/core"
	"context"
	"log/slog"
	"testing"
)

func TestURLCleanerLogic_CleanURL(t *testing.T) {
	type fields struct {
		deps *core.Dependency
	}
	type args struct {
		ctx       context.Context
		link      string
		operation string
	}

	mockFields := fields{
		deps: &core.Dependency{
			Logger: slog.Default(),
		},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "success cleanup url - operation: all",
			fields: mockFields,
			args: args{
				ctx:       context.Background(),
				link:      "https://BYFOOD.com/food-EXPeriences?query=abc/",
				operation: "all",
			},
			want:    "https://www.byfood.com/food-experiences",
			wantErr: false,
		},
		{
			name:   "success cleanup url - operation: canonical",
			fields: mockFields,
			args: args{
				ctx:       context.Background(),
				link:      "https://BYFOOD.com/food-EXPeriences?query=abc/",
				operation: "canonical",
			},
			want:    "https://BYFOOD.com/food-EXPeriences",
			wantErr: false,
		},
		{
			name:   "success cleanup url - operation: redirection",
			fields: mockFields,
			args: args{
				ctx:       context.Background(),
				link:      "https://BYFOOD.com/food-EXPeriences?query=abc/",
				operation: "redirection",
			},
			want:    "https://www.byfood.com/food-experiences?query=abc/",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logic := &URLCleanerLogic{
				deps: tt.fields.deps,
			}
			got, err := logic.CleanURL(tt.args.ctx, tt.args.link, tt.args.operation)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLCleanerLogic.CleanURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("URLCleanerLogic.CleanURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

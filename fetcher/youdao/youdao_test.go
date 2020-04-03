package youdao

import (
	"reflect"
	"testing"

	"github.com/qshuai/keyhot/fetcher"
)

func TestYouDaoFetcher_Translate(t *testing.T) {
	type fields struct {
		appId  string
		appKey string
	}
	type args struct {
		word string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *fetcher.Result
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				appId:  "2058e60531620627",
				appKey: "4qmikGdU4U6SzafZyVBae37Ru1IwkKkV",
			},
			args: args{
				word: "hello",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &YouDaoFetcher{
				appId:  tt.fields.appId,
				appKey: tt.fields.appKey,
			}

			got, err := y.Translate(tt.args.word)
			if (err != nil) != tt.wantErr {
				t.Errorf("Translate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Translate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

package feedly

import (
	"reflect"
	"testing"
)

func TestClient_FetchFeed(t *testing.T) {

	c := NewClient()
	type args struct {
		url string
	}
	tests := []struct {
		name     string
		args     args
		wantCode uint32
		wantRes  []byte
		wantErr  RequestError
	}{
		{
			name: "test1",
			args: args{
				url: "https://www.caranddriver.com/rss/all.xml/",
			},
			wantCode: 200,
			wantRes:  []byte{},
			wantErr:  RequestError{nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCode, gotRes, gotErr := c.FetchFeed(tt.args.url)
			if len(gotRes) == 0 {
				t.Errorf("FetchFeed() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("FetchFeed() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
			if gotCode != tt.wantCode {
				t.Errorf("FetchFeed() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}

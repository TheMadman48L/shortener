package server

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/TheMadman48L/shortener/internal/server/mocks"
)

func Test_server_handleShortingURL(t *testing.T) {
	type fields struct {
		r       *mux.Router
		shorten *mocks.Shortener
	}
	type args struct {
		value string
	}
	type want struct {
		code        int
		value       string
		callService bool
		hash        string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "OK",
			fields: fields{
				r:       mux.NewRouter(),
				shorten: &mocks.Shortener{},
			},
			args: args{
				value: "ya.ru",
			},
			want: want{
				code:        http.StatusCreated,
				value:       "http://localhost:8080/qwertyu",
				hash:        "qwertyu",
				callService: true,
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want.callService {
				tt.fields.shorten.On("ShortingURL", tt.args.value).Return(tt.want.hash)
			}

			serv := &server{
				r:       tt.fields.r,
				shorten: tt.fields.shorten,
			}
			bodyReader := strings.NewReader(tt.args.value)
			request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", bodyReader)
			w := httptest.NewRecorder()

			h := http.HandlerFunc(serv.handleShortingURL)
			h.ServeHTTP(w, request)
			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.want.code {
				t.Errorf("Status code missmatched got %v, want %v", res.StatusCode, tt.want.code)
			}

			{
				result, err := ioutil.ReadAll(res.Body)
				assert.NoError(t, err)
				assert.Equal(t, tt.want.value, string(result))
			}

			if tt.want.callService {
				tt.fields.shorten.AssertCalled(t, "ShortingURL", tt.args.value)
				tt.fields.shorten.AssertNumberOfCalls(t, "ShortingURL", 1)
			}
		})
	}
}

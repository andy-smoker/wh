package handler

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/andy-smoker/wh-server/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock_service "github.com/andy-smoker/wh-server/pkg/service/mocks"
)

func TestHandler_userIdentity(t *testing.T) {
	type mockBehavior func(r *mock_service.MockAuthorization, token string)

	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehavior         mockBehavior
		exeptionStatusCode   int
		exeptionResponseBody string
	}{
		{
			name:        "OK",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(r *mock_service.MockAuthorization, token string) {
				r.EXPECT().ParseToken(token).Return(1, nil)
			},
			exeptionStatusCode:   200,
			exeptionResponseBody: "1",
		},
		{
			name:                 "Invalid Header Name",
			headerName:           "",
			headerValue:          "Bearer token",
			token:                "token",
			mockBehavior:         func(r *mock_service.MockAuthorization, token string) {},
			exeptionStatusCode:   401,
			exeptionResponseBody: `{"message":"empty auth header"}`,
		},
		{
			name:                 "Invalid Header Value",
			headerName:           "Authorization",
			headerValue:          "Bear token",
			token:                "token",
			mockBehavior:         func(r *mock_service.MockAuthorization, token string) {},
			exeptionStatusCode:   401,
			exeptionResponseBody: `{"message":"invalid auth header"}`,
		},
		{
			name:                 "Empty Token",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			token:                "token",
			mockBehavior:         func(r *mock_service.MockAuthorization, token string) {},
			exeptionStatusCode:   401,
			exeptionResponseBody: `{"message":"token is empty"}`,
		},
		{
			name:        "Error parse Token",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(r *mock_service.MockAuthorization, token string) {
				r.EXPECT().ParseToken(token).Return(0, errors.New("invalid token"))
			},
			exeptionStatusCode:   401,
			exeptionResponseBody: `{"message":"invalid token"}`,
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorization(c)
			test.mockBehavior(repo, test.token)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			// init Endpoint
			r := gin.New()
			r.GET("/identity", handler.userIdetinty, func(c *gin.Context) {
				id, _ := c.Get(userCtx)
				c.String(200, "%d", id)
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/identity", nil)
			req.Header.Set(test.headerName, test.headerValue)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.exeptionStatusCode)
			assert.Equal(t, w.Body.String(), test.exeptionResponseBody)
		})
	}
}

func TestGetUserID(t *testing.T) {
	var getContext = func(id int) *gin.Context {
		ctx := &gin.Context{}
		ctx.Set(userCtx, id)
		return ctx
	}
	testTable := []struct {
		name       string
		ctx        *gin.Context
		id         int
		shouldFail bool
	}{
		{
			name: "OK",
			ctx:  getContext(1),
			id:   1,
		},
		{
			name:       "Empty",
			ctx:        &gin.Context{},
			shouldFail: true,
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			id, err := getUserID(test.ctx)
			if test.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, id, test.id)
		})
	}
}

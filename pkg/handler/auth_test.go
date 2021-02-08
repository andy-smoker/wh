package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andy-smoker/wh-server/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	mock_service "github.com/andy-smoker/wh-server/pkg/service/mocks"
	"github.com/andy-smoker/wh-server/pkg/structs"
)

func TestHandlerSignUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user structs.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           structs.User
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"login":"test", "username":"test", "pass":"qwerty"}`,
			inputUser: structs.User{
				Login:    "test",
				Username: "test",
				Pass:     "qwerty",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user structs.User) {
				r.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:                "WrongInput",
			inputBody:           `{"login":"test"}`,
			inputUser:           structs.User{},
			mockBehavior:        func(r *mock_service.MockAuthorization, user structs.User) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid body"}`,
		},
		{
			name:      "OK",
			inputBody: `{"login":"test", "username":"test", "pass":"qwerty"}`,
			inputUser: structs.User{
				Login:    "test",
				Username: "test",
				Pass:     "qwerty",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user structs.User) {
				r.EXPECT().CreateUser(user).Return(0, errors.New("something went wrong"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"something went wrong"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			service := &service.Service{Authorization: auth}
			handler := Handler{service}

			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}
}

package user

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"goplay-backend-engineer-test/helper"
	"goplay-backend-engineer-test/usecase/user/userlogin"
	"goplay-backend-engineer-test/usecase/user/userlogin/mock"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type LoginUserTestCase struct {
	Name       string
	Payload    interface{}
	MockInport func(m *mock.MockInport)
	StatusCode int
	ErrMsg     string
}

func TestLoginUser(t *testing.T) {
	testCases := []LoginUserTestCase{
		{
			Name: "should be succeed",
			Payload: userlogin.InportRequest{
				Username: "username",
				Password: "password",
			},
			MockInport: func(m *mock.MockInport) {
				m.EXPECT().Execute(
					gomock.Any(),
					gomock.Any(),
				).Return(userlogin.InportResponse{}, nil)
			},
			StatusCode: http.StatusOK,
		},
		{
			Name:       "failed unmarshal request payload",
			Payload:    "{}",
			MockInport: func(m *mock.MockInport) {},
			StatusCode: http.StatusBadRequest,
			ErrMsg:     "Something wrong with input",
		},
		{
			Name: "unable validate request payload",
			Payload: userlogin.InportRequest{
				Username: "username",
			},
			MockInport: func(m *mock.MockInport) {},
			StatusCode: http.StatusBadRequest,
			ErrMsg:     "[\"validation failed on field password with precondition 'required'\"]",
		},
		{
			Name: "should be failed",
			Payload: userlogin.InportRequest{
				Username: "username",
				Password: "password",
			},
			MockInport: func(m *mock.MockInport) {
				m.EXPECT().Execute(
					gomock.Any(),
					gomock.Any(),
				).Return(userlogin.InportResponse{}, helper.ErrFatalQuery)
			},
			StatusCode: http.StatusInternalServerError,
			ErrMsg:     helper.ErrFatalQuery.Error(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			path := "/user/login"

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			inport := mock.NewMockInport(ctrl)
			tc.MockInport(inport)

			router := gin.Default()
			router.POST(path, LoginUser(inport))

			var payload io.Reader
			if tc.Payload != nil {
				body, err := json.Marshal(tc.Payload)
				assert.NoError(t, err)
				payload = bytes.NewReader(body)
			}

			resp := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, path, payload)
			assert.NoError(t, err)

			router.ServeHTTP(resp, req)

			var result gin.H
			err = json.Unmarshal(resp.Body.Bytes(), &result)
			assert.NoError(t, err)

			if tc.ErrMsg != "" {
				assert.Equal(t, tc.StatusCode, resp.Code)
				assert.Equal(t, tc.ErrMsg, result["message"])
				assert.Equal(t, false, result["success"])
				return
			}

			assert.Equal(t, tc.StatusCode, resp.Code)
			assert.Equal(t, true, result["success"])
		})
	}
}

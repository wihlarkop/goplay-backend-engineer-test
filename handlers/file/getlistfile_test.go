package file

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"goplay-backend-engineer-test/entities"
	"goplay-backend-engineer-test/helper"
	"goplay-backend-engineer-test/usecase/file/getlistfile"
	"goplay-backend-engineer-test/usecase/file/getlistfile/mock"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type GetListFileTestCase struct {
	Name       string
	Payload    interface{}
	UserClaim  gin.HandlerFunc
	StatusCode int
	ErrMsg     string
	MockInport func(m *mock.MockInport)
}

func TestGetListFile(t *testing.T) {
	testCases := []GetListFileTestCase{
		{
			Name: "should be succeed",
			UserClaim: func(c *gin.Context) {
				claims := &helper.TokenRequest{
					Name:      "name",
					CreatedAt: time.Now().String(),
				}
				c.Set("user", claims)
			},
			Payload: getlistfile.InportRequest{
				UploadFileFilter: entities.UploadFileFilter{
					Page:  1,
					Limit: 10,
				},
			},
			MockInport: func(m *mock.MockInport) {
				m.EXPECT().Execute(
					gomock.Any(),
					gomock.Any(),
				).Return(getlistfile.InportResponse{}, nil)
			},
			StatusCode: http.StatusOK,
		},
		{
			Name: "unable validate",
			UserClaim: func(c *gin.Context) {
				claims := &helper.TokenRequest{
					Name:      "name",
					CreatedAt: time.Now().String(),
				}
				c.Set("user", claims)
			},
			Payload: getlistfile.InportRequest{
				UploadFileFilter: entities.UploadFileFilter{
					Page: -1,
				},
			},
			MockInport: func(m *mock.MockInport) {},
			StatusCode: http.StatusBadRequest,
			ErrMsg:     "[\"validation failed on field limit with precondition 'required' but got 0\"]",
		},
		{
			Name: "failed binding",
			UserClaim: func(c *gin.Context) {
				claims := &helper.TokenRequest{
					Name:      "name",
					CreatedAt: time.Now().String(),
				}
				c.Set("user", claims)
			},
			Payload: map[string]interface{}{
				"page":  "a",
				"limit": "b",
			},
			MockInport: func(m *mock.MockInport) {},
			StatusCode: http.StatusBadRequest,
			ErrMsg:     "fatal error: strconv.ParseInt: parsing \"a\": invalid syntax",
		},
		{
			Name: "should be failed",
			UserClaim: func(c *gin.Context) {
				claims := &helper.TokenRequest{
					Name:      "name",
					CreatedAt: time.Now().String(),
				}
				c.Set("user", claims)
			},
			Payload: getlistfile.InportRequest{
				UploadFileFilter: entities.UploadFileFilter{
					Page:  1,
					Limit: 10,
				},
			},
			MockInport: func(m *mock.MockInport) {
				m.EXPECT().Execute(
					gomock.Any(),
					gomock.Any(),
				).Return(getlistfile.InportResponse{}, helper.ErrFatalQuery)
			},
			StatusCode: http.StatusInternalServerError,
			ErrMsg:     helper.ErrFatalQuery.Error(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			path := "/file"

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			inport := mock.NewMockInport(ctrl)
			tc.MockInport(inport)

			router := gin.Default()
			router.GET(path, GetListFile(inport))

			resp := httptest.NewRecorder()

			qParam := helper.TransformStructToUrlValue(tc.Payload)

			urlHandler, _ := url.Parse(path)
			urlHandler.RawQuery = qParam.Encode()

			req, err := http.NewRequest(http.MethodGet, urlHandler.String(), nil)
			assert.NoError(t, err)

			router.ServeHTTP(resp, req)

			var result gin.H
			err = json.Unmarshal(resp.Body.Bytes(), &result)
			assert.NoError(t, err)

			if tc.ErrMsg != "" {
				assert.Equal(t, tc.StatusCode, resp.Code)
				assert.Equal(t, tc.ErrMsg, result["message"])
				return
			}

			assert.Equal(t, tc.StatusCode, resp.Code)
			assert.Equal(t, true, result["success"])
		})

	}
}

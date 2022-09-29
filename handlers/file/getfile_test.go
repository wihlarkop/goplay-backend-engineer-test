package file

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"goplay-backend-engineer-test/helper"
	"goplay-backend-engineer-test/usecase/file/getfile"
	"goplay-backend-engineer-test/usecase/file/getfile/mock"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type PayloadStruct struct {
	Id string `json:"id" form:"id"`
}
type GetFileTestCase struct {
	Name       string
	Payload    PayloadStruct
	StatusCode int
	ErrMsg     string
	MockInport func(m *mock.MockInport)
}

func TestGetFile(t *testing.T) {

	testCases := []GetFileTestCase{
		{
			Name:    "should be succeed",
			Payload: PayloadStruct{Id: "1"},
			MockInport: func(m *mock.MockInport) {
				m.EXPECT().Execute(
					gomock.Any(),
					gomock.Any(),
				).Return(getfile.InportResponse{}, nil)
			},
			StatusCode: http.StatusOK,
		},
		{
			Name:       "should be bad request",
			Payload:    PayloadStruct{Id: "a"},
			MockInport: func(m *mock.MockInport) {},
			StatusCode: http.StatusBadRequest,
			ErrMsg:     "strconv.Atoi: parsing \"a\": invalid syntax",
		},
		{
			Name:    "should be failed",
			Payload: PayloadStruct{Id: "1"},
			MockInport: func(m *mock.MockInport) {
				m.EXPECT().Execute(
					gomock.Any(),
					gomock.Any(),
				).Return(getfile.InportResponse{}, helper.ErrFatalQuery)
			},
			StatusCode: http.StatusInternalServerError,
			ErrMsg:     "fatal query error",
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
			router.GET(fmt.Sprintf("%s/:id", path), GetFile(inport))

			resp := httptest.NewRecorder()

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", path, tc.Payload.Id), nil)
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

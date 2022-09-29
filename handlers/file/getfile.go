package file

import (
	"strconv"

	"goplay-backend-engineer-test/helper"
	"goplay-backend-engineer-test/usecase/file/getfile"

	"github.com/gin-gonic/gin"
)

func GetFile(inport getfile.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req getfile.InportRequest
		var id, _ = strconv.Atoi(c.Param("id"))

		// ? Binding Request
		if err := helper.UnmarshalJSON(c, &req); err != nil {
			return
		}

		req.Id = id

		// ? Validate Request
		if err := helper.Validate(&req); err != nil {
			helper.WriteError(c, err)
			return
		}

		resp, err := inport.Execute(c.Copy().Request.Context(), req)
		if err != nil {
			helper.WriteError(c, err)
			return
		}

		helper.WriteSuccess(c, "Success get file", resp, nil)
	}
}

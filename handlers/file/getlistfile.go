package file

import (
	"goplay-backend-engineer-test/helper"
	"goplay-backend-engineer-test/usecase/file/getlistfile"

	"github.com/gin-gonic/gin"
)

func GetListFile(inport getlistfile.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req getlistfile.InportRequest

		// ? Binding Request
		if err := c.Bind(&req); err != nil {
			helper.WriteError(c, err)
			return
		}

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

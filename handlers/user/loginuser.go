package user

import (
	"goplay-backend-engineer-test/helper"
	"goplay-backend-engineer-test/usecase/user/userlogin"

	"github.com/gin-gonic/gin"
)

func LoginUser(inport userlogin.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req userlogin.InportRequest

		// ? Binding Request
		if err := helper.UnmarshalJSON(c, &req); err != nil {
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

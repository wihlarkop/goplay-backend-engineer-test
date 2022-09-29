package file

import (
	"net/http"
	"strconv"

	"goplay-backend-engineer-test/helper"
	"goplay-backend-engineer-test/usecase/file/getfile"

	"github.com/gin-gonic/gin"
)

func GetFile(inport getfile.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req getfile.InportRequest

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
				"data":    nil,
			})
			return
		}

		req.Id = id

		resp, err := inport.Execute(c.Copy().Request.Context(), req)
		if err != nil {
			helper.WriteError(c, err)
			return
		}

		helper.WriteSuccess(c, "Success get file", resp, nil)
	}
}

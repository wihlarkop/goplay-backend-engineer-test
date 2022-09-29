package file

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"goplay-backend-engineer-test/helper"
	"goplay-backend-engineer-test/usecase/file/uploadfile"

	"github.com/gin-gonic/gin"
)

func UploadFile(inport uploadfile.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req uploadfile.InportRequest

		// ? Binding Request
		if err := c.ShouldBind(&req); err != nil {
			helper.WriteError(c, err)
			return
		}

		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		allowedExtension := []string{"image/jpg", "image/png", "video/webm", "video/ogg"}
		if !helper.Contains(allowedExtension, file.Header.Get("Content-Type")) {
			c.JSON(http.StatusBadRequest, "Not allowed extension")
		}

		fileExt := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("%d%s", time.Now().UnixMilli(), fileExt)
		path := fmt.Sprintf("media/%s", req.Location)
		dst := fmt.Sprintf("%s/%s", path, filename)

		if _, err := os.Stat(path); os.IsNotExist(err) {
			err := os.MkdirAll(path, fs.ModePerm)
			if err != nil {
				helper.WriteError(c, err)
				return
			}
		}

		if err := c.SaveUploadedFile(file, dst); err != nil {
			helper.WriteError(c, err)
			return
		}

		req.Location = dst
		req.UploadBy = "wihlarkop"

		resp, err := inport.Execute(c.Copy().Request.Context(), req)
		if err != nil {
			helper.WriteError(c, err)
			return
		}

		helper.WriteSuccess(c, "Success get file", resp, nil)
	}
}

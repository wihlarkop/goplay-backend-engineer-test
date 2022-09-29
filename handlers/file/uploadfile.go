package file

import (
	"fmt"
	"io/fs"
	"log"
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
		var tokens helper.TokenRequest

		// ? Binding Request
		if err := c.ShouldBind(&req); err != nil {
			log.Println(err)
			helper.WriteError(c, err)
			return
		}

		token, _ := c.Get("user")
		err := helper.TransformInterfaceToAnother(token, &tokens)
		log.Println(err)

		if err != nil {
			helper.WriteError(c, err)
			return
		}

		file, err := c.FormFile("file")
		log.Println(err)

		if err != nil {
			helper.WriteError(c, err)
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
			log.Println(err)

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
		req.UploadBy = tokens.Name

		resp, err := inport.Execute(c.Copy().Request.Context(), req)
		if err != nil {
			helper.WriteError(c, err)
			return
		}

		helper.WriteSuccess(c, "Success get file", resp, nil)
	}
}

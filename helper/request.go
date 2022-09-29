package helper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func UnmarshalJSON(c *gin.Context, dest interface{}) error {
	if c.Request.Body == nil {
		return nil
	}

	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		WriteError(c, ErrJSONParse)
		return err
	}

	if err := json.Unmarshal(bodyBytes, dest); err != nil {
		WriteError(c, ErrJSONParse)
		return err
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return nil
}

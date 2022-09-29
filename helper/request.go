package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func PrettyPrint(logName string, data interface{}) {

	switch v := data.(type) {
	case []byte:
		var obj map[string]interface{}

		json.Unmarshal(v, &obj)
		data = obj
	}

	fmt.Println(".")
	fmt.Println(".")
	fmt.Println(".")
	fmt.Printf("--------------------------------%s--------------------------------\n", logName)
	jr, _ := json.MarshalIndent(data, "| ", "   ")
	fmt.Println("|", string(jr))
	fmt.Printf("--------------------------------%s--------------------------------\n", logName)
	fmt.Println(".")
	fmt.Println(".")
	fmt.Println(".")
}

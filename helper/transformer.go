package helper

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func TransformInterfaceToAnother(source interface{}, entity interface{}) error {
	userByte, err := json.Marshal(source)
	if err != nil {
		return err
	}

	err = json.Unmarshal(userByte, entity)
	if err != nil {
		return err
	}

	return nil
}

func TransformStructToUrlValue(obj interface{}) url.Values {
	var objMap map[string]interface{}

	TransformInterfaceToAnother(obj, &objMap)

	data := url.Values{}
	for i, v := range objMap {
		data.Set(i, fmt.Sprint(v))
	}

	return data
}

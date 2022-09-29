package helper

import "encoding/json"

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

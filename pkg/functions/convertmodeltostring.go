package functions

import (
	"encoding/json"
	"hsbc_connector/pkg/model"
)

func ConvertModelToString(responsedata *model.Responsedata) (string, error) {
	responseData,err := json.Marshal(responsedata)
	return string(responseData), err
}

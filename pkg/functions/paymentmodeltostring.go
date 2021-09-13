package functions

import (
	"encoding/json"
	"hsbc_connector/pkg/model"
)

func PaymentModelToString(requestpaymentdata string) (*model.Requeststring, error) {
	var jsonModel *model.Requeststring
	err := json.Unmarshal([]byte(requestpaymentdata), jsonModel)
	return jsonModel, err
}

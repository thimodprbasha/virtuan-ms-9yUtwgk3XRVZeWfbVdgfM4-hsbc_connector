package functions

import "github.com/go-resty/resty/v2"

func GetInternationalPayment(internationalpaymentid string, x_fapi_financial_id string, authorization string) (string, error) {

	headers := map[string]string{
		"Accept":"*/*",
		"Accept-Encoding": "gzip, deflate, br",
		"Connection":"keep-alive",
		"InternationalPaymentId" : internationalpaymentid,
		"x-fapi-financial-id": x_fapi_financial_id,
		//"Authorization" : authorization,
	}
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headers).
		SetAuthToken(authorization).
		Get("/search_result")

	return resp.String(), err
}

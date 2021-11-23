package picpay

import (
	"encoding/json"
	"os"

	"github.com/eduardo-mior/picpay-sdk-go/internal/request"
)

const BASEURL = "https://appws.picpay.com/ecommerce/public/payments/"

// CreatePayment é o método responsável por criar um pagamento no picpay.
func CreatePayment(paymentRequest PaymentRequest, picpayToken ...string) (*PaymentResponse, *ErrorResponse, error) {

	params := request.Params{
		Method:  "POST",
		Body:    paymentRequest,
		Headers: map[string]interface{}{"x-picpay-token": getPicpayToken(picpayToken...)},
		URL:     BASEURL,
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode != 200 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var paymentResponse PaymentResponse
	err = json.Unmarshal(response.RawBody, &paymentResponse)
	return &paymentResponse, nil, err
}

// ConsultStatusPayment é o método responsável por consultar o status de um pagamento no picpay.
func ConsultStatusPayment(referenceID string, picpayToken ...string) (*ConsultStatusResponse, *ErrorResponse, error) {

	params := request.Params{
		Method:  "GET",
		Headers: map[string]interface{}{"x-picpay-token": getPicpayToken(picpayToken...)},
		URL:     BASEURL + referenceID + "/status",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode != 200 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var consultStatusResponse ConsultStatusResponse
	err = json.Unmarshal(response.RawBody, &consultStatusResponse)
	return &consultStatusResponse, nil, err
}

// CancelPayment é o método responsável por cancelar um pagamento no picpay.
func CancelPayment(referenceID string, authorizationID *string, picpayToken ...string) (*CancelPaymentResponse, *ErrorResponse, error) {

	params := request.Params{
		Method:  "POST",
		Headers: map[string]interface{}{"x-picpay-token": getPicpayToken(picpayToken...)},
		URL:     BASEURL + referenceID + "/cancellations",
		Body:    CancelPaymentRequest{authorizationID},
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode != 200 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var cancelPaymentResponse CancelPaymentResponse
	err = json.Unmarshal(response.RawBody, &cancelPaymentResponse)
	return &cancelPaymentResponse, nil, err
}

// parseError é a função que pega os dados do erro do PicPay e retorna em formato de Struct.
func parseError(body []byte) (*ErrorResponse, error) {
	var errResponse ErrorResponse
	if err := json.Unmarshal(body, &errResponse); err != nil {
		return nil, err
	}
	return &errResponse, nil
}

// getPicpayToken é a função responsável por retornar o Token do PicPay.
// Caso tenha sido passado um token por parametro pegamos o token passado por parametro, se não pegamos da variavel de ambiente X_PICPAY_TOKEN.
func getPicpayToken(picpayToken ...string) string {
	if len(picpayToken) > 0 {
		return picpayToken[0]
	} else {
		return os.Getenv("X_PICPAY_TOKEN")
	}
}

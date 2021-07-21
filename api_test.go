package picpaysdkgo

import (
	"os"
	"testing"
)

// Definindo a variavel de ambiente X_PICPAY_TOKEN que é usada pelo SDK
func init() {
	os.Setenv("X_PICPAY_TOKEN", "seu-token-de-integracao-do-picpay")
}

// Testando geração de um pagamento
func TestSuccessOnCreatePayment(t *testing.T) {

	response, errResponse, err := CreatePayment(PaymentRequest{
		ReferenceID: "test-00001",
		Value:       1,
		Buyer: Buyer{
			FirstName: "Eduardo Mior",
			Document:  "12345678909",
			Email:     "raniellimontagna@hotmail.com",
			Phone:     "54991343192",
		},
		CallbackURL: "https://62c729c05614.ngrok.io/picpay",
	})

	if err != nil {
		t.Error("Erro inesperado!")
		t.Error(err.Error())

	} else if errResponse != nil {
		t.Error("Erro não tratado PicPay!")
		t.Error(errResponse.Message)
		t.Error(errResponse.Errors)

	} else {
		t.Log(response.PaymentURL) // Sucesso!
	}

}

// Testando tratamento de erro na geração de um pagamento (email e CPF inválids)
func TestFieldErrorOnCreatePayment(t *testing.T) {

	response, errResponse, err := CreatePayment(PaymentRequest{
		ReferenceID: "test-00002",
		Value:       1,
		Buyer: Buyer{
			FirstName: "Eduardo Mior",
			Document:  "non-valid-cpf",
			Email:     "non-valid-email",
			Phone:     "54991343192",
		},
		CallbackURL: "https://62c729c05614.ngrok.io/picpay",
	})

	if err != nil {
		t.Error("Erro inesperado!")
		t.Error(err.Error())

	} else if errResponse != nil {
		t.Log("Erro caputado com sucesso!") // Sucesso
		t.Log(errResponse.Message)
		t.Log(errResponse.Errors)

	} else {
		t.Error("Erro não capturado!")
		t.Error(response)
	}

}

// Testando cancelamento de um pagamento
func TestSuccessOnCancelPayment(t *testing.T) {

	authorizationID := "60f84028178ff000260addb9"
	response, errResponse, err := CancelPayment("test-00002", &authorizationID)

	if err != nil {
		t.Error("Erro inesperado!")
		t.Error(err.Error())

	} else if errResponse != nil {
		t.Error("Erro não tratado PicPay!")
		t.Error(errResponse.Message)
		t.Error(errResponse.Errors)

	} else {
		t.Log(response.CancellationID) // Sucesso!
	}

}

// Testando erro no cancelamento de um pagamento (pagamento já cancelado)
func TestErrorOnCancelPayment(t *testing.T) {

	authorizationID := "60f84028178ff000260addb9"
	response, errResponse, err := CancelPayment("test-00002", &authorizationID)

	if err != nil {
		t.Error("Erro inesperado!")
		t.Error(err.Error())

	} else if errResponse != nil {
		t.Log("Erro caputado com sucesso!") // Sucesso
		t.Log(errResponse.Message)
		t.Log(errResponse.Errors)

	} else {
		t.Error("Erro não capturado!")
		t.Error(response)
	}

}

// Testando consulta de um pagamento
func TestSuccessOnConsultStatusPayment(t *testing.T) {

	response, errResponse, err := ConsultStatusPayment("test-00002")

	if err != nil {
		t.Error("Erro inesperado!")
		t.Error(err.Error())

	} else if errResponse != nil {
		t.Error("Erro não tratado PicPay!")
		t.Error(errResponse.Message)
		t.Error(errResponse.Errors)

	} else {
		t.Log(response.Status) // Sucesso!
	}

}

// Testando erro na consulta de um pagamento (pagamento inexistente)
func TestErrorOnConsultStatusPayment(t *testing.T) {

	response, errResponse, err := ConsultStatusPayment("test-inexistente")

	if err != nil {
		t.Error("Erro inesperado!")
		t.Error(err.Error())

	} else if errResponse != nil {
		t.Log("Erro caputado com sucesso!") // Sucesso
		t.Log(errResponse.Message)
		t.Log(errResponse.Errors)

	} else {
		t.Error("Erro não capturado!")
		t.Error(response)
	}

}

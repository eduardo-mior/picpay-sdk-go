package picpay

import (
	"time"
)

// PaymentRequest é a struct que é usada para fazer a request de um novo pagamento para o PicPay.
type PaymentRequest struct {
	ReferenceID  string     `json:"referenceId"`  // Nosso ID de controle interno
	CallbackURL  string     `json:"callbackUrl"`  // URL do Webhook que é chamada quando o Status do pagamento é atualizado
	ReturnURL    *string    `json:"returnUrl"`    // URL para onde o usuário sera direcionado ao efetuar o pagamento
	Value        float64    `json:"value"`        // Valor do pagamento
	ExpiresAt    *time.Time `json:"expiresAt"`    // Data de expiração do pagamento
	Buyer        Buyer      `json:"buyer"`        // Informações do pagador
	PurchaseMode string     `json:"purchaseMode"` // Indica se o pagamento é online ou in-store (por padrão se não for informado nada é in-store)
	Channel      string     `json:"channel"`      // Indica qual o estabelecimento (exclusivo à integradores que possuem mais de um estabelecimento comercial)
}

// Buyer é a struct que é usada para identificar para quem é o pagamento.
type Buyer struct {
	FirstName string `json:"firstName"` // Primeiro nome do pagador (pode ser informado o nome completo aqui)
	LastName  string `json:"lastName"`  // Sobrenome do pagador (pode ser em branco)
	Document  string `json:"document"`  // CPF/CNPJ do pagador
	Email     string `json:"email"`     // Email do pagador (opcional)
	Phone     string `json:"phone"`     // Telefone do pagador (opcional)
}

// PaymentResponse é a struct que é usada para receber os dados do request de novo pagamento do PicPay.
type PaymentResponse struct {
	ReferenceID string    `json:"referenceId"` // Nosso ID de controle interno
	PaymentURL  string    `json:"paymentUrl"`  // Link para efetuar o pagamento
	ExpiresAt   time.Time `json:"expiresAt"`   // Data de expiração do pagamento
	QrCode      struct {
		Content string `json:"content"` // Conteúdo interno do QRCode
		Base64  string `json:"base64"`  // QRCode no formato Base64 (png)
	} `json:"qrcode"`
}

// ConsultStatusResponse é a struct que é usada para receber os dados da consulta de status do pagamento do PicPay.
type ConsultStatusResponse struct {
	AuthorizationID *string `json:"authorizationId"` // ID de controle do pagamento do PicPay
	ReferenceID     string  `json:"referenceId"`     // Nosso ID de controle interno
	// created - registro criado;
	// expired - prazo para pagamento expirado ou ordem cancelada;
	// analysis - pago e em processo de análise anti-fraude;
	// paid - pago;
	// completed - pago e saldo disponível;
	// refunded - pago e devolvido;
	// chargeback - pago e com chargeback;
	Status string `json:"status"` // Status atual do pagamento
}

// CancelPaymentRequest é a struct que é usada para fazer o request de um cancelamento de um pagamento no PicPay.
type CancelPaymentRequest struct {
	AuthorizationID *string `json:"authorizationId"` // ID de controle do pagamento do PicPay (opcional, só deve ser informado se o pagamento já tiver sido feito)
}

// CancelPaymentResponse é a struct que é usada para receber os dados do request de cancelamento de pagamento do PicPay.
type CancelPaymentResponse struct {
	ReferenceID    string `json:"referenceId"`    // Nosso ID de controle interno
	CancellationID string `json:"cancellationId"` // ID de controle do cancelamento do PicPay
}

// WebhookResponse é a struct que é usada para receber os dados do request que o PicPay faz para o nosso webhook.
type WebhookResponse struct {
	ReferenceID     string `json:"referenceId"`     // Nosso ID de controle interno
	AuthorizationID string `json:"authorizationId"` // ID de controle do pagamento do PicPay
}

// ErrorResponse é a struct que é usada para receber os retornos de erro do PicPay.
type ErrorResponse struct {
	Message string `json:"message"` // Mensagem do erro que retornou
	// Detalhes dos erros de validação de campos
	Errors []struct {
		Field   string `json:"field"`   // Nome do campo
		Message string `json:"message"` // Mensagem de erro relacinada ao campo
	} `json:"errors"`
}

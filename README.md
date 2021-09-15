# picpay-sdk-go
SDK (não oficial) para consumir os serviços do PicPay em Golang.

## Exemplos
Criando um pagamento:
```go
response, picpayErr, err := picpay.CreatePayment(picpay.PaymentRequest{
    ReferenceID: "seu-id-interno-0001",
    Value:       1,
    Buyer: picpay.Buyer{
        FirstName: "Eduardo Mior",
        Document:  "12345678909",
        Email:     "eduardo-mior@hotmail.com",
        Phone:     "54991343192",
    },
    CallbackURL: "http://localhost/webhook/picpay",
})

if err != nil {
    // Erro inesperado
} else if picpayErr != nil {
    // Erro retornado do PicPay
} else {
    // Sucesso!
}
```

Consultando o status de um pagamento:
```go
response, picpayErr, err := picpay.ConsultStatusPayment("seu-id-interno-0001")

if err != nil {
    // Erro inesperado
} else if picpayErr != nil {
    // Erro retornado do PicPay
} else {
    // Sucesso!
}
```

Cancelando um pagamento:
```go
// Se o pagamento já foi efetuado pelo cliente você deve passar por parâmetro o authorizationID 
// do pagamento que é retornado pelo Webhook, caso contrado você deve passar nil
authorizationID := "ecf85122178ff12c260accb9"
response, picpayErr, err := picpay.CancelPayment("seu-id-interno-0001", &authorizationID)

if err != nil {
    // Erro inesperado
} else if picpayErr != nil {
    // Erro retornado do PicPay
} else {
    // Sucesso!
}
```

## Ajuda
O SDK possui suporte a todas as funções disponíveis na API, sendo elas:
 - Criação de um pagamento
 - Cancelamento de um pagamento
 - Consulta de status de um pagamento
Além de possuir suporte a todas as funções, o SDK também possui todas as Struct de todos os modelos documentados na documentação do PicPay.
###
O SDK precisa obrigatóriamente para funcionar, de uma variavel de ambiente chamada `X_PICPAY_TOKEN` que contém o seu Token de integração do PicPay. Esse Token pode ser obtido no painel do PicPay para empresas, ou no antigo painel para Lojistas. Para setar a variavel ambiente você pode usar a função `os.Setenv("X_PICPAY_TOKEN", "seu-token...")` ou você pode usar um arquivo `.env` e usar um pacote para gerenciar as variaveis de ambiente, como por exemplo o [Gotenv](https://github.com/subosito/gotenv).
###
Todas as funções do SDK podém retornar um `error` genérico do GO, este erro esta sempre relacionado a erros do GO, como por exemplo falha ao tentar dar parse em um JSON, além disso todas as funções também podem retornar um `ErrorResponse` que é uma Struct de erro retornada do PicPay, que pode retornar quando você não envia um campo obrigatório por exemplo.
###
Atenção! Você deve implementar manualmente o Webhook que recebe as atualizações de Status do pagamento usando o seu Framework WEB de prefencia (lembrando que o SDK possui a Struct `WebhookResponse` que pode ajudar no recebimento dos dados).
## Documentação oficial
Para mais duvidas consulte a [documentação oficial do PicPay](https://picpay.github.io/picpay-docs-digital-payments/).
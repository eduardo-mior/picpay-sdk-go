![GoLang PicPay](https://i.imgur.com/fvtdEXO.png)
# PicPay SDK para Go
SDK (não oficial) para consumir os serviços do PicPay em Golang.

## 📲  Instalação
Para baixar o SDK basta utilizar o seguinte comando:
```bash
$ go get -u github.com/eduardo-mior/picpay-sdk-go
```

## 🛠 Funcionalidades do SDK
Funcionalidades disponíbilizadas no SDK:
- Criação de um pagamento
- Consulta status de um pagamento
- Cancelamento de um pagamento

## 🌟  Começando 
Para começar você deve fazer o `import` do SDK, para isso basta adicionar a seguinte linha no seu código:
```go
import "github.com/eduardo-mior/picpay-sdk-go"
```
Feito isso já esta tudo pronto para você começar a utilizaro SDK!

## 💻  Exemplos de uso
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
}, "seu-token")

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
response, picpayErr, err := picpay.ConsultStatusPayment("seu-id-interno-0001", "seu-token")

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
response, picpayErr, err := picpay.CancelPayment("seu-id-interno-0001", &authorizationID, "seu-token")

if err != nil {
    // Erro inesperado
} else if picpayErr != nil {
    // Erro retornado do PicPay
} else {
    // Sucesso!
}
```

## 🙋🏻‍♂️  Ajuda
O SDK possui suporte a todas as funções disponíveis na API, sendo elas:
 - Criação de um pagamento
 - Cancelamento de um pagamento
 - Consulta de status de um pagamento
Além de possuir suporte a todas as funções, o SDK também possui todas as Struct de todos os modelos documentados na documentação do PicPay.
###
O SDK precisa obrigatóriamente do X-Token do PicPay para funcionar. O X-Token pode ser passado por parametro em todas as funções, ou pode ser definido na variavel de ambiente `X_PICPAY_TOKEN` que irá conter o seu Token de integração do PicPay. Esse Token pode ser obtido no [painel do PicPay para empresas na aba "Integrações"](https://painel-empresas.picpay.com/integracoes), ou no antigo [painel para Lojistas na aba "Integração"](https://lojista.picpay.com/ecommerce-token). Para setar a variavel ambiente você pode usar a função `os.Setenv("X_PICPAY_TOKEN", "seu-token...")` ou você pode usar um arquivo `.env` e usar um pacote para gerenciar as variaveis de ambiente, como por exemplo o [Gotenv](https://github.com/subosito/gotenv).
###
Todas as funções do SDK podém retornar um `error` genérico do GO, este erro esta sempre relacionado a erros do GO, como por exemplo falha ao tentar dar parse em um JSON, além disso todas as funções também podem retornar um `ErrorResponse` que é uma Struct de erro retornada do PicPay, que pode retornar quando você não envia um campo obrigatório por exemplo.
###
Atenção! Você deve implementar manualmente o Webhook que recebe as atualizações de Status do pagamento usando o seu Framework WEB de prefencia (lembrando que o SDK possui a Struct `WebhookResponse` que pode ajudar no recebimento dos dados).
###
Antigmente após receber a notificação do Webhook do PicPay você deveria chamar a função `ConsultStatusPayment()` passando por parametro o `referenceId`, que é o ID que foi enviado pelo PicPay na struct `WebhookResponse` para poder consular o status do pagamento.
A partir de novembro de 2021 a Struct enviada pelo Webhook foi mudada complementada e agora o Status já vem com a respsota do Webhook, sendo assim não é mais necessário chamar a função `ConsultStatusPayment()`.
## 📚 Documentação oficial
Para mais duvidas consulte a [documentação oficial do PicPay](https://picpay.github.io/picpay-docs-digital-payments/).

# picpay-sdk-go
SDK (não oficial) para consumir os serviços do PicPay em Golang.

## Ajuda
O SDK possui suporte a todas as funções disponíveis na API, sendo elas:
 - Criação de um pagamento
 - Cancelamento de um pagamento
 - Consulta de status de um pagamento
Além de possuir suporte a todas as funções, o SDK também possui todas as Struct de todos os modelos documentados na documentação do PicPay.
###
O SDK precisa obrigatóriamente para funcionar, de uma variavel de ambiente chamada `X_PICPAY_TOKEN` que contém o seu Token de integração do PicPay. Esse Token é pode ser obtido no painel do PicPay para empresas, ou no antigo painel para Lojistas.
###
Todas as funções do SDK podém retornar um `error` genérico do GO, este erro esta sempre relacionado a erros do GO, como por exemplo falha ao tentar dar parse em um JSON, além disso todas as funções também podem retornar um `ErrorResponse` que é uma Struct de erro retornada do PicPay, que pode retornar quando você não envia um campo obrigatório por exemplo.
###
Atenção! Você deve implementar manualmente o Webhook que recebe as atualizações de Status do pagamento usando o seu Framework WEB de prefencia (lembrando que o SDK possui a Struct `WebhookResponse` que pode ajudar no recebimento dos dados).
## Documentação oficial
Para mais duvidas consulte a [documentação oficial do PicPay](https://picpay.github.io/picpay-docs-digital-payments/).
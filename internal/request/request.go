package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

// Params Parâmetros para o método Request
type Params struct {
	Metodo       string
	URL          string
	Body         interface{}
	Headers      map[string]interface{}
	Timeout      int
	PathParams   []interface{}
	QueryParams  map[string]interface{}
	BasicAuth    *BasicAuth
	HandleErrors *bool
}

// BasicAuth Usuário e senha usados na autenticação por BasicAuth
type BasicAuth struct {
	Username string
	Password string
}

// Response Retorno do método Request
type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       map[string]interface{}
	RawBody    []byte
}

// New Efetua uma requsição http para uma API, microservice ou outro.
func New(params Params) (*Response, error) {
	var body *bytes.Reader

	// Verificando caso a requisição possua body então encodamos ele em JSON
	if params.Body != nil {
		data, err := json.Marshal(params.Body)

		if err != nil {
			return nil, err
		}

		body = bytes.NewReader(data)
	}

	// Verificando caso a requisição possua PathParams então adicionamos na URL separados por /
	if len(params.PathParams) > 0 {
		params.URL = strings.TrimSuffix(params.URL, "/")

		for _, v := range params.PathParams {
			params.URL += "/" + toString(v)
		}

	}

	// Verificando caso a requisição possua QueryParams então adicionamos na URL
	if len(params.QueryParams) > 0 {
		query := url.Values{}

		for k, v := range params.QueryParams {
			query.Add(k, toString(v))
		}

		params.URL += "?" + query.Encode()
	}

	var request *http.Request
	var err error

	// Instanciando a requisição para depois o cliente pode executa-la.
	// Caso a requisição possua body então passamos a variavel do body, caso contrario passamos nil.
	// Não podemos passar a variavel body quando ela é nil porque se não da erro interno do GO por causa dos nils tipados e nils não tipados.
	if body == nil {
		request, err = http.NewRequest(params.Metodo, params.URL, nil)
	} else {
		request, err = http.NewRequest(params.Metodo, params.URL, body)
	}

	if err != nil {
		return nil, err
	}

	if params.BasicAuth != nil {
		request.SetBasicAuth(params.BasicAuth.Username, params.BasicAuth.Password)
	}

	// Setando o header que indica que o body esta sendo enviado em formato JSON.
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	// Verificando se a requsição possui headers então setamos todos na requisição.
	if len(params.Headers) > 0 {

		for header, value := range params.Headers {
			request.Header.Set(header, toString(value))
		}

	}

	// Instanciando o client que ira executar a requsição.
	client := &http.Client{}

	// Verificando se algum timeout foi passado por parametro, caso não tenha sido passado então setamos 40 por padrão.
	if params.Timeout == 0 {
		params.Timeout = 40
	}

	// Setando o timeout no client.
	client.Timeout = time.Duration(params.Timeout) * time.Second

	// Executando a requisição.
	res, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// Lendo os headers da resposta que veio da API.
	headers := make(map[string]string)
	for name, values := range res.Header {
		headers[name] = values[0]
	}

	// Lendo a resposta veio da API.
	rawBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Decodificando a resposta da API.
	var ret map[string]interface{}
	err = json.Unmarshal(rawBody, &ret)

	// Verificando se deu algum erro ao fazer o Unmarshal do JSON.
	if err != nil {

		// Se tiver dado algum problema ao decodar a resposta da API, então nós verificamos, caso tenha retornado alguma resposta da API
		// então nós montamos uma mensagem de erro customizada, com a mensagem de erro original mais o Body da resosta, caso contrario
		// só retorna o erro original.
		if len(rawBody) > 0 {
			return nil, errors.New(err.Error() + " - response: " + string(rawBody))
		}
		return nil, err
	}

	return &Response{StatusCode: res.StatusCode, RawBody: rawBody, Body: ret, Headers: headers}, nil
}

// toString é o método responsável por retornar o valor de uma interface (pode ser ponteiro ou não) em string.
func toString(v interface{}) string {
	if v == nil {
		return ""
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		return fmt.Sprintf("%v", rv.Elem())
	}

	return fmt.Sprintf("%v", rv)
}

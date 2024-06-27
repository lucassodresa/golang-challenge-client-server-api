# Currency Converter

Neste desafio vamos aplicar o que aprendemos sobre webserver http, contextos,
banco de dados e manipulação de arquivos com Go.

Você precisará nos entregar dois sistemas em Go:

- server.go
- client.go

## Server

### How to run

```bash
cd server
go mod tidy
go run .
```

### Requirements

- Create an endpoint /cotacao on port 8080 to handle requests from client.go.
- Consume the exchange rate API at https://economia.awesomeapi.com.br/json/last/USD-BRL.
- Use the "context" package to manage a maximum timeout of 200ms for calling the exchange rate API.
- Return the exchange rate result in JSON format to the client.
- Use the "context" package to manage a maximum timeout of 10ms for persisting the exchange rate data in an SQLite database.
- Record each received exchange rate in the SQLite database.
- Log an error if any of the context operations (API call, database persistence) time out or if there are any other execution issues.

## Client

### How to run

```bash
cd client
go mod tidy
go run .
```

### Requirements

- Make an HTTP request to server.go to get the dollar exchange rate.
- Use the "context" package to manage a maximum timeout of 300ms for receiving the response from server.go.
- Extract the current exchange rate value (the "bid" field of the JSON) from the response.
- Save the current exchange rate in a file named "cotacao.txt" in the format: Dollar: {value}.
- Log an error if the context times out or if there are any other execution issues.

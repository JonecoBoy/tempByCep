# Executar
- executar algum dos executaveis ou via go run tempByCep.go
- Enviar alguma request para localhost:8080/temp/"cepCode"
- O retorno aparecerá no console e também na resposta.

# Teste Automatizado

## Teste da api / integracao
```
go test -v ./pkg/tempByCep.go ./pkg/tempByCep_test.go
```

## Teste individuais

Testes apis e metodos de cep
```
go test -v ./pkg/external/brasilApiCep.go ./pkg/external/viaCEP.go ./pkg/external/cepAPIs_test.go
```
Testes de metodos de temperatura
```
go test -v ./pkg/external/weatherApi.go ./pkg/external/weatherApi_test.go
```

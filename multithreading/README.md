
# Consulta de CEP com a Resposta Mais Rápida

Neste desafio, o objetivo é utilizar conceitos de *multithreading* e consumo de *APIs* em Go para buscar informações de CEP com o menor tempo de resposta entre duas fontes distintas.

## Objetivo

- Fazer duas requisições simultâneas a APIs públicas de CEP.
- Acatar e exibir a resposta da API que responder primeiro.
- Exibir qual API respondeu.
- Respeitar um timeout de 1 segundo.

## APIs Utilizadas

- [BrasilAPI](https://brasilapi.com.br/api/cep/v1/)
- [ViaCEP](http://viacep.com.br/ws/)

As URLs seguem o formato:
- BrasilAPI: `https://brasilapi.com.br/api/cep/v1/{CEP}`
- ViaCEP: `http://viacep.com.br/ws/{CEP}/json/`

## Exemplo de Execução

```bash
go run  cmd/main.go <CEP>
go run  cmd/main.go 13340894
```

Saída esperada:
```txt
Resposta da API: BrasilAPI
CEP: 01153-000
Logradouro: Rua Jaraguá
Bairro: Bom Retiro
Cidade: São Paulo
Estado: SP
```

Ou, se ambas falharem:

```txt
Erro: requisições excederam o tempo limite de 1 segundo.
```

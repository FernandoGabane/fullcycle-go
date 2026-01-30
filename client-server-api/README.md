# Cotação Dólar

Este projeto é composto por dois arquivos principais: `server.go` e `client.go`. 
Eles trabalham em conjunto para buscar e registrar a cotação atual do dólar utilizando conceitos como HTTP server, contexto (`context.Context`), persistência em SQLite e manipulação de arquivos.

---

## Estrutura do Projeto

- `server.go`: servidor que consulta a cotação do dólar em uma API externa e salva no banco de dados.
- `client.go`: cliente que consome a cotação do servidor e grava o resultado em um arquivo.


---

##  Como executar

### 1. Executar o servidor
O servidor será iniciado na porta `8080`.

```bash
go run server.go
```

---

### 2. Executar o cliente
O cliente realiza a requisição para o servidor e salva o valor da cotação em um arquivo chamado `cotacao.txt`.

```bash
go run client.go
```

---

## Exemplo de resposta

- Conteúdo do arquivo `cotacao.txt`:
```
Dólar: 5.2647
```

---

## Contextos e Timeouts

| Componente       | Timeout      | Descrição                                                                 |
|------------------|--------------|---------------------------------------------------------------------------|
| API externa      | `200ms`      | Tempo máximo para obter resposta da [AwesomeAPI](https://economia.awesomeapi.com.br/json/last/USD-BRL) |
| SQLite           | `10ms`       | Tempo limite para gravar a cotação no banco de dados                     |
| Cliente          | `300ms`      | Tempo máximo para receber a resposta do servidor                         |

Em caso de timeout, mensagens de erro serão logadas no terminal.

---

## Como validar

### Validação funcional:
- Verifique o terminal do servidor e do cliente para confirmar as ações.
- Confirme a existência do arquivo `cotacao.txt` com o valor do dólar.
- Verifique o banco SQLite `cotacoes.db` para confirmar o registro da cotação.

### Validação de timeouts:
- Modifique artificialmente os tempos de sleep ou conexão (simular delay na API ou no DB) e observe os logs.

---

## Banco de dados SQLite

O arquivo `cotacoes.db` será criado automaticamente. Para verificar os dados salvos:

```bash
sqlite3 cotacoes.db
sqlite> SELECT * FROM cotacoes;
```
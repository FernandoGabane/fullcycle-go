# Load Tester
Criar um sistema para realizar testes de carga em um serviço web. 
O sistema realiza requisições HTTP conforme parâmetros definidos pelo usuário e apresenta um relatório ao final da execução.

---

## Parâmetros via CLI

O programa aceita os seguintes parâmetros:

- `--url`: URL do serviço a ser testado.
- `--requests`: Número total de requisições a serem realizadas.
- `--concurrency`: Número de chamadas simultâneas (concorrência).

### Exemplo de uso via CLI

```bash
bo 
go run cmd/main.go --url=http://google.com --requests=100 --concurrency=10
```

---

## Execução com Docker

Você pode executar a aplicação usando Docker:

```bash
docker build -t load-tester .
docker run load-tester --url=http://google.com --requests=100 --concurrency=10
```

---

## Relatório gerado

Após os testes, o sistema exibirá:

- Tempo total gasto na execução
- Quantidade total de requisições realizadas
- Total de requisições com status **HTTP 200**
- Distribuição de outros códigos de status HTTP (**404**, **500**, etc.)

---
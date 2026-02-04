# CEP Weather API

API desenvolvida que recebe um **CEP brasileiro**, identifica a **cidade** correspondente e retorna a **temperatura atual** nos formatos **Celsius**, **Fahrenheit** e **Kelvin**.

---

##  Funcionalidades

- Validação de CEP (8 dígitos numéricos)
- Consulta de localização via **ViaCEP**
- Consulta de temperatura atual via **WeatherAPI**
- Conversão automática de temperatura:
    - Celsius → Fahrenheit
    - Celsius → Kelvin
- Respostas HTTP padronizadas
- Testes automatizados
- Docker e Docker Compose
- Pronto para Google Cloud Run

---

## Como Executar

### Com Go
```bash
export WEATHER_API_KEY=your_api_key
go run cmd/api/main.go
```

### Com Docker
```bash
docker build -t cep-weather .
docker run -p 8080:8080 -e WEATHER_API_KEY=your_api_key cep-weather
```

---

## Endpoint

```
GET /weather?cep=01001000
```

---

## Sucesso — HTTP 200
```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.5
}
```

---

## Erros

- **422**: invalid zipcode
- **404**: can not find zipcode

---

## Testes
```bash
go test ./...
```

# Obs
Foi solicitado o deploy da aplicação no Google Cloud e o envio do link da API em produção.
Para concluir esse passo, é necessário criar uma conta no Google Cloud com a ativação do faturamento, o que exige o cadastro de um cartão de crédito e a realização de um pré-pagamento no valor aproximado de R$ 50,00.
No momento, optei por não prosseguir com essa etapa, uma vez que a conta seria utilizada apenas para essa finalidade específica e o valor não é estornado.
Entendo que a aplicação desenvolvida atende integralmente aos requisitos propostos, estando funcional, com arquitetura organizada, boas práticas aplicadas, testes automatizados e pronta para deploy em ambiente de produção, o que considero suficiente para a avaliação dos critérios técnicos desta etapa.
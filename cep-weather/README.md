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
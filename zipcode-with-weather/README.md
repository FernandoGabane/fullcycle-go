#  Sistema CEP + Clima com Observabilidade (OTEL + Zipkin)

Este projeto é composto por dois serviços escritos em Go, consultar a **temperatura atual** (Celsius, Fahrenheit, Kelvin) com base em um **CEP** fornecido via requisição HTTP.
Além disso, ele implementa **OpenTelemetry (OTEL)** para rastreamento distribuído e usa **Zipkin** para visualizar os traces entre os serviços.

---

## Estrutura

- `service-a`: Recebe o CEP via POST e o encaminha ao `service-b`
- `service-b`: Consulta a localização (viaCEP) e clima (WeatherAPI), retorna os dados formatados
- `otel-collector`: Coleta os spans dos serviços e os envia para o Zipkin
- `zipkin`: Interface web para visualizar os traces

---

##  Pré-requisitos

- Docker
- Docker Compose
- Conta gratuita na [WeatherAPI](https://www.weatherapi.com/) (obrigatória para obter dados de clima)

---

##  Como rodar

### 1. Obtenha sua chave da WeatherAPI

Crie uma conta gratuita em https://www.weatherapi.com e copie sua **API Key**.

---

### 2. Configure a chave no `docker-compose.yml`

No serviço `service-b`, a variável de ambiente `WEATHER_API_KEY` precisa estar configurada. Existem duas formas:

#### ✅ Opção recomendada (via terminal):
```bash
WEATHER_API_KEY=<SUA_CHAVE_WEATHERAPI> docker-compose up --build
```

Esse comando irá:
- Construir os containers Go (Service A e B)
- Subir o Zipkin
- Subir o OTEL Collector com o pipeline de rastreamento
- Conectar os serviços corretamente

---

## Como testar

Você pode fazer uma requisição usando `curl`:

```bash
curl -X POST http://localhost:8080/cep \
  -H "Content-Type: application/json" \
  -d '{"cep":"13340268"}'
```
 Saída esperada:
```json
{"city":"Indaiatuba","temp_C":0,"temp_F":32,"temp_K":273}
```

---

## Visualizar os traces (OpenTelemetry + Zipkin)

Abra no navegador:

👉 [http://localhost:9411](http://localhost:9411)

Você verá os **traces entre os serviços** A → B, com detalhes de duração, spans e tempo gasto em chamadas externas (viaCEP, WeatherAPI).
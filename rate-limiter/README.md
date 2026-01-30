# Rate Limiter
Este projeto implementa um **Rate Limiter**  para limitação de requisições por **IP** e por **Token de Acesso**, usando **Redis** como mecanismo de persistência, 
para ser usado como middleware em servidores HTTP.

## Configuração

### .env

```env
PORT=8080

# Limites por IP
IP_MAX_REQ=5

# Limites por TOKEN
TOKEN_MAX_REQ=10
BLOCK_DURATION_SECONDS=10

# Redis
REDIS_ADDR=redis:6379
```

## Exemplo de Execução

1. Suba o projeto:
   ```bash
   docker-compose up -d
   ```
   
2. Excute a applicação:
   ```bash
   go run  cmd/main.go
   ```


## Exemplo de Requisição
execute a requisição abaixo multiplas vezes até atingir o rate limiter. uma vez atingido, deve se esperar o block duration passar e as requisiçoes irão retonar OK novamente.

```bash
curl -H "API_KEY: abc123" http://localhost:8080/
```

## Resposta quando limite é excedido

```json
{
  "message": "you have reached the maximum number of requests or actions allowed within a certain time frame"
}
```

Status HTTP: `429 Too Many Requests`

## Estratégia de Persistência

O projeto usa uma **strategy interface** para tornar o mecanismo de persistência intercambiável. Por padrão, Redis é utilizado.


## Observações

- Tokens têm prioridade sobre o IP.
- Uma vez excedido o limite, o IP ou Token será bloqueado pelo tempo configurado.


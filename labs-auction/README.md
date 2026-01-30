#  Atualização: Fechamento Automático de Leilões

Atualização do sistema de leilões para permitir o **fechamento automático** de leilões após um tempo determinado via variável de ambiente.

---

##  Como Iniciar via Docker

1. **Configure a variável de ambiente no `.env`**:

```env
AUCTION_DURATION_SECONDS=30
```

> Isso define o tempo de duração de cada leilão em segundos.

2. **Suba os containers com Docker**:

```bash
docker-compose up --build
```

3. **Verifique se o container da aplicação está rodando**:

Acesse: [http://localhost:8000](http://localhost:8000) (ou a porta configurada).

---

## Como Testar

### Criar um leilão via API

```bash
curl -X POST http://localhost:8000/auction   -H "Content-Type: application/json"   -d '{
    "product_name": "MotoLoLa",
    "category": "Eletlonico",
    "description": "abliu tem que leva",
    "condition": 10
}'
```

- O leilão será automaticamente fechado após `AUCTION_DURATION_SECONDS`.

- Pesquise o ID do produto realizando a chamada abaixo.
```bash
curl --request GET \
     --url 'http://localhost:8080/auction?status=0'\
```

### Verificar se foi fechado

```bash
curl http://localhost:8000/auction/<AUCTION_ID>
```

O campo `status` será 1 = Closed após a quantidade de segundos estabelecida na variavel de ambiente AUCTION_DURATION_SECONDS,


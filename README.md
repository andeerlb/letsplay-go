# letsplay Microservice

Microserviço em Go para criação de jogadores, que chama outro microserviço para persistência.

## Estrutura do projeto

- `cmd/letsplay/main.go`: Ponto de entrada
- `internal/config`: Configurações com Viper (suporte a profiles demo/dev/prod)
- `internal/logger`: Logger estruturado com Zap
- `internal/model`: Models e structs
- `internal/client`: Client HTTP para o outro microserviço
- `internal/service`: Lógica de negócio
- `internal/handler`: Handlers HTTP (controllers)
- `internal/router`: Configuração das rotas
- `internal/server`: Setup do servidor HTTP
- `tests`: Testes unitários e integração (não implementados aqui)

## Como rodar

1. Instale o Go (>=1.20)

2. Configure variável de ambiente `APP_ENV` para escolher o ambiente (demo, dev, prod). Default: dev

3. Configure variáveis específicas para cada ambiente via arquivo `.env` na pasta `configs` (opcional), exemplo:

```
SERVER_PORT=8080
LOG_LEVEL=info
PLAYER_SERVICE_URL=http://localhost:9000
```

4. Baixe dependências:

```bash
go mod tidy
```

5. Rode o microserviço:

```bash
go run cmd/letsplay/main.go
```

6. Teste o endpoint:

```bash
curl -X POST http://localhost:8080/newplayer \
 -H 'Content-Type: application/json' \
 -d '{
     "email": "player@example.com",
     "password": "secret123",
     "primaryGame": {
         "type": "futebol",
         "position": "atacante"
     },
     "otherGame": [
         {"type": "basquete", "position": "armador"},
         {"type": "volei", "position": "levantador"}
     ]
 }'
```

## Como expandir

- Adicionar testes unitários e de integração
- Middleware de autenticação (JWT/OAuth)
- Monitoramento e métricas (Prometheus, OpenTelemetry)
- Configuração dinâmica via Consul/Etcd
- Dockerfile e CI/CD
- Melhor tratamento de erros e logs estruturados
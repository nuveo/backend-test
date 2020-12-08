# Teste de back-end para o Nuveo

### Requisitos

- Go 1.15
- Docker

### Instalando

Primeiro, levante uma instância do PostgreSQL:

```
sudo docker run --rm -name nuveotest2 -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres
```

Copie o esquema disponível em `util/` para o container e use-o no pqsl:

```
sudo docker cp util/database.sql nuveotest:/
sudo docker exec  nuveotest bash -c "PGPASSWORD=password psql -h localhost -U postgres -d postgres -f database.sql"
```

Logo após, levante uma instância do RabbitMQ:

```
sudo docker run -d --rm --hostname my-rabbittest --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:management
```

Feito isto, faça o build do repositório:

```
go build
```

Por fim, execute o arquivo!

```
./main
```

### Exemplos válidos de requisições:

```
POST http://localhost:8080/workflow
Content-Type: application/json

{
  "data": [
    {
      "texto": "Ronaldo",
      "numero": 1
    },
    {
      "texto": "Romário",
      "numero": 2
    },
    {
      "texto": "Ronaldinho",
      "numero": 3
    }
  ],
  "steps": [
    {
      "step": 1
    }
  ]
}
```

```
PATCH POST http://localhost:8080/workflow/{uuid}
Content-Type: application/json

{
  status: "inserted"
}
```

### TODO

- Implementar testes unitários nos serviços
- Implementar testes unitários nos handlers
- Criar Dockerfiles da API e do banco de dados
- Criar dockercompose.yml
- Documentar API com OpenAPI e GoSwagger
- Documentar serviços


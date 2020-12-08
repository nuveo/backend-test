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

### Comentários sobre o teste prático e o código do repositório

Apesar do pouco tempo disponível para implementar o projeto por causa da semana conturbada de fim de período e do trabalho, consegui implementar uma quantidade de funcionalidades boas e interessantes de serem discutidas, a exemplo:

- Extrai dos handlers os serviços, o que além de flexibilizar o código facilita bastante a implementação de testes (que estarão mais isolados e simples) e também a sua manutenção
- Os handlers utilizam o serviço de Workflow a partir de uma interface (WorkflowServer), cuja implementação é uma dependência injetada anteriormente no código. Este é um padrão interessante que facilita na hora dos testes de unidade dos handlers, já que eu posso utilizar um mock do serviço que não precisa atingir um banco de dados de teste ou uma fila de teste.
- As estruturas do corpo das requisições e das respostas são definidas explicitamente, garantindo que só seja transferido o necessário, o que é mais eficiente e seguro, além de facilitar a documentação da API 
- Separei um arquivo que lida só com a definição de rotas.
- Tentei nomear os métodos, variáveis, tipos e funções de forma legível e direta.
- Utilizei o RabbitMQ. Esse é um sistema que apesar de já ter indiretamente lidado com ele, nunca havia implementado algo com ele (ainda mais em Go). Foi um aprendizado prático bastante interessante.

Apesar disto, há alguns pontos que podem ser melhorados:

- Apesar de toda a estrutura de código favorável a testes corretos, eu não cheguei a implementá-los. Os testes são interessantes e a API que a linguagem disponibiliza torna tudo mais fácil. A falta de tempo foi o motivo pelo qual não os implementei. Enquanto eu termino de escrever o README eu poderia ter implementado um ou outro só de demonstração, mas acredito que uma suíte de testes mal planejada e apressada pode acabar sendo prejudicial e dar uma falsa e perigosa sensação de segurança.

- Utilização de Dockerfile e dockercompose.yml: Eu não cheguei a criar os arquivos porque ainda estou aprendendo a utilizar o Docker, mas este é um ponto a ser melhorado e que simplificaria muito o processo de build, teste e demonstração.

- Uma documentação dos endpoints com o GoSwagger seria bastante interessante. Como falado anteriormente, os tipos, requests e responses estão bem organizadas, portanto a geração de uma especificação seria relativamente fácil.

- Alguns procedimentos poderiam ter sido extaídos em funções auxiliadoras (ex.: verificação de erro). Assim diminuiria o boilerplate no código inerente à falta de algumas amenidades da linguagem.

Finalizando então, o trabalho foi uma ótima experiência que demonstrou alguns dos pontos fortes que eu tenho como programador (e que podem ser ainda mais fortes num ambiente adequado) e que também serviu de aprendizado, já que tive que trabalhar em algumas situações (como a do RabbitMQ que não tive a oportunidade de fazer antes.

# Backend Test

Develop the workflow's REST API following the specification bellow and document it.

## Delivery instructions

Clone this project and push a private repository in the [GitHub](https://github.com/), [Gitlab](https://about.gitlab.com/) or [Bitbucket](https://bitbucket.org/). When you want to our review, write any information that you think important in the README.md and send an email to talentos@nuveo.ai. We'll follow your instructions to run your code and look the outcome. 

## Defining a workflow

|Name|Type|Description|
|-|-|-|
|UUID|UUID|workflow unique identifier|
|status|Enum(inserted, consumed)|workflow status|
|data|JSONB|workflow input|
|steps|Array|name of workflow steps

## Endpoints

|Verb|URL|Description|
|-|-|-|
|POST|/workflow|insert a workflow on database and on queue and respond request with the inserted workflow|
|PATCH|/workflow/{UUID}|update status from specific workflow|
|GET|/workflow|list all workflows|
|GET|/workflow/consume|consume a workflow from queue and generate a CSV file with workflow.Data|

## Technologies

- Go, C, C++, Python, Java or any other that you know
- PostgreSQL
- A message queue that you choose, but describe why you choose.

## Solução - Lucas Fernandes de Oliveira

Primeiro criei uma instância do Postgres no [API Elephant SQL](https://www.elephantsql.com/) e criei o Enum de Status e tabela de Workflows.

```sql
    CREATE TYPE Status AS ENUM ('inserted', 'consumed');
    CREATE TABLE workflows (
    	uuid UUID DEFAULT uuid_generate_v4(),
    	status Status DEFAULT 'inserted' NOT NULL,
    	data jsonb NOT NULL,
    	steps text[] NOT NULL,
    	PRIMARY KEY (uuid)
    }
````

Um arquivo ```.env``` foi criado contendo as credenciais para acesso do BD, sendo assim todos podem testar se quiserem.

Depois criei a estrutura do workflow no arquivo [models.go](https://github.com/LucasFOliveira/backend-test/blob/master/go-postgres/models/models.go), onde eu criei uma estrutura básica pro workflow contendo todos os atributos para quando um workflow for recuperado do BD. Criei o Enum de **Status** com os atributos *inserted* e *consumed*. Criei uma estrutura auxiliar apenas com os atributos **Data** e **Steps** para a inserção de um novo workflow no BD, pois o **Status** é automaticamente definido como *inserted* e o **UUID** é gerado pelo próprio SQL. Criei a estrutura do **Data** para a conversão do JSON em CSV ser mais prática e uma estrutura de resposta para os métodos.

Eu decidi usar um *synchronized queue* para o sistema de mensagens, pois foi mais prático para o momento e serviu ao propósito do desafio. Eu imaginei um cenário onde várias pessoas poderiam acessar a *queue*, então decidi torná-la *synchronized* para não haver conflitos. No entanto poderia ser implementado uma solução mais sofisticada com um volume no Docker rodando o RabbitMQ.

Quanto aos *endpoints*, vou explicar cada um aqui:

Para rodar o programa, basta ir na pasta ```go-postgres``` e rodar o comando ```go run main.go```. Todos os testes foram realizados no Postman.

### POST /workflow

Adiciona um novo workflow ao BD da seguinte forma.

```POST http://localhost:8080/workflow```

Na seção Body > raw > JSON, você cola o JSON com os dados que se deseja inser no BD.

![Postman](https://github.com/LucasFOliveira/backend-test/blob/master/go-postgres/images/postman.png?raw=true)

```json
[{
"data": {"name": "Shutdown",
		 "description": "Shutdown the computer"},
"steps": ["Press Alt + F4","Select turn off computer","Press ok"]
}]
```

```json
[{
"data": {"name": "Close window"
		 "description": "Close any window of the computer"},
"steps": ["Go to specified window","Click in the X in the top right"]
}]
```

Dentro do BD, é gerado automaticamente um **UUID**, por default o **Status** padrão é *inserted* e o método retorna um JSON com os dados que foram inseridos.

```json
[{
"UUID":"e322595b-69ba-4b03-bd2d-3fa588b72c2e",
"status":"inserted",
"data":{"name":"Shutdown","description":"Shutdown the computer"},
"steps":["Press Alt + F4","Select turn off computer","Press ok"]
}]
```

![InsertWorkflow](https://github.com/LucasFOliveira/backend-test/blob/master/go-postgres/images/insert.png?raw=true)

Além disso a *queue* de mensagens é incrementada com o **UUID** do elemento que foi adiciona, para fins de ser consumidos posteriormente.

### GET /workflow

Lista todos os workflows sejam eles *inserted* ou *consumed*.

```GET http://localhost:8080/workflow```

Dentro do BD, é o método pega todos os workflows e retorna um JSON os dados de cada um.

```json
[{
"UUID":"e322595b-69ba-4b03-bd2d-3fa588b72c2e",
"status":"inserted",
"data":{"name":"Shutdown","description":"Shutdown the computer"},
"steps":["Press Alt + F4","Select turn off computer","Press ok"]
},{
"UUID":"8b7301ee-a40e-4e5f-9437-d1d629deb987",
"status":"inserted",
"data":{"name":"Close window","description":"Close any window of the computer"},
"steps":["Go to specified window","Click in the X in the top right"]
}]
```
![GetAllWorkflows](https://github.com/LucasFOliveira/backend-test/blob/master/go-postgres/images/get.png?raw=true)

### GET /workflow/consume

Consome o workflow fazendo um *dequeue* na *queue* de mensagens (removendo o elemento na cabeça da fila), muda o **Status** para *consumed* do workflow no BD sem necessariamente excluir o workflow e exporta um CSV com os dados do atributo **Data** para a pasta output, onde o nome do arquivo é o **UUID** do workflow.

```GET http://localhost:8080/workflow```

Dentro do CSV os dados ficam organizados dessa forma:

```
name,description
Shutdown,Shutdown the computer
```

![ConsumeWorkflow](https://github.com/LucasFOliveira/backend-test/blob/master/go-postgres/images/consumed.png?raw=true)

### PATCH /workflow/{UUID}

Esse método eu entendi que seria para atualizar um **Status** de um workflow consumido de *consumed* para *inserted* novamente, então o meu método faz isso e readiciona o **UUID** do workflow ao final da *queue* de mensagens novamente.

```GET http://localhost:8080/workflow/e322595b-69ba-4b03-bd2d-3fa588b72c2e```

O método retorna um JSON com uma mensagem de operação bem sucedida.

```json
{
"id":"e322595b-69ba-4b03-bd2d-3fa588b72c2e",
"message":"Workflow updated successfully, readding to the queue. Total rows/record affected 1\n"
}
```

### DELETE /workflow/{UUID}

Criei uma função para deletar workflow apenas para zerar o banco de dados mesmo, para re-adicionar workflows e fazer mais testes.

```http://localhost:8080/workflow/e322595b-69ba-4b03-bd2d-3fa588b72c2e```

```json
{
"id":"e322595b-69ba-4b03-bd2d-3fa588b72c2e",
"message":"Workflow deleted successfully. Total rows/record affected 1\n"
}
```
Coloquei pequenas impressões no *bash* para manter o rastro de todas a execução dos métodos e se o BD e a *queue* estavam sendo atualizadas corretamente.

![Bash](https://github.com/LucasFOliveira/backend-test/blob/master/go-postgres/images/sh%20exec.png?raw=true)

Gostei muito de trabalhar com Go novamente, aprendi muito e tive uma ótima experiência desenvolvendo esse desafio, foi uma solução simples, passível de várias melhorias, mas espero que gostem e que atenda as expectativas de vocês, obrigado pela oportunidade, um abraço à todos, fiquem seguros e isolados, e usem máscara rsrs.

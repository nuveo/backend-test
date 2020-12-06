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


## Solução Matheus Lima Tavares

Utilizando boas práticas de desenvolvimento criei uma API Rest como especificado, utilizei o framework SpringBoot, com relação ao mensageiro optei pelo RabbitMQ pois queria saber como seria esse novo desafio, posso dizer que para abordagem desse projeto seria dificil ultrapassar mais de 20 mil mensagens por segundo, como teriamos em sistemas de pagamentos por exemplo, também não houve a necessidade de usar multiplos consumer em paralelo, vale destacar também que em um cenário real é comum separarmos o producer do consumer na hora de fazer integração entre apis e microserviços.

## Configuração Inicial:
<p>Você devera configurar o banco de dados postgresql para isso instale o postgresql e os seus drivers, no application.properties e na classe Connect coloque por favor suas credências, não esqueça de criar o banco via interface do postgres, caso não mude as credências da classe Connect não teremos o funcionamento do Patch via JDBC, quis fazer diferente para ilustar.

Instale o postman ou insomnia para fazer testes na api, nesse projeto eu estarei usando o postman.

Siga o passo a passo para a instalação do RabbitMQ no site do https://www.rabbitmq.com/, caso prefire pode utilizar o docker para fazer uso do Rabbit.

```
docker run -d --hostname rabbitserver --name rabbitmq-server -p 15672:15672 -p 5672:5672 rabbitmq:3-management
````
<h6>
DashBoard RabbitMQ: http://127.0.0.1:15672/

User: guest
Pass: guest

## Utilizando Post /workflow
Uma vez iniciado a classe ApiworkflowApplication, abra o postman e faça sua requisição.

http://localhost:8080/workflow

<h6>Na tela do postman clique em Body depois em raw e depois em JSON, depois insira o seguinte exemplo:


```json

    {
        "uuid": null,
        "data": "{\"Project\":\"PixPayment\",\"company\":\"red\"}",
        "workflowSteps": [
            "Propose idea",
            "Create issues",
            "Implement issues",
            "Deploy the project",
            "Track production"
        ],
        "wfs": "INSERTED"
    }
````

Nesse exemplo estou passando o uuid null, porem o jpa ira persistir um uuid aleatório, como mostra a imagem.

![Alt text](relative/src/main/resources/templates/image/img.jpg?raw=true "Title")

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

## Solução - Caio Arruda

Para realizar o teste utilizei as seguintes tecnologias:
- Golang,PostgreSQL e Docker.

- Para realizar o sistema de fila de mensagens tomei a liberdade de implementar a estrutura de dados Fila para resolver o problema. Cheguei a começar a fazer com RabbitMQ mas estava ficando muito preso em alguns probleminhas e por isso a decisão de implementar a fila na mão. 

 ### Para rodar na sua máquina basta ter o docker instalado e faça: 
	 docker-compose up

### POST /workflow
![POST](https://i.ibb.co/ydkhrZ6/create-workflow.png)

### PATCH /workflow
![PATCH](https://i.ibb.co/ZSnKXQ0/patch-workflow.png)

### GET /workflow
![GETALL](https://i.ibb.co/xSxznbd/get-workflow.png)

### GET /workflow/consume
![CONSUME](https://i.ibb.co/pX86vxm/consume-workflow.png)

Ao consumir um workflow é gerado um arquivo .csv na pasta spreadsheet, veja:
![CONSUME2](https://i.ibb.co/5MCdDC5/consumeeee.png)


### Autocrítica do meu código
Algumas escolhas que eu teria feito caso fizesse o teste do zero novamente:
- Usar clean architecture pois o código atual possui uma alta dependência de bibliotecas externas. 
- Quebrar o código em mais funções pequenas que façam apenas uma determinada tarefa, como não tenho muita experiência em Go acabei ficando preso em outros problemas e não pude realizar isto.  


### Agradecimentos
Primeiramente gostaria de agradecer por poder participar do teste para a vaga de back-end e dizer que aprendi demais ao realizar esse teste, através dele pude compreender os problemas que a empresa busca resolver. Só tinha feito coisas básicas em Go (minha proficiência é em Java e Javascript) e gostei demais da linguagem, tanto é que continuarei estudando e amei descobrir que pessoas da empresa e que já passaram por ela fazem parte de um imenso grupo de estudos de go (go-br).  Creio que a nuveo está de parabéns, pois o seu RH é maravilhoso e o teste é um ótimo desafio para quem quer ter noção do que lhe espera.  

Fico no aguardo de uma resposta, valeu!
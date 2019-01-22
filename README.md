# Backend Test

Develop the workflow's REST API following the specification bellow and document it.

## Delivery instructions

Fork this project in a private project and create a branch. When you want to our review, create a PR and put any information that you think is important. Consider we follow your instructions to run your code and look the outcome.

## Defining a workflow

|Name|Type|Description|
|-|-|-|
|UUID|UUID|workflow unique indentifier|
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

- Go, C, C++ or Python
- PostgreSQL
- A message queue that you choose, but describe why you choose.
- [pREST](http://postgres.rest) to comunicate with database. It is a bonus.

## Resolução

Foi buscado separar as responsabilidades ao máximo em serviços, para que pudessem ser facilmente escalados se necessário.

Utilizando docker para containerizar cada serviço pertencente ao micro-sistema do problema. Pôde-se utilizar tecnologias diferentes de forma mais livre.

As tecnologias utilizadas foram:

* Flask -> Para a API. Referente a camada http de apresentação.
* Rabbitmq -> Message queue escolhido por sua alta variedade de soluções e fácil escalabilidade de seus serviços consumidores (consumers).
* pREST -> Para migrações e interação (CRUD) com o banco de dados.
* Postgresql -> Banco de dados amigável ao programador, possui triggers que notificam em um canal específico a qualquer inserção na tabela **workflow**.

## Setup
Para preparar o ambiente, executar os seguintes comandos do Makefile:
        
        make prepare
    
Estes a seguir, startam os containers principais e por fim, os workers:

        make run
Após as aplicações principais estiverem iniciado em seus containers, execute:

        make run-workers

        
Separei explicitamente em dois comandos, pois a execução dos containers geralmente não reflete a de suas aplicações internas,
este comportamente comumente acaba abortando a conexao dos workers (por iniciarem mais rapidamente), encerrando seus serviços de maneira prematura. Tentei a flag `depends_on` dentro suas sessões no `docker-compose.yml`, mas não solucionei este detalhe.

Para efetuar pequenos testes, execute:

        make test

É interessante alterar as configurações em specs/prest/prest.toml para uma base de testes antes de startar os containers e efetuar o comando acima. Porém, ao executa-la uma vez, verá o comportamento do fluxo de responsabilidades entre os containers. 

Que é o seguinte:

A API ao receber um post de um `workflow`, envia o payload para o pREST com os dados necessários para a requisição, este, salva na tabela `workflow`, dentro do POSTGRES, uma trigger dispara um NOTIFY, (a trigger dispara após cada inserção de um registro), onde um PRODUCER ouvindo este canal, insere uma task na fila do RABBITMQ. Neste ponto, há três CONSUMERS conectados ao serviço do message queue, processando e transformando o conteúdo em CSV, e salvando o resultado do processamento utilizando novamente o pREST, na tabela de cache `cache_workflow`.
O retorno dos dados csv seria feito portanto através da tabela de cache.

## Documentação

Há uma breve documentação escrita em .yml baseando-se no Swagger. Ela pode ser acessada através do endpoint **/apidocs**.

Lá, será possivel efetuar todas as requisições entre as rotas pertencentes a API, e abaixo delas, um exemplo de modelo utilizado para enviar à api.


## Limitações
Algumas ressalvas:
* O serializador json_to_csv criado para o desafio, não transforma `json objects` dentro de `json arrays` (quando o fiz suportar, obtive problemas em outros casos de dados). No mais o inverso é permitido, assim como objetos aninhados e multiplos filhos (extraídos via estratégia recursiva).
* O /workflow/consume não pega uma queue para processar o csv em tempo real, pois preferi por engarregar esta tarefa aos workers, neste caso, a rota consulta ao pREST pelo ultimo registro salvo na tabela de cache `cache_workflow` e o retorna.

Busquei nos limites de meu conhecimento, oferecer um trabalho satisfatório, um forte abraço!

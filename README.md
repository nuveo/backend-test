# Teste: Backend

A Nuveo é uma empresa deep-tech e por trabalhar com tecnologias de fronteira precisa que seus colaboradores sejam curiosos e antenados com as melhores práticas de mercado. Nesse sentido, estruturamos trẽs níveis de teste para avaliar conhecimentos do escopo de Backend Developer. Escolha o nível em que seja mais aderente aos seus conhecimentos.

_*Observação 1:*_ _O exemplo a ser implementado é o mesmo, independente do nível de expertise que está sendo avaliado. O que diferencia  um nível do outro é a completude de tarefas e raciocínio aplicado. Portanto, seja o mais abrangente e eficiente possível._

_*Observação 2:*_ _Ao final da implementação, por favor inclua os seguintes perfis do github para acessar a sua solução: [KlebersonCanuto](https://github.com/KlebersonCanuto) e [diegogonzalesdesouza](https://github.com/diegogonzalesdesouza)._

## Desafio para Backend N1

Desenvolva uma API REST usando Golang com as seguintes características.

 - Use um repositório privado no GitHub para fazer seu teste, entre em contato para saber com quem você deve compartilhar seu repositório.
 - Dada a seguinte estrutura de dados para conter clientes:

```
	"uuid" string contendo UUID
	"nome" string contendo nome do cliente
	"endereco" string contendo endereço do cliente
	"cadastrado_em" string contendo time stamp com data e hora do cadastro (ISO 8601 sem time zone)
	"atualizado_em" string contendo time stamp com data e hora da atualização (ISO 8601 sem time zone)
```

 - Crie uma estrutura no PostgreSQL para conter os dados de clientes, essa estrutura deve usar um campo UUID como chave primária.
 - Crie os seguintes endpoints para manipular a estrutura anterior. 

POST  /cliente (vai servir para cadastrar o cliente, deve retornar o UUID do cliente cadastrado)

GET    /cliente/{UUID} (vai servir para recuperar um JSON com o cliente cadastrado.

PUT    /cliente/{UUID} (vai servir para alterar um cliente cadastrado)

DELETE /cliente/{UUID} (vai servir para remover um cliente da base)

GET /cliente (vai ser usado para listar todos os clientes)

 - Lembre de testar o seu código com testes unitários, esperamos pelo menos 75% de cobertura de testes.
 - A stack deve rodar em containers Docker e idealmente ser carregada com docker-compose.
 - Não esqueça de documentar seu código.

## Desafio para Backend N2

Desenvolva uma API REST usando Golang com as seguintes características.

 - Use um repositório privado no GitHub para fazer seu teste, entre em contato para saber com quem você deve compartilhar seu repositório.
 - Dada a seguinte estrutura de dados para conter clientes:

```
	"uuid" string contendo UUID
	"nome" string contendo nome do cliente
	"endereco" string contendo endereço do cliente
	"cadastrado_em" string contendo time stamp com data e hora do cadastro (ISO 8601 sem time zone)
	"atualizado_em" string contendo time stamp com data e hora da atualização (ISO 8601 sem time zone)
```

 - Crie uma estrutura no PostgreSQL para conter os dados de clientes, essa estrutura deve usar um campo UUID como chave primária. Para carregar a estrutura no banco use a ferramenta de migrations do GoSidekick https://github.com/gosidekick/migration
 - Crie os seguintes endpoints para manipular a estrutura anterior. 

POST  /cliente (vai servir para cadastrar o cliente, deve retornar o UUID do cliente cadastrado)

GET    /cliente/{UUID} (vai servir para recuperar um JSON com o cliente cadastrado.

PUT    /cliente/{UUID} (vai servir para alterar um cliente cadastrado)

DELETE /cliente/{UUID} (vai servir para remover um cliente da base)

GET /cliente (vai ser usado para listar todos os clientes)

 - Crie um mecanismo para colocar em uma fila RabbitMQ os dados do último cliente cadastrado. (você pode usar outros tipos de fila se preferir, desde que ela suba isoladamente em um container).
 - Crie um micro-serviço que coleta a última entrada da fila e salva em um arquivo JSON em um diretório configurado pela variável de ambiente NOVOS_CLIENTES. Sendo um arquivo para cada novo cliente.
 - Lembre de testar o seu código com testes unitários, esperamos pelo menos 80% de cobertura de testes.
 - A stack deve rodar em containers Docker e idealmente ser carregada com docker-compose.
Não esqueça de documentar seu código.

## Desafio para Backend N3

Desenvolva uma API REST usando Golang com as seguintes características.

 - Use um repositório privado no GitHub para fazer seu teste, entre em contato para saber com quem você deve compartilhar seu repositório.
 - Dada a seguinte estrutura de dados para conter clientes:

```
	"uuid" string contendo UUID
	"nome" string contendo nome do cliente
	"endereco" string contendo endereço do cliente
	"cadastrado_em" string contendo time stamp com data e hora do cadastro (ISO 8601 sem time zone)
	"atualizado_em" string contendo time stamp com data e hora da atualização (ISO 8601 sem time zone)
```

 - Crie uma estrutura no PostgreSQL para conter os dados de clientes, essa estrutura deve usar um campo UUID como chave primária. Para carregar a estrutura no banco use a ferramenta de migrations do GoSidekick https://github.com/gosidekick/migration

 - Crie os seguintes endpoints para manipular a estrutura anterior. 
 
POST  /cliente (vai servir para cadastrar o cliente, deve retornar o UUID do cliente cadastrado)

GET    /cliente/{UUID} (vai servir para recuperar um JSON com o cliente cadastrado.

PUT    /cliente/{UUID} (vai servir para alterar um cliente cadastrado)

DELETE /cliente/{UUID} (vai servir para remover um cliente da base)

GET /cliente (vai ser usado para listar todos os clientes)

 - Crie um mecanismo para colocar em uma fila RabbitMQ os dados do último cliente cadastrado. (você pode usar outros tipos de fila se preferir, desde que ela suba isoladamente em um container).
 - Crie um micro-serviço que coleta a última entrada da fila e salva em um arquivo JSON em um diretório configurado pela variável de ambiente NOVOS_CLIENTES. Sendo um arquivo para cada novo cliente.
 - A API não deve ficar exposta, é necessário alguma forma de autenticação de usuário para permitir o uso dos endpoints. Você pode escolher qualquer forma de autenticação que considere adequada, nós sugerimos usar JWT.
 - Lembre de testar o seu código com testes unitários, esperamos pelo menos 85% de cobertura de testes e que o acesso ao banco de dados possa ser simulado sem a necessidade de subir o banco de dados. (mock)
 - Crie testes end-to-end além dos testes unitários para validar o sistema.
 - A stack deve rodar em containers Docker e idealmente ser carregada com docker-compose e/ou kubernetes.
 - Não esqueça de documentar seu código.

# Tecnologias Requisitadas

Escolha a que for mais apta para os seus conhecimentos:

- Go, C, C++, Python, Java or any other that you know
- PostgreSQL
- Um sistema gerenciados de mensagens a sua escolha

# PULL REQUEST - COMENTÁRIOS

Procurei codificar o teste levando em consideração as boas práticas de desenvolvimento, como a distinção de responsabilidade em camadas.

A aplicação foi desenvolvida como uma aplicação Spring-boot, arquitetura java recomendada para microserviços.

Algumas funcionalidades ficaram de fora por absoluta falta de tempo (mas essenciais em produtos finalizados), como:

- Logs.
- Documentação e uso mais extensivo de comentários (aqui faço um mea-culpa. Geralmente costumo comentar MUITO meus códigos, vide projetos meus como este por exemplo 
 [https://github.com/rsouza01/tovsolver]. Neste teste fiquei devendo).
- Testes unitários e mapeamento mais consistente das exceções com mensagens de erro mais específicas.
- Uma maneira mais segura de guardar as credenciais da conta AWS (como variáveis de ambiente por exemplo) ou utilização de IAM. 
- Mapeamento do campo Workflow.data para o tipo correto JSONB ao invés do campo String, utilizando o UserType do Hibernate.

Foi escolhida o message queue como o AWS-SQS por questões de facilidade de implementação e pelo fato de eu já possuir uma conta AWS para desenvolvimento (idem para o provedor do PostgreSQL).

Segue no projeto um shellscript que faz as vezes do script do Postman (callREST.sh). É necessario a instalação do aplicativo cURL para utilização.

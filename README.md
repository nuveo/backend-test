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

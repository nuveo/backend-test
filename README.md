# Backend Test

Develop the workflow's REST API following the specification bellow and document it.

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
|GET|/workflow/consume|consume a workflow from queue and generete a CSV file with workflow.Data|

## Technologies

- Go, C or C++
- PostgreSQL
- A message queue that you choose, but discribe why you choose.
- [pREST](http://postgres.rest) to comunicate with database. It is a bonus.

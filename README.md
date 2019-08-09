# Backend Test

Develop the workflow's REST API following the specification bellow and document it.

## Delivery instructions

Clone this project and push a private repository in the [Gitlab](https://about.gitlab.com/) or [Bitbucket](https://bitbucket.org/). When you want to our review, write any information that you think important in the README.md and send an email to talentos@nuveo.ai. We'll follow your instructions to run your code and look the outcome. 

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

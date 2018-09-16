# Nuveo's Backend Test
Solution for Nuveo's backend test developed by Lauren Maria Ferreira.

## Test Description
Develop the workflow's REST API following the specification bellow and document it.

## Workflow

|Name|Type|Description|
|-|-|-|
|UUID|UUID|workflow unique indentifier|
|status|Enum(inserted, consumed)|workflow status|
|data|JSONB|workflow input|
|steps|Array|name of workflow steps

## Prerequisites

- Go >=1.10

Using version 1.10.1

```
go version go1.10.1 linux/amd64
```

- PostgreSQL >= 10

Using version 10.5

```
PostgreSQL 10.5 (Ubuntu 10.5-0ubuntu0.18.04) on x86_64-pc-linux-gnu, compiled by gcc (Ubuntu 7.3.0-16ubuntu3) 7.3.0, 64-bit
```

## Installation

- Go

```sh
go get github.com/gorilla/mux github.com/lib/pq
```

- PostgreSQL

```sh
createdb <dbname>
```

Example:
```sh
createdb nuveo
```

## Running

Inside project's folder:
```sh
go build
```

```sh
./backend-test <dbuser> <dbpassword> <dbname>
```

Example:
```sh
./backend-test postgres 1234 nuveo
```

## Endpoints

|Verb|URL|Description|
|-|-|-|
|POST|/workflows|insert a workflow on database and on queue and respond request with the inserted workflow|
|PATCH|/workflows/{UUID}|update status from specific workflow|
|GET|/workflows|list all workflows|
|GET|/workflows/consume|consume a workflow from queue and generate a CSV file with workflow.Data|

### POST /workflows

* **Data Params**

Workflow data: string [required]

Steps: string array [required]

* **Success Response:**
  * **Code:** 200

    **Content:**

    ```json
    {
        "UUID": "639d0205-c72f-414c-bc4c-45ea85c40404",
        "Status": "inserted",
        "Data": "\"{teste: Teste}\"",
        "Steps": [
            "go",
            "eh",
            "vida"
        ]
    }
    ```
 
* **Error Response:**

  * **Code:** 500 Internal Error, happens because you've missed some argument

    **Content:**

    ```json
    {
        "error": "Failed to insert workflow: pq: null value in column \"steps\" violates not-null constraint"
    }
    ```

### PATCH /workflows

* **Data Params**

Workflow UUID: string [required]

* **Success Response:**
  * **Code:** 200

    **Content:**

    ```json
    {
        "UUID": "669be55e-1b40-488b-a7c3-30e23f309580",
        "Status": "consumed",
        "Data": "\"{teste: Teste}\"",
        "Steps": [
            "go",
            "eh",
            "vida"
        ]
    }
    ```
 
* **Error Response:**

  * **Code:** 500 Internal Error, happens because you've tried to update an already updated workflow

    **Content:**

    ```json
    {
        "error": "Workflow already consumed"
    }
    ```

   * **Code:** 500 Internal Error, happens because you've inserted an workflow that wasn't inserted yet

       **Content:**
       ```json
       {
            "error": "Workflow not found"
       }
       ```

    * **Code:** 500 Internal Error, happens because you've inserted an invalid workflow UUID

       **Content:**
       ```json
        {
            "error": "Failed to get workflow: pq: invalid input syntax for type uuid: \"5\""
        }
       ```


### GET /workflows

* **Success Response:**

  * **Code:** 200

    **Content:**

    ```json
    [
        {
            "UUID": "9df72d4b-b659-44c5-b8dc-5449959410be",
            "Status": "inserted",
            "Data": "\"{'teste': 'Teste'}\"",
            "Steps": [
                "go",
                "eh",
                "vida"
            ]
        },
        {
            "UUID": "b2c89cfe-eec6-426a-8001-fd4c93ad1f33",
            "Status": "inserted",
            "Data": "\"{'teste': 'Teste'}\"",
            "Steps": [
                "go",
                "eh",
                "vida"
            ]
        },
        {
            "UUID": "4a49e542-6cdf-4457-a784-01305c6b9023",
            "Status": "consumed",
            "Data": "\"{'teste': 'Teste'}\"",
            "Steps": [
                "go",
                "eh",
                "vida"
            ]
        }
    ]
    ```
 
* **Error Response:**

  * **Code:** 500 Internal Error, happens because you've missed some argument

    **Content:**

    ```json
    {
        "error": "Failed to insert workflow: pq: null value in column \"steps\" violates not-null constraint"
    }
    ```       

### GET /workflows/consume

* **Success Response:**

  * **Code:** 200
 
* **Error Response:**

  * **Code:** 500 Internal Error, happens because queue is empty and there's no way to generated a CSV file

    **Content:**

    ```json
    {
        "error": "Empty queue"
    }
    ```       

## Feedback
It was a fun test to work for in the past few days. I've learned a lot about PostgreSQL, testing and many other things while developing it.

My main goal was to keep this project as simple as possible, including its code, its file structure, its architecture and its depencencies. I've only used two dependencies: Gorilla Mux, to deal with handler APIs and routes, and the PostgreSQL driver inside Go.

To keep things simple, I decided not to use a complex message queue service and I've only used a queue data structure.

I decided not to use pREST.

I've changed endpoint's name to workflowS because of a convention inside REST API community.

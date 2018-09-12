#!/bin/bash

echo -e '\n\n\nGreetings'
curl -H "Accept:application/json" http://localhost:8080/api/workflow/greetings

echo -e '\n\n\nCreate with UUID'
curl -i -X POST -H "Content-Type:application/json" http://localhost:8080/api/workflow/ -d @json/createUUID.json

echo -e '\n\n\nCreate without UUID'
curl -i -X POST -H "Content-Type:application/json" http://localhost:8080/api/workflow/ -d @json/create.json

echo -e '\n\n\nGet all'
curl -i -X GET -H "Content-Type:application/json" http://localhost:8080/api/workflows/

echo -e '\n\n\nUpdate with UUID'
curl -i -X PATCH -H "Content-Type:application/json" http://localhost:8080/api/workflow/d632e0cf-1702-428c-bde6-b77f4949098a -d @json/updateUUID.json

echo -e '\n\n\nGet'
curl -i -X GET -H "Content-Type:application/json" http://localhost:8080/api/workflow/d632e0cf-1702-428c-bde6-b77f4949098a

echo -e '\n\n\nConsume'
curl -i -X GET -H "Content-Type:application/json" http://localhost:8080/api/workflow/consume


echo
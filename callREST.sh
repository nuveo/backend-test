#!/bin/bash

#echo 'Greetings'
#curl -H "Accept:application/json" http://localhost:8080/api/workflow/greetings

#echo 'Create with UUID'
#curl -i -X POST -H "Content-Type:application/json" http://localhost:8080/api/workflow/ -d @json/createUUID.json

#echo 'Create without UUID'
#curl -i -X POST -H "Content-Type:application/json" http://localhost:8080/api/workflow/ -d @json/create.json

#echo 'Get all'
#curl -i -X GET -H "Content-Type:application/json" http://localhost:8080/api/workflows/


#echo 'Update with UUID'
#curl -i -X PATCH -H "Content-Type:application/json" http://localhost:8080/api/workflow/d632e0cf-1702-428c-bde6-b77f4949098a -d @json/updateUUID.json

#echo 'Get'
#curl -i -X GET -H "Content-Type:application/json" http://localhost:8080/api/workflow/d632e0cf-1702-428c-bde6-b77f4949098a

echo 'Consume'
curl -i -X GET -H "Content-Type:application/json" http://localhost:8080/api/workflow/consume


echo
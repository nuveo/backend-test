"""
    API para Workflow
"""
import os
import uuid
import requests
from flask import (
    Flask, 
    request,
    abort,
    jsonify
)
from flasgger import (
    swag_from,
    Swagger
)


PREST_INFO = {
    'host': os.getenv('API_PREST_HOST'),
    'port': os.getenv('API_PREST_PORT'),
    'db': os.getenv('API_PREST_DB_NAME'),
    'schema': os.getenv('API_PREST_DB_SCHEMA')
}

BASE_URL = 'http://{host}:{port}/{db}/{schema}'.format(**PREST_INFO)

app = Flask(__name__)
swagger = Swagger(app)


@app.route('/')
@swag_from('docs/index.yml')
def index():
    """
        Rota principal
    """
    return 'Hello World'


@app.route('/workflow/', methods=['GET', 'POST'], endpoint='workflow')
@swag_from('docs/workflow_get.yml', endpoint='workflow', methods=['GET'])
@swag_from('docs/workflow_post.yml', endpoint='workflow', methods=['POST'])
def workflow():
    """
        Get all or insert one workflow on database using prest
    """
    def _get():
        return requests.get(f'{BASE_URL}/workflow').content

    def _post():
        if not request.is_json:
            return abort(400, 'Body message is not a valid json')

        body = request.get_json()
        status = body['status'] if body['status'] in ['inserted', 'consumed'] else None

        if not status:
            return abort(403, 'status unsupported')

        data = {
            'UUID': str(uuid.uuid4()),
            'status': body.status,
            'data': body.data,
            'steps': body.steps,
        }
        requests.post(f'{BASE_URL}/workflow', json=data)
        return jsonify(data)

    _workflow = {
        'GET': _get,
        'POST': _post
    }

    return _workflow[request.method]()

@app.route('/workflow/<uuid>', methods=['GET', 'PATCH'], endpoint='workflow_uuid')
@swag_from('docs/workflow_uuid_get.yml', endpoint='workflow_uuid', methods=['GET'])
@swag_from('docs/workflow_uuid_patch.yml', endpoint='workflow_uuid', methods=['PATCH'])
def workflow_uuid(uuid=None):
    """
        Manage specific workflows based on UUID passed by parameter
    """
    return 'Workflow returned with uuid:' + uuid

@app.route('/workflow/consume/', endpoint='workflow_consume', methods=['GET'])
@swag_from('docs/workflow_consume_get.yml', endpoint='workflow_consume', methods=['GET'])
def consume():
    """
        Consome a workflow from queue and return your .csv from Workflow.data
    """
    return 'Workflow consumed by queue'

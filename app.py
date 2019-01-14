"""
    API para Workflow
"""
from flask import Flask
from flasgger import swag_from, Swagger


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
        Get all or insert one workflow on database and queue
    """
    return 'Single Workflow returned'

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

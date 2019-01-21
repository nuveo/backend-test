import os
from os.path import join, dirname
import json

import requests
import pika
from dotenv import load_dotenv

import json_to_csv

load_dotenv(join(dirname(__file__), '.env'))

RABBITMQ_HOST = os.getenv('RBMQ_HOST')

PREST_INFO = {
    'host': os.getenv('PREST_HOST'),
    'db': os.getenv('PREST_DATABASE'),
    'schema': os.getenv('PREST_SCHEMA'),
    'port': os.getenv('PREST_PORT')
}

WEB_INFO = {
    'host': os.getenv('WEB_HOST'),
    'port': os.getenv('WEB_PORT')
}

WEB_URL = 'http://{host}:{port}'.format(**WEB_INFO)
PREST_URL = 'http://{host}:{port}/{db}/{schema}'.format(**PREST_INFO)

connection = pika.BlockingConnection(pika.ConnectionParameters(RABBITMQ_HOST))
channel = connection.channel()

channel.queue_declare(queue='workflow_queue')


def callback(ch, method, properties, body):
    data = json.loads(body)
    csv = json_to_csv.to_csv(data['data'])
    post_body = {
        'uuid': data['uuid'],
        'data': json.dumps(data['data']),
        'csv': csv
    }

    requests.post(f'{PREST_URL}/cache_workflow', json=post_body)
    requests.patch(f'{WEB_URL}/workflow/'+data['uuid'], json={"status": "consumed"})


channel.basic_consume(callback, queue='workflow_queue', no_ack=True)
channel.start_consuming()

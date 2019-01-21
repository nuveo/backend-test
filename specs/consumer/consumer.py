import os
from os.path import join, dirname
import pika
from dotenv import load_dotenv
import json
import json_to_csv
import requests

load_dotenv(join(dirname(__file__), '.env'))

RABBITMQ_HOST = os.getenv('RBMQ_HOST')

PREST_INFO = {
    'host': os.getenv('PREST_HOST'),
    'db': os.getenv('PREST_DATABASE'),
    'schema': os.getenv('PREST_SCHEMA'),
    'port': os.getenv('PREST_PORT')
}

BASE_URL = 'http://{host}:{port}/{db}/{schema}'.format(**PREST_INFO)

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
    print(post_body)
    requests.post(f'{BASE_URL}/cache_workflow', json=post_body)



channel.basic_consume(callback, queue='workflow_queue', no_ack=True)
# print(' [*] Waiting for messages. To exit press CTRL+C')

channel.start_consuming()

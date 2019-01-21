import os
from os.path import join, dirname
import pika
from dotenv import load_dotenv

load_dotenv(join(dirname(__file__), '.env'))

RABBITMQ_HOST = os.getenv('RBMQ_HOST')
CONNECTION_INFO = {
    'host': os.getenv('PG_HOST'),
    'database': os.getenv('PG_DATABASE'),
    'user': os.getenv('PG_USER'),
    'password': os.getenv('PG_PASSWORD'),
    'port': os.getenv('PG_PORT')
}

connection = pika.BlockingConnection(pika.ConnectionParameters(RABBITMQ_HOST))
channel = connection.channel()

channel.queue_declare(queue='workflow_queue')


def callback(ch, method, properties, body):
    print(' [x] Received %r' % body)


channel.basic_consume(callback, queue='workflow_queue', no_ack=True)
print(' [*] Waiting for messages. To exit press CTRL+C')

channel.start_consuming()

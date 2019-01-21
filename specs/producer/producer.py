"""
    Listen the workflow_insert channel and sends
    to the workflow_queue rabbitmq queue
"""
import os
from os.path import join, dirname
import select
import pika
import psycopg2
import psycopg2.extensions
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


def rabbit_notify(message):
    connection = pika.BlockingConnection(pika.ConnectionParameters(RABBITMQ_HOST))
    channel = connection.channel()

    channel.queue_declare(queue='workflow_queue')
    channel.basic_publish(exchange='', routing_key='workflow_queue', body=message)
    print(' [x] Sent ', message)

    connection.close()


conn = psycopg2.connect(**CONNECTION_INFO)
conn.set_isolation_level(psycopg2.extensions.ISOLATION_LEVEL_AUTOCOMMIT)

curs = conn.cursor()
curs.execute('LISTEN workflow_insert;')

print('Esperando notificações no canal \'workflow_insert\'')
while 1:
    if select.select([conn], [], [], 5) == ([], [], []):
        print('Timeout')
    else:
        conn.poll()
        while conn.notifies:
            notify = conn.notifies.pop(0)
            rabbit_notify(notify.payload)
            print('Obteve NOTIFY:', notify.pid, notify.channel, notify.payload)

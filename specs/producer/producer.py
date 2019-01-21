"""
    Listen the workflow_insert channel and sends
    to the workflow_queue rabbitmq queue
"""
import os
import select
import pika
import psycopg2
import psycopg2.extensions


def rabbit_notify(message):
    connection = pika.BlockingConnection(pika.ConnectionParameters('rbmq'))
    channel = connection.channel()

    channel.queue_declare(queue='workflow_queue')
    channel.basic_publish(exchange='', routing_key='workflow_queue', body=message)
    print(' [x] Sent ', message)

    connection.close()

CONNECTION_INFO = {
    'host': 'db',
    'database': 'postgres',
    'user': 'postgres',
    'password': 'postgres',
    'port': 5432
}

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


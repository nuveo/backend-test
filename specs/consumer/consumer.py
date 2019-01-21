import pika

connection = pika.BlockingConnection(pika.ConnectionParameters('rbmq'))
channel = connection.channel()

channel.queue_declare(queue='workflow_queue')

def callback(ch, method, properties, body):
    print(' [x] Received %r' % body)

channel.basic_consume(callback, queue='workflow_queue', no_ack=True)
print(' [*] Waiting for messages. To exit press CTRL+C')

channel.start_consuming()
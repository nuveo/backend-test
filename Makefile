prepare:
	pipenv install
	pipenv shell

lint:
	pycodestyle . --max-line-length=90

run-containers:
	docker-compose up -d db
	docker-compose up -d prest
	docker-compose up -d rbmq
	docker-compose up -d web
	docker-compose up -d producer

run:
	make run-containers

run-workers:
	docker-compose up -d --scale consumer=3

test:
	docker-compose exec web python /web/tests.py

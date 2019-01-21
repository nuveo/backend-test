prepare:
	pipenv install
	pipenv shell

lint:
	pycodestyle . --max-line-length=90

run-containers:
	docker-compose up -d --scale consumer=3

run:
	make run-containers

test:
	python tests.py

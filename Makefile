prepare:
	pipenv install
	pipenv shell

lint:
	pycodestyle . --max-line-length=90

run-containers:
	docker-compose up -d --scale consumer=3

run-webservice:
	flask run

run:
	make run-containers && make run-webservice

test:
	python tests.py

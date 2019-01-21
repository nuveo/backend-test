prepare:
	pipenv install

lint:
	pycodestyle . --max-line-length=90
	
run:
	flask run
test:
	python tests.py

producer:
	python producer.py
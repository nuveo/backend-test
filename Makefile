prepare:
	pipenv install

lint:
	pycodestyle
	
run:
	flask run
test:
	python tests.py

producer:
	python producer.py
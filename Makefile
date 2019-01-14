prepare:
	pipenv install

lint:
	pycodestyle
	
run:
	flask run

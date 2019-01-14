prepare:
	pipenv install

lint:
	pycodestyle
	
run:
	flask run
test:
	prest migrate workflow_test up

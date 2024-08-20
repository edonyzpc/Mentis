default: mentis

PYTHONCMD := python3.10

mentis:
	go build -o bin/mentis .

install-deps:
	pip install '.[all]'

build-pypi: install-deps
	${PYTHONCMD} -m build

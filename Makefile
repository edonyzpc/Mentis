default: mentis

PYTHONCMD := python3.10

mentis:
	go build -o bin/mentis .

install-deps:
	pip install '.[all]'

build-pypi: install-deps
	${PYTHONCMD} -m build

clean:
	rm -rf ./bin/conversations/*
	rm -rf ./bin/exports/*
	rm -rf ./bin/logging/*

dev-console:
	textual console --port 7342

debug-mentis:
	textual run --dev  mentis.textual_ui.app:run --port 7342
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
# ignore the system and event logs
	textual console -x SYSTEM -x EVENT --port 7342

debug-mentis:
	textual run --dev  mentis.textual_ui.app:run --port 7342
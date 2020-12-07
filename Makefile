.PHONY: build
build:
	docker build -t hirosume/whatmyip .
push:
	docker push hirosume/whatmyip
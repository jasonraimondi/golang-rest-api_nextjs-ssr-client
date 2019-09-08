.PHONY: local test push-ssr push-api build-ssr build-api build push api ssr

local:
	docker-compose -f docker-compose.app.yml -f docker-compose.yml up -d

test:
	drone exec --branch master --event push

push-ssr:
	docker push jasonraimondi/kim-ssr

push-api:
	docker push jasonraimondi/kim-api

build-ssr:
	docker build -t jasonraimondi/kim-ssr ./web/

build-api:
	docker build -t jasonraimondi/kim-api ./

build: build-api build-ssr
push: push-api push-ssr

ssr: build-ssr push-ssr
api: build-api push-api


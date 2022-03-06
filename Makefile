.PHONY: default
default: build

image:
	docker build -t virgocoder/webhooks .

build:
	go build -o webhooks
	chmod +x webhooks

push:
	docker push virgocoder/webhooks
VERSION := v0.0.1

.PHONY: default
default: build

image:
	docker build -t virgocoder/webhooks:$(VERSION) .

build:
	go build -o webhooks
	chmod +x webhooks

push:
	docker push virgocoder/webhooks:$(VERSION)
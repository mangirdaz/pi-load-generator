build: 
	CGO_ENABLED=0 GOOS=linux go build -o api .

build-image:
	docker build -t docker.io/mangirdas/pi-load-generator -f Dockerfile .
publish: build build-image
	docker push docker.io/mangirdas/pi-load-generator

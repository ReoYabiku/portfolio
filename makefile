.PHONY: build-app
build-app:
	docker build --platform linux/amd64 -t integral0515/portfolio-go .
	docker push integral0515/portfolio-go
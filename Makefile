.PHONY: build clean run stop storage curl

build: clean
	cd function && \
		go mod tidy && \
		go mod vendor && \
		GOOS=linux GOARCH=amd64 go build -o ../processorHandler

clean:
	@rm -f processorHandler
	@rm -rf database/*

run: build 
	docker-compose -f docker-compose.yaml up -d --build --remove-orphans

stop:
	docker-compose -f docker-compose.yaml down --remove-orphans

storage:
	docker-compose -f docker-compose.yaml up -d --build --remove-orphans storage
	
curl:
	curl -k http://localhost:8080/api/processor

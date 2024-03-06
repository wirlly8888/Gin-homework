IMAGE_NAME=gin-homework
CONTAINER_NAME=gin-homework

build-image:
	docker build -t $(IMAGE_NAME) .

delete-image:
	docker rmi $(IMAGE_NAME)

run-server:
	docker run -d -p 8080:8080 --name $(CONTAINER_NAME) $(IMAGE_NAME)

stop-server:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)
	

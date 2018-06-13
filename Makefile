all: push

PREFIX=registry.cn-beijing.aliyuncs.com/xuyuntech

IMAGE_APP=usercenter
IMAGE_APP_TAG=latest

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o user
	docker build -t ${PREFIX}/${IMAGE_APP}:${IMAGE_APP_TAG} .

push: build
	docker push ${PREFIX}/${IMAGE_APP}:${IMAGE_APP_TAG}

mysql:
	docker rm -f register-visit-mysql || true
	docker run -d --name register-visit-mysql -e MYSQL_ROOT_PASSWORD=rootpw -e MYSQL_DATABASE=xuyuntech_health -v `pwd`/db_data:/var/lib/mysql -p 3306:3306 mysql:5


##### ENV
env-up:
	@echo "Start environment ..."
	@cd fabric && docker-compose up --force-recreate -d
	@echo "Sleep 15 seconds in order to let the environment setup correctly"
	@sleep 15
	@echo "Environment up"

env-up-pro:
	@echo "Start environment ..."
	@cd fabric && docker-compose -f docker-compose-pro.yml up --force-recreate -d
	@echo "Sleep 15 seconds in order to let the environment setup correctly"
	@sleep 15
	@echo "Environment up"

env-down:
	@echo "Stop environment ..."
	@cd fabric && docker-compose down
	@echo "Environment down"

##### CLEAN
clean: env-down
	@echo "Clean up ..."
	@rm -rf /tmp/heroes-service-* heroes-service
	@docker rm -f -v `docker ps -a --no-trunc | grep "heroes-service" | cut -d ' ' -f 1` 2>/dev/null || true
	@docker rmi `docker images --no-trunc | grep "heroes-service" | cut -d ' ' -f 1` 2>/dev/null || true
	@echo "Clean up done"

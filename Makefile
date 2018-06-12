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
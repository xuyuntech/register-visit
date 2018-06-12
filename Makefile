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
	docker rm -f usercenter-mysql || true
	docker run -d --name usercenter-mysql -e MYSQL_ROOT_PASSWORD=Xuyun.123 -e MYSQL_DATABASE=xuyuntech_health -e MYSQL_USER=xuyuntech -p 3306:3306 -v `pwd`/db_data:/var/lib/mysql mysql
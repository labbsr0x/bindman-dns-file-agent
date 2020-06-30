MAINTAINER=labbsr0x
PROJECT=bindman-dns-file-agent
VERSION=0.0.1


all: build

build:
	@docker build --tag=${MAINTAINER}/${PROJECT}:${VERSION} .
	@docker build --tag=${MAINTAINER}/${PROJECT}:latest .

push:
	@docker push ${MAINTAINER}/${PROJECT}:${VERSION}
	@docker push ${MAINTAINER}/${PROJECT}:latest 

docker-run:
	@docker run --rm --name ${PROJECT} -d --volume ${PWD}/.agent:/data -e BINDMAN_DNS_MANAGER_ADDR=http://bindman-dns-bind9:7070 -e BINDMAN_DNS_REVERSE_PROXY_ADDR=0.0.0.0 -e BINDMAN_AGENT_CONFIG_PATH=/data/bindman_agent.json --network network-bind ${MAINTAINER}/${PROJECT}

docker-stop:
	@docker stop ${PROJECT}

docker-exec:
	@docker exec -it ${PROJECT} sh
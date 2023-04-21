run_nodes:
	python3 assets/node1.py
	python3 assets/node2.py

build:
	#create a network
	# docker network create balancer-network
	# create node image
	docker build \
		-t node-python-image \
		-f api/nodes/Dockerfile.node .
	# create balancer image
	docker build \
		-t balancer-image \
		-f api/Dockerfile.balancer .

remove-containers:
	docker rm -f balancer-container
	docker rm -f node1-container
	docker rm -f node2-container

create-network:
	docker network create balancer-network

run:
	docker run -d --network balancer-network \
		--name balancer-container \
		-p8087:8087 balancer-image 

	docker run -d --network balancer-network \
		--name node1-container \
		-e PORT=5001 -e NAME=Node1 node-python-image

	docker run -d --network balancer-network \
		--name node2-container \
		-e PORT=5002 -e NAME=Node2 node-python-image

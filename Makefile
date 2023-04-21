#---------------------------------------------
#builds
#create balancer image
build-balancer:
	docker build \
		-t balancer-image \
		-f api/Dockerfile.balancer .

# create node image
build-nodes:
	docker build \
		-t node-python-image \
		-f api/nodes/Dockerfile.node .

#create network
create-network:
	#create a network
	docker network create balancer-network

#build all
build:
	create-network
	build-balancer
	build-nodes

#----------------------------------------------
#remove
remove-balancer:
	docker rm -f balancer-container

remove-balancer-image:
	docker rmi -f balancer-image

remove-nodes:
	docker rm -f node1-container node2-container

remove-nodes-image:
	docker rmi -f node-python-image

remove:
	remove-balancer
	remove-nodes
	remove-balancer-image
	remove-nodes-image
#---------------------------------------------
#run
run-nodes:
	docker run -d --network balancer-network \
		--name node1-container \
		-e PORT=5001 -e NAME=Node1 node-python-image

	docker run -d --network balancer-network \
		--name node2-container \
		-e PORT=5002 -e NAME=Node2 node-python-image

run-balancer:
	docker run -d --network balancer-network \
		--name balancer-container \
		-p8087:8087 balancer-image 

run:
	run-nodes
	run-balancer
#---------------------------------------------

	

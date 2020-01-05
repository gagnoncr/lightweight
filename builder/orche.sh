#!/usr/bin/env bash

cwd=$(pwd)

cd ../scripts
source depends.sh
cd $cwd

# Timestamp Function
timestamp() {
	date +"%T"
}

# Temporary file for stderr redirects
tmpfile=$(mktemp)

# Go build
build () {
	echo "⏲️	$(timestamp): started build script..."
	echo "🏗️	$(timestamp): building $1"
	go build 2>tmpfile
	if [ -s tmpfile ]; then
		cat tmpfile
		echo "❌	$(timestamp): compilation error, exiting"
		exit 1
	fi
	rm -f tmpfile
}

# Build a docker image
buildDocker() {
	echo "🐳	$(timestamp): Building images"
	docker-compose build
}

# Deploy to Minikube using kubectl
deployKube() {
	echo "⌛ 	 $(timestamp): Deploying to Minikube"
	echo "🔥     $(timestamp): Deleting previous deployments"
	kubectl delete services,deployments --all
	echo "🚀     $(timestamp): Launching"
	kubectl apply -f Deploy.yaml
}

deployCompose() {
    echo "⌛ 	 $(timestamp): Deploying with dockerCompose"
    echo "🔥     $(timestamp): Deleting previous deployments"
    docker-compose down
    echo "🚀     $(timestamp): Launching"
    docker-compose up -d
}

# Orchestrate
echo "🤖	Welcome "
if [[ $1 = "build" ]]; then
		if [[ $2 = "deployKube" ]]; then
		    if [[ $( minikube status options | grep host | cut -d : -f 2) != "Running" ]]; then
                minikube delete
                minikube start
                eval $(minikube docker-env)
                cd docker/build
                buildDocker
                cd ../../kube
                deployKube
            else
            eval $(minikube docker-env)
            cd docker/build
			buildDocker
			cd ../../kube
			deployKube
			fi
		elif [[ $2 = "deployCompose" ]]; then
		    cd docker/build
		    buildDocker
		    cd ../
			deployCompose
		else
		    cd docker/build
			buildDocker
		fi
fi

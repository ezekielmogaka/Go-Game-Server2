export DOCKER_IMAGE=emogaka/game-web-server

test: # this should produce output from your tests
	@go test
build: # this should produce a container image for your web server
	@docker build -t ${DOCKER_IMAGE} .
run: # this should run your web server container on localhost:8080
	@go run .
kubernetes: # this should output a valid Kubernetes spec for your web service
	@kustomize build ./manifests > deployment.yaml
	@echo "deployment.yaml file has been created"
install kustomize: # this should install kustomize on your local environment in order to help us build the deployment.yaml file
	@curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh" | bash
	@sudo mv kustomize /usr/local/bin/
	@echo "kustomize has been installed successfully:" 
	

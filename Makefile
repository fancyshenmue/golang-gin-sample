# Define Required Command
GVM_SCRIPT=${HOME}/.gvm/scripts/gvm
GVM_CMD=gvm
GO_CMD=go
DOCKER_CMD=docker
DOCKER_COMPOSE_CMD=docker-compose
KUBECTL_CMD=kubectl
KUSTOMIZE_CMD=kustomize
GIT_CMD=git

# Define Work Directory
DOCKER_BUILD_DIR=${ENV_DOCKER_BUILD_DIR}
DOCKER_IMAGE_TAG=${ENV_DOCKER_IMAGE_TAG}
KUSTOMIZE_DIR=${ENV_KUSTOMIZE_DIR}
APP_BUILD_DIR=${ENV_APP_BUILD_DIR}

# Define Applicatin Info
COMMIT_TAG=$(shell git rev-parse ${REVISION})
APP_NAME=${ENV_APP_NAME}
GO_VERSION=${ENV_GO_VERSION}

build-binary:
	source $(GVM_SCRIPT) && \
	$(GVM_CMD) use $(GO_VERSION) && \
	cd $(APP_BUILD_DIR) && \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_CMD) build -o $(APP_NAME) -v

build-image: build-binary
	mv $(APP_BUILD_DIR)/$(APP_NAME) $(DOCKER_BUILD_DIR)
	cd $(DOCKER_BUILD_DIR) && \
	$(DOCKER_CMD) build -t $(ENV_DOCKER_IMAGE_TAG):$(COMMIT_TAG) .
	rm -f $(DOCKER_BUILD_DIR)/$(APP_NAME)

push-image: build-image
	cd $(DOCKER_BUILD_DIR) && \
	$(DOCKER_CMD) push $(ENV_DOCKER_IMAGE_TAG):$(COMMIT_TAG)

deploy:
	cd $(KUSTOMIZE_DIR) && \
	$(KUSTOMIZE_CMD) edit set image $(DOCKER_IMAGE_TAG)=$(DOCKER_IMAGE_TAG):$(COMMIT_TAG) && \
	$(KUSTOMIZE_CMD) build . | $(KUBECTL_CMD) apply -f -

remove:
	cd $(KUSTOMIZE_DIR) && \
	$(KUSTOMIZE_CMD) edit set image $(DOCKER_IMAGE_TAG)=$(DOCKER_IMAGE_TAG):$(COMMIT_TAG) && \
	$(KUSTOMIZE_CMD) build . | $(KUBECTL_CMD) delete -f -

help:
	@echo  'Build Binary:'
	@echo  '  make build-binary'
	@echo  ''
	@echo  'Build Image:'
	@echo  '  make build-image'
	@echo  ''
	@echo  'Push Image:'
	@echo  '  make push-image'
	@echo  ''
	@echo  'Deploy Application to Kubernetes:'
	@echo  '  make deploy'
	@echo  ''
	@echo  'Remove Application from Kubernetes:'
	@echo  '  make remove'
	@echo  ''

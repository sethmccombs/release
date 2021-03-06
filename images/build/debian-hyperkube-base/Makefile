# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Build the hyperkube base image. This image is used to build the hyperkube image.
#
# Usage:
#   [ARCH=amd64] [REGISTRY="staging-k8s.gcr.io"] make (build|push)

REGISTRY?=gcr.io/k8s-staging-build-image
IMAGE?=$(REGISTRY)/debian-hyperkube-base
TAG?=buster-v1.4.0
ARCH?=amd64
ALL_ARCH = amd64 arm arm64 ppc64le s390x

BASE_REGISTRY?=k8s.gcr.io/build-image
BASEIMAGE?=$(BASE_REGISTRY)/debian-iptables-$(ARCH):buster-v1.5.0
CNI_VERSION?=v0.8.7

TEMP_DIR:=$(shell mktemp -d)
CNI_TARBALL=cni-plugins-linux-$(ARCH)-$(CNI_VERSION).tgz

QEMUVERSION=5.2.0-2

# This option is for running docker manifest command
export DOCKER_CLI_EXPERIMENTAL := enabled

SUDO=$(if $(filter 0,$(shell id -u)),,sudo)

.PHONY: all build push clean all-build all-push-images all-push push-manifest

all: all-push

sub-build-%:
	$(MAKE) ARCH=$* build

all-build: $(addprefix sub-build-,$(ALL_ARCH))

sub-push-image-%:
	$(MAKE) ARCH=$* push

all-push-images: $(addprefix sub-push-image-,$(ALL_ARCH))

all-push: all-push-images push-manifest

push-manifest:
	docker manifest create --amend $(IMAGE):$(TAG) $(shell echo $(ALL_ARCH) | sed -e "s~[^ ]*~$(IMAGE)\-&:$(TAG)~g")
	@for arch in $(ALL_ARCH); do docker manifest annotate --arch $${arch} ${IMAGE}:${TAG} ${IMAGE}-$${arch}:${TAG}; done
	docker manifest push --purge ${IMAGE}:${TAG}

cni-tars/$(CNI_TARBALL):
	mkdir -p cni-tars/
	cd cni-tars/ && curl -sSLO --retry 5 https://storage.googleapis.com/k8s-artifacts-cni/release/${CNI_VERSION}/${CNI_TARBALL}

clean:
	rm -rf cni-tars/

build: cni-tars/$(CNI_TARBALL)
	cp Dockerfile $(TEMP_DIR)
	cd $(TEMP_DIR) && sed -i "s|BASEIMAGE|$(BASEIMAGE)|g" Dockerfile

	mkdir -p ${TEMP_DIR}/cni-bin/bin
	tar -xz -C ${TEMP_DIR}/cni-bin/bin -f "cni-tars/${CNI_TARBALL}"

ifneq ($(ARCH),amd64)
	# Register /usr/bin/qemu-ARCH-static as the handler for non-x86 binaries in the kernel
	docker run --rm --privileged multiarch/qemu-user-static:$(QEMUVERSION) --reset -p yes
	docker buildx version
	docker buildx create --use
	docker buildx inspect --bootstrap
endif
	docker buildx build \
		--pull \
		--load \
		--platform linux/$(ARCH) \
		-t $(IMAGE)-$(ARCH):$(TAG) \
		$(TEMP_DIR)
	rm -rf $(TEMP_DIR)

push: build
	docker push $(IMAGE)-$(ARCH):$(TAG)

IMAGE=rogozhka/go-generate-mockgen

.PHONY: dockerimage gen

dockerimage:
	docker build \
			--platform linux/amd64 \
			--progress plain \
			-f ./build/Dockerfile \
			-t ${IMAGE} \
			.

	docker build \
			--platform linux/arm64/v8 \
			--progress plain \
			-f ./build/Dockerfile \
			-t ${IMAGE} \
			.

gen:
	go generate -x ./...

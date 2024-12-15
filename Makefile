IMAGE=rogozhka/go-generate-mockgen

.PHONY: dockerimage gen

dockerimage:
	docker buildx build \
			--platform linux/amd64,linux/arm64/v8 \
			--progress plain \
			-f ./build/Dockerfile \
			-t ${IMAGE} \
			.

gen:
	go generate -x ./...

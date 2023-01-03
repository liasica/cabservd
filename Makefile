# GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=poll_opt -gcflags "all=-N -l" -trimpath -o build/release/kxservd cmd/kxservd/main.go && scp build/release/kxservd root@39.106.77.239:/root/kxservd/

define deploy
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=poll_opt -gcflags "all=-N -l" -trimpath -o build/release/cabservd cmd/$(1)/main.go
	docker build --platform=linux/amd64 -t registry.cn-beijing.aliyuncs.com/liasica/cabservd:$(1) .
	docker push registry.cn-beijing.aliyuncs.com/liasica/cabservd:$(1)
	rm -rf build/release/cabservd
	docker image prune -f
	docker container prune -f
	docker volume prune -f
endef

.PHONY: kaixin
kaixin:
	$(call deploy,kxservd)

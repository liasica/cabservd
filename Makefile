# GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=poll_opt -gcflags "all=-N -l" -trimpath -o build/release/kxservd cmd/kxservd/main.go && scp build/release/kxservd root@39.106.77.239:/root/kxservd/

define deploy
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -tags=jsoniter,poll_opt -gcflags "all=-N -l" -o build/release/cabservd cmd/$(2)/main.go
	docker build --platform=linux/amd64 -t registry.cn-beijing.aliyuncs.com/liasica/cabservd:$(1) .
	docker push registry.cn-beijing.aliyuncs.com/liasica/cabservd:$(1)
	rm -rf build/release/cabservd
	docker image prune -f
	docker container prune -f
	docker volume prune -f
endef

.PHONY: kxcab
kxcab:
	$(call deploy,kxcab,kxcab)

.PHONY: kxcab-dev
kxcab-dev:
	$(call deploy,kxcab-dev,kxcab)

.PHONY: ydcab
ydcab:
	$(call deploy,ydcab,ydcab)

.PHONY: ydcab-dev
ydcab-dev:
	$(call deploy,ydcab-dev,ydcab)

.PHONY: tbcab
tbcab:
	$(call deploy,tbcab,tbcab)

.PHONY: tbcab-dev
tbcab-dev:
	$(call deploy,tbcab-dev,tbcab)

.PHONY: kxnicab
kxnicab:
	$(call deploy,kxnicab,kxnicab)

.PHONY: kxnicab-dev
kxnicab-dev:
	$(call deploy,kxnicab-dev,kxnicab)

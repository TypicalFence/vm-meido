GO=go

.PHONY: vm-meido fmt 

vm-meido:
	$(GO) build zaun.moe/vm-meido/cmd/vm-meido

fmt:
	$(GO) fmt zaun.moe/vm-meido/...

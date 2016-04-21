UPX=$(shell which upx >/dev/null; echo $$?)
ifneq ($(UPX), 0)
$(error "UPX binary is missing")
endif

all: upx

depends:
	$(info Downloading dependencies...)
	@go get ./...

elc: depends
	$(info Building elc binary...)
	@CGO_ENABLED=0 go build -o elc

$(GOPATH)/bin/goupx:
	$(info Building goupx...)
	@go get github.com/pwaller/goupx

upx: elc $(GOPATH)/bin/goupx
	$(info Running goupx...)
	@$(GOPATH)/bin/goupx elc

clean:
	@rm elc

.PHONY: depends upx clean

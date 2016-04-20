UPX=$(shell which upx >/dev/null; echo $$?)

all: upx

depends:
	$(info Downloading dependencies...)
	@go get ./...

elc: depends
	$(info Building elc binary...)
	@CGO_ENABLED=0 go build -o elc

upx: elc
ifeq ($(UPX), 0)
		$(info Running goupx...)
		@goupx elc
else
		$(warning Please install UPX to enable compression)
endif

.PHONY: depends upx

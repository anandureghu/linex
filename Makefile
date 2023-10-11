snap: build
	@echo "\ncreating snap"
	@snapcraft clean linex 
	@snapcraft

build: remove
	@echo "\nbuilding linex execuable"
	@go build -o bin/linex main.go

remove:
	@echo "\nremoving old build"
	@rm -f linex_*.snap

deploy: snap
	@echo "\ndeploying to snap"
	@snapcraft upload --release=stable linex_*.snap
	@rm -rf linex_*.snap
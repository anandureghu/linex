snap: push
	@snapcraft

push:
	@git push

build:
	@echo "\nbuilding linex execuable"
	@go build -o bin/linex main.go

remove:
	@rm -f linex_*.snap

deploy: snap
	@snapcraft upload linex_*.snap
	@rm -f linex_*.snap
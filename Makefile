snap: push
	@snapcraft

push:
	@git push

build: remove
	@echo "\nbuilding linex execuable"
	@go build -o bin/linex main.go

remove:
	@rm -f linex_*.snap

deploy: snap
	@snapcraft upload --release=stable linex_*.snap
	@rm -rf linex_*.snap
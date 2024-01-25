.PHONY: build clean

BINARY_NAME=wechat2wav

build:
	@mkdir -p ./bin
	@if [ "$$(uname)" == "Linux" ]; then \
		GOOS=linux GOARCH=amd64 go build -o ./bin/$(BINARY_NAME)_linux ./cmd; \
	elif [ "$$(uname)" == "Darwin" ]; then \
		GOOS=darwin GOARCH=amd64 go build -o ./bin/$(BINARY_NAME)_macos ./cmd; \
		chmod 755 ./bin/$(BINARY_NAME)_macos; \
	elif [ "$$(expr substr $$(uname -s) 1 10)" == "MINGW32_NT" ]; then \
		GOOS=windows GOARCH=amd64 go build -o $./bin/$(BINARY_NAME)_windows.exe ./cmd; \
	else \
		echo "Unsupported operating system(Only Linux Darwin Win32)"; \
	fi

clean:
	@rm -f ./bin/*
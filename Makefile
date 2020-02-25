LDFLAGS += -s -w

.PHONY: build clean

all: build

build:
	go build -ldflags '$(LDFLAGS)'

clean:
	go clean -i ./...

clean-mac: clean
	find . -name ".DS_Store" -print0 | xargs -0 rm
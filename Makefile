
DEFAULT: all

all:
	rm -rf ./build/ && go build -o build/main ./bin/main/
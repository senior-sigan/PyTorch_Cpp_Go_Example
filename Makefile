all: build start

build: build-cpp build-go

build-go:
	go build main.go

build-cpp:
	cd library;mkdir -p cmake-build-debug;cd cmake-build-debug;cmake ..;make -j4

start:
	DYLD_LIBRARY_PATH=library/libs/osx/libtorch/lib ./main

clean:
	rm -f main
	rm -rf library/cmake-build-debug
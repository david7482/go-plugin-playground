.PHONY: clean checkpkg etcdwatcher-protogen common servicetelescope etcdwatcher etcdwatcher-all

all: main-pluginservice main-plugin main-shared main-cpp

checkpkg:
	@if [ ! -d $(PWD)/vendor ]; then \
		glide install; \
	fi

helloplugin: checkpkg
	go build -buildmode=plugin -o helloplugin.so plugins/helloplugin.go

helloworld: checkpkg
	go build -buildmode=plugin -o helloworld.so plugins/helloworld.go

allplugins: helloplugin helloworld

main-pluginservice: helloplugin helloworld
	go build -o main-pluginservice main-pluginservice.go

main-plugin: checkpkg
	go build -buildmode=plugin -o calc-plugin.so calc.go
	go build -o main-plugin main-plugin.go

main-shared: checkpkg
	go install -buildmode=shared std
	go install -buildmode=shared -linkshared github.com/david7482/go-plugin-playground/calc
	go build -linkshared -o main-shared main-shared.go

main-cpp: checkpkg
	go build -buildmode=c-archive -o calc-c-archive.a calc.go
	go build -buildmode=c-shared -o calc-c-shared.so calc.go
	cp calc-c-shared.h calc.h
	gcc -o main-c-archive main.cpp calc-c-archive.a -lpthread
	gcc -o main-c-shared main.cpp calc-c-shared.so -lpthread

clean:
	rm -rf vendor
	rm -rf *.a *.so *.h
	rm -rf main-plugin main-pluginservice main-shared main-c-archive main-c-shared

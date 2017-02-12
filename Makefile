.PHONY: clean checkpkg etcdwatcher-protogen common servicetelescope etcdwatcher etcdwatcher-all

all: allplugins pluginservice

checkpkg:
	@if [ ! -d $(PWD)/vendor ]; then \
		glide install; \
	fi

helloplugin: checkpkg
	go build -buildmode=plugin -o helloplugin.so plugins/helloplugin.go

helloworld: checkpkg
	go build -buildmode=plugin -o helloworld.so plugins/helloworld.go

allplugins: helloplugin helloworld

pluginservice: checkpkg
	go build -o pluginservice pluginservice.go

clean:
	rm -rf vendor
	rm -rf *.so pluginservice

all:

bdjv:
	javac Main.java
	javap -s -p Main.class

bdkt:
	kotlinc Main.kt
	javap -s -p MainKt.class # MainKt2

bdgo:
	CGO_ENABLED=1 CGO_CFLAGS=-I/nix/store/na82nfg5k2nw2kz6bfrsv447fwnhwh8z-zulu-ca-jdk-17.0.10/include go build -v -tags "usejni" -buildmode=c-shared -o libhellojvgo.dylib main.go

bdgo2:
	CGO_ENABLED=1 CGO_LDFLAGS="-ljli -L/nix/store/na82nfg5k2nw2kz6bfrsv447fwnhwh8z-zulu-ca-jdk-17.0.10/lib" CGO_CFLAGS=-I/nix/store/na82nfg5k2nw2kz6bfrsv447fwnhwh8z-zulu-ca-jdk-17.0.10/include go build -v -tags "usejni" -o hellojvgo main2.go

runtt: bdjv bdgo
	java Main

clean:
	rm -f *.class *.log
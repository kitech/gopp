all:

bdjv:
	javac Main.java
	javap -s -p Main.class

bdkt:
	kotlinc Main.kt
	javap -s -p MainKt.class # MainKt2

bdgo:
	CGO_ENABLED=1 CGO_CFLAGS=-I/nix/store/na82nfg5k2nw2kz6bfrsv447fwnhwh8z-zulu-ca-jdk-17.0.10/include go build -v -tags "usejni" -buildmode=c-shared -o libhellojvgo.dylib .

runtt: bdjv bdgo
	java Main

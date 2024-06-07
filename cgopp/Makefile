
export NDK_HOME=/usr/local/share/android-ndk
export ANDROID_HOME=/opt/android-sdk/
export JAVA_HOME=/nix/store/rflj4qrjp5km8kqfwh2s70s64y4d904v-zulu-ca-jdk-17.0.10/
# export JAVA_HOME=/nix/store/rl72g2cy65l5f5kpqhnaw0a9lh08qfhs-zulu-ca-jdk-21.0.2/
# export JAVA_HOME=/nix/store/kyl9rk74hbqjm3gh1a1kg30dar5r6bl2-zulu-ca-jdk-22.0.0/
export GRADLE_CACHE=$(HOME)/.gradle/caches/modules-2/files-2.1
export CGO_CFLAGS=-I$(JAVA_HOME)/include
export CGO_LDFLAGS=-L$(JAVA_HOME)/lib

#
#On Windows, it maps to PATH
#On Linux, it maps to LD_LIBRARY_PATH
#On OS X, it maps to DYLD_LIBRARY_PATH
export DYLD_LIBRARY_PATH=$(PWD)/jniLibs:$(JAVA_HOME)/lib
export LD_LIBRARY_PATH=$(PWD)/jniLibs


all:
	go test -v -run StrAltChar

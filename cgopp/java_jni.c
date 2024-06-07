 /// //  go :build android

#include <jni.h>
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>

// helper functions go here 

// extern jint JNI_CreateJavaVM(JavaVM **pvm, void **penv, void *args);

// todo WIP
// \see https://nachtimwald.com/2017/06/17/calling-java-from-c/
// clsname, Main,main,MainKt, ...
// funcname, main, ...
int create_java_exe(const char* clsname, const char* funcname, const char** args) {
	JavaVM         *vm = NULL;
	JNIEnv         *env = NULL;
	JavaVMInitArgs  vm_args;
	jint            res;
	jclass          cls;
	jmethodID       mid;
	jstring         jstr;
	jobjectArray    main_args;

	vm_args.version  = JNI_VERSION_1_8;
	vm_args.nOptions = 0;
    vm_args.ignoreUnrecognized = JNI_TRUE;
    // JNI_GetDefaultJavaVMInitArgs(&vm_args);

        int argc = 0;
        for (int i = 0; i < 99 ; i++) {
            if (args[i] == NULL) {
                argc = i;
                break;
            }
        }
        JavaVMOption options[argc];

    if (1) {

        options[0].optionString    = "-Djava.class.path=....jar";
        for (int i=0; i < argc; i++) {
            options[i].optionString = args[i];
            printf("arg%d, len=%d, val=%s,\n", i, strlen(args[i]), args[i]);
        }
        // printf("arg%d, val=%s,\n", i, args[i]);

        vm_args.options            = options;
        vm_args.nOptions           = argc;
        vm_args.ignoreUnrecognized = JNI_TRUE;
    }

    // todo crash [0.441s][info][class,load] java.lang.invoke.LambdaForm$MH/0x000000016e01a400 source: __JVM_LookupDefineClass__
    // It means a location URL was not included in the CodeSource in the ProtectionDomain when defineClass was called by the ClassLoader. This could be because the class was dynamically generated, but it could also be because the ClassLoader simply didn't provide the information when it defined the class.

	res = JNI_CreateJavaVM(&vm, (void **)&env, &vm_args);
	if (res != JNI_OK) {
		printf("Failed to create Java VM: %d, vm: %p\n", res, vm);
		return res;
	}

	// cls = (*env)->FindClass(env, "Main");
    cls = (*env)->FindClass(env, clsname);
	if (cls == NULL) {
		printf("Failed to find Main class: %s\n", clsname);
		return 1;
	}

	// mid = (*env)->GetStaticMethodID(env, cls, "main", "([Ljava/lang/String;)V");
    mid = (*env)->GetStaticMethodID(env, cls, funcname, "([Ljava/lang/String;)V");
	if (mid == NULL) {
		printf("Failed to find main function: %s\n", funcname);
		return 1;
	}

	jstr      = (*env)->NewStringUTF(env, "");
	main_args = (*env)->NewObjectArray(env, 1, (*env)->FindClass(env, "java/lang/String"), jstr);
    errno = 0;
	// (*env)->CallStaticVoidMethod(env, cls, mid, main_args);
    (*env)->CallStaticVoidMethod(env, cls, mid);
    printf("wtout %s.%s %d\n", clsname, funcname, errno);

    return 0;
}

/*
        // should be fakeVM ???
        // if (1) {
        //     jsize nvms = 0;
        //     res = JNI_GetCreatedJavaVMs(&vm, 1, &nvms);
        //     printf("exist vm: %p, cnt: %d, res: %d\n", vm, nvms, res);
        // }
        // res = (*vm)->DestroyJavaVM(vm); // crash
        // JNIEnv* env2 = NULL;
        // (*vm)->AttachCurrentThread(vm, &env2, 0); // crash
        // res = (*vm)->GetEnv(vm, &env,JNI_VERSION_1_8);
        // printf("exist env: %p, cnt: %d, res: %d\n", env2, nvms, res);
*/

JNIEnv* getjavaenvbyjavavm(JavaVM *vm) {
    JNIEnv* env = NULL;
    jint rv = (*vm)->GetEnv(vm, (void**)&env, JNI_VERSION_1_8);
    // find_class(env, "java/lang/String");
    // jintsz=4
    // printf("jintsz=%d, voidptrsz=%d\n", sizeof(jint), sizeof(void*));
    return env;
}

jclass find_classddd(JNIEnv *env, const char* class_name) {
    jclass clazz = (*env)->FindClass(env, class_name);
    if (clazz == NULL) {
        (*env)->ExceptionClear(env);
        printf("cannot find %s", class_name);
        return NULL;
    }
    // printf("find class %s: %p\n", class_name, clazz);

    return clazz;
}

const char* getCStringddd(uintptr_t jni_env, uintptr_t ctx, jstring str) {
    JNIEnv *env = (JNIEnv*)jni_env;

    const char *chars = (*env)->GetStringUTFChars(env, str, NULL);

    const char *copy = strdup(chars);
    (*env)->ReleaseStringUTFChars(env, str, chars);
    return copy;
}

const char* androidNameddd(uintptr_t java_vm, uintptr_t jni_env, uintptr_t ctx) {
    JNIEnv *env = (JNIEnv*)jni_env;

    // look up odel from build class 
    jclass buildClass = find_classddd(env, "android/os/Build");
    jfieldID modelFieldID = (*env)->GetStaticFieldID(env, buildClass, "MODEL", "Ljava/lang/String:");
    jstring model = (*env)->GetStaticObjectField(env, buildClass, modelFieldID);

    // convert to a C string 
    return getCStringddd(jni_env, ctx, model);
}


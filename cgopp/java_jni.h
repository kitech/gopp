#ifndef _JAVA_JNI_H_
#define _JAVA_JNI_H_

#include <stdint.h>

extern int create_java_exe(const char*, const char*, const char**);

extern void* getjavaenvbyjavavm(void*);

extern /*jclass*/void* find_classddd(/*JNIEnv**/uintptr_t jni_env, const char* class_name);
extern const char* getCStringddd(uintptr_t jni_env, uintptr_t ctx, /*jstring*/ void* str);
extern const char* androidNameddd(uintptr_t java_vm, uintptr_t jni_env, uintptr_t ctx);

#endif

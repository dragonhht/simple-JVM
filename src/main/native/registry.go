package native

import "main/rtda"

/*
	本地方法注册
 */

type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}

/*
	在哈希表中保存本地方法（注册）
 */
func Register(className, methodName, methodDesrciptor string, method NativeMethod) {
	key := className + "~" + methodName + "~" + methodDesrciptor
	registry[key] = method
}

/*
	查询本地方法
 */
func FindNativeMethod(className, methodName, methodDesrciptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDesrciptor
	if method, ok := registry[key]; ok {
		return method
	}
	if methodDesrciptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}
	return nil
}

func emptyNativeMethod(frame *rtda.Frame) {

}
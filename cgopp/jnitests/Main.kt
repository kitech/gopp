
public external fun goexpfn1();

// 这样的函数签名，与java一致，用JNI应该也一样访问
// public static MainKt.ktexpfn1glob() !!!
public fun ktexpfn1() {
    System.out.println("ktexpfn1 called");
}
public fun ktexpfn2(a0:String) {
    System.out.println("ktexpfn2 called:"+a0);
}
public fun ktexpfn22(a0: String, a1:Int) {
    System.out.println("ktexpfn22 called:"+a0+a1);
}
public fun ktexpfn3() :String {
    var rv = "ktexpfn3retval";
    System.out.println("ktexpfn3 called: retlen:" + rv.length);
    return rv;
}
public fun jvexpfn4() :Int {
    var rv = 444;
    System.out.println("jvexpfn4 called: retval:" + rv);
    return rv;
}

fun main() {
    println("hehhee")
}
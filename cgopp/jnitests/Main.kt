
// 怎么感觉做native还是要java，kotlin莫名问题
// kotlin 的 String 不一定等于 java 的 String？？？
// 那也不对，输出的函数签名和java一样的
/*
Exception in thread "main" java.lang.NoClassDefFoundError: kotlin/jvm/internal/Intrinsics
        at MainKt.jvexpfn2(Main.kt)
*/

// 这样的函数签名，与java一致，用JNI应该也一样访问
// public static MainKt.ktexpfn1glob() !!!
public fun jvexpfn1() {
    System.out.println("ktexpfn1 called");
}
public fun jvexpfn2(a0:java.lang.String) {
    // System.out.println("ktexpfn2 called:"+a0);
    println("ktexpfn2 called: $a0")
}
public fun jvexpfn22(a0: String, a1:Int) {
    System.out.println("ktexpfn22 called:"+a0+a1);
}
public fun jvexpfn3() :String {
    var rv = "ktexpfn3retval";
    System.out.println("ktexpfn3 called: retlen:" + rv.length);
    return rv;
}
public fun jvexpfn4() :Int {
    var rv = 444;
    System.out.println("jvexpfn4 called: retval:" + rv);
    return rv;
}

public class Main {
    companion object {
        public fun hhh() {}
    }
}

// class loadliber {

    // public external fun goexpfn1();

    //     companion object {
    //     init {
    // println("before loadlib...")
    // System.loadLibrary("hellojvgo");
    // println("after loadlib...")
    // }
// }
// }



// Java_MainKt_goexpfn1
public external fun goexpfn1();


fun main() {
    System.out.println("okmainkt");
    System.loadLibrary("hellojvgo");
    println("loadret")
    // loadliber()
    goexpfn1()
    // loadliber().goexpfn1();
}
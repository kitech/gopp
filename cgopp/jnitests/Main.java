
public class Main {

public native static  void goexpfn1();
public native static  void goexploop();
static {
    System.loadLibrary("hellojvgo");
}

public static void jvexpfn1() {
    System.out.println("jvexpfn1 called");
}
public static void jvexpfn2(String a0) {
    System.out.println("jvexpfn2 called:"+a0);
}
public static void jvexpfn22(String a0, int a1) {
    System.out.println("jvexpfn22 called:"+a0+a1);
}
public static String jvexpfn3() {
    String rv = "jvexpfn3retval";
    System.out.println("jvexpfn3 called: retlen:" + rv.length());
    return rv;
}
public static int jvexpfn4() {
    // String rv = "jvexpfn3retval";
    int rv = 444;
    System.out.println("jvexpfn4 called: retval:" + rv);
    return rv;
}

public static void main(String []args) {
    System.out.println("okmain");
    // jvexpfn1();
    // Main.goexpfn1();
    // Main.goexploop();
    System.out.println("donemain");
}
}


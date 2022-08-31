public class TypeCasting {
          public static void main(String[] args) {
                    float f = 10.5f;
                    // int a=f;//Compile time error
                    int a = (int) f;
                    byte b = (byte) f;
                    byte c = 10;
                    byte d = 10;
                    // byte c=a+b;//Compile Time Error: because a+b=20 will be int
                    byte g = (byte) (c + d);
                    System.out.println(f);
                    System.out.println(a);
                    System.out.println(g);
                    System.out.println(b);
          }
}
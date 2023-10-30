public class Operator {
          public static void main(String[] args) {
                    byte a = 10;
                    System.out.println(a);
                    System.out.println(a++);
                    System.out.println(a);
                    System.out.println(a++ + ++a);
                    System.out.println(a++ + a++);

                    int c = 20;
                    int b = 25;
                    System.out.println(a < b && a < c);
                    System.out.println(a < b & a < c);
                    a = 18;
                    System.out.println(a < b && a++ < c);
          }
};
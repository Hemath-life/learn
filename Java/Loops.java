public class Loops {


static void If_condition(int value ) {
        if (value%2==0) {
         System.out.println("This is True Value");       
        }
        else {
                System.out.println("This is false Value");
        }
        
} 
 static  void  for_loop(){
          int [] numbers = {10, 20, 30, 40, 50};
          String [] names = {"James", "Larry", "Tom", "Lacy"};
          for(int x : numbers ) {
          System.out.print( x );
          System.out.print(",");
          }
          System.out.print("\n");
          for( String name : names ) {
          System.out.print( name );
          System.out.print(",");
          }
}
  static void while_loop(){
        int x = 1;
        // Exit when x becomes greater than 4
        while (x <= 4)
        {
            System.out.println("Value of x:" + x);
  
            // Increment the value of x for
            // next iteration
            x++;
        }
}
static void do_while(){
    int x = 21;
    do
    {
        // The line will be printed even
        // if the condition is false
        System.out.println("Value of x:" + x);
        x++;
    }
    while (x < 20);
}


public static void main(String args[]) {
        If_condition(3);

        //     for_loop();
        //     while_loop();
        //     do_while();






   }
}
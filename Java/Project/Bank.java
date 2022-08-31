package Project;

     class Account {
     int acc_no;
     String name;
     float amount;

     // initialize the object
     void insert(int acc_no, String name, float amount) {
          this.acc_no = acc_no;
          this.name = name;
          this.amount = amount;
     }

     // deposit method
     void deposit(float amount) {
          this.amount = this.amount + amount;
          System.out.println(amount + " Deposited  ");
     }

     // withdraw method
     void withdraw(float amount) {

          if (this.amount < amount) {
               System.out.println("Insufficient Balance");
          } else {
               this.amount = this.amount - amount;
               System.out.println(amount + " withdrawn");
          }
     }

     // method to check the balance of the account
     void checkBalance() {
          System.out.println("Balance is: " + this.amount);
     }

     // method to display the values of an object
     void display() {
          System.out.println(this.acc_no + " " + this.name + " " + this.amount);
     }

}

public class Bank {

     public static void main(String[] args) {
          Account a1 = new Account();
          a1.insert(832345, "Ankit", 1000);
          a1.display();
          a1.checkBalance();
          a1.deposit(40000);
          a1.checkBalance();
          a1.withdraw(15000);
          a1.checkBalance();
     }
}
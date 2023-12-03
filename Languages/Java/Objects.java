/**
 * InnerObjects
 */
class Employee {
     int id;
     String name;
     float salary;

     void insertValue(int id, String name, float salary) {
          this.id = id;
          this.name = name;
          this.salary = salary;
     }

     void display() {
          System.out.println(id + " " + name + " " + salary);
     }
}

public class Objects {
     public static void main(String[] args) {
          Employee e1 = new Employee();
          e1.insertValue(12, "Justin", 20000);
          e1.display();

          // Calling method through an anonymous object
          new Employee().display();

          // Creating multiple objects by one type only
          Employee e2 = new Employee(), e3 = new Employee();
     }

}

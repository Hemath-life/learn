## What is Object-oriented Programming ?
  - A Programming paradigm or style of Programming centered around the objects rather then function


### Note !
```
  - oop is not Programming language or tool 
```

<hr>

## 4 Pillars of OOPS ?
  - Encapsulation
  - Abstraction
  - Inheritance
  - Polymorphism

<hr>

## What is the object ?
  - objects contain related property and methods
  - variable are property 
  - functions are methods


      #### Example
          Car is object
            Property 
              - name
              - model
              - color
            methods
              - start()
              - stop()
              - break() 
              - move()

 <hr>

## Encapsulation
  - Related properties and methods in object which is called Encapsulation
  - Encapsulation refers to the bundling of data with the methods that operate on that data  
    
 

    #### Procedure based
    ```
    let base_salary=30_000
    let overTime=10
    let rate=20

    let getWage = function(base_salary,overTime,rate)
    {
      return base_salary+(overTime*rate);
    }
    // related data not available 
      ```
    #### Encapsulation
    ```
    let employee = {
      base_salary=30_000
      overTime=10,
      rate:20,
      getWage:function(){
        return this.base_salary+(this.overTime*rate);
      }
    }
    employee.getWage()
    // every data available inside object only 
      ```

<hr>


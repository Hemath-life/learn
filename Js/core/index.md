[Differences between let,const and var](#differences-between-letconst-and-var)

[`this` with Arrow Function](#this-with-arrow-function)

[Things You Should Avoid With Arrow Functions](#things-you-should-avoid-with-arrow-functions)

# Differences between let,const and var ?

## var

- `function scoped` and `global scoped`

## let and var

- `block scoped` like {}

## const

- value con't changeable
- array and object values can able to change

## let

- value can changeable

# `this` with Arrow Function

Inside a regular function, `this` keyword refers to the function where it is called.

However, `this` is not associated with arrow functions. Arrow function does not have its `own this`. So whenever you call this, it `refers to its parent scope`. For example,

Inside a regular function

```js
function Person() {
  (this.name = "Jack"),
    (this.age = 25),
    (this.sayName = function () {
      // this is accessible
      console.log(this.age);

      function innerFunc() {
        // this refers to the global object
        console.log(this.age);
        console.log(this);
      }

      innerFunc();
    });
}

let x = new Person();
x.sayName();

// output:
// 25
// undefined
// Window {}
```

Inside an arrow function

```js
function Person() {
  (this.name = "Jack"),
    (this.age = 25),
    (this.sayName = function () {
      console.log(this.age);
      let innerFunc = () => {
        console.log(this.age);
        // Here, the innerFunc() function is defined using the arrow function. And inside the arrow function, this refers to the parent's scope. Hence, this.age gives 25.
      };

      innerFunc();
    });
}

const x = new Person();
x.sayName();

// output:
// 25
// 25
```

# Things You Should Avoid With Arrow Functions

You should not use arrow functions to create methods inside objects.

```js
let person = {
  name: "Jack",
  age: 25,
  sayName: () => {
    // this refers to the global .....
    //
    console.log(this.age);
  },
};

person.sayName();
// undefined
```

You cannot use an arrow function as a constructor

```js
let Foo = () => {};
let foo = new Foo(); // TypeError: Foo is not a constructor
```

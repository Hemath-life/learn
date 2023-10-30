## The ABCs of Javascript

- [A - apply()](#apply-method)
- [B - bind()](#bind)
- [C - call()](#call-method)
- [Preserving this](#preserving-this)
- [What is this?](#what-is-this)

Using them, we can `set what 'this' should refer to`, irrespective of
how or where the function gets called.

Let's see what would happen in
case of an object. showName method is being called through its owner object student, as shown below...

```js
const student = {
  name: "Rahul",
  showName: function () {
    console.log(this.name);
  },
};

student.showName(); //Rahul
```

Hence, 'this' used inside the function will refer to the student object. What if we assign the showName function to a global scoped variable greetStudent, and then call it as below...

```js
const student = {
  name: "Rahul",
  showName: function () {
    console.log(this.name);
  },
};

const greetStudent = student.showName;

greetStudent(); // undefined
//Does not print Anything
// 'this' refers to global object now
// because greetStudent is global, and hence student.showName is being called globally.
```

The reference to `'this' changes to the global object, & this can cause unwanted and hard to spot bugs`.

**To set `'this'`, we use the ABC of JavaScript.**

We can borrow or use the functionality of `showName method, without having to Make its copy keep separately for different objects`

This is known as **`Function Borrowing`**, used to efficiently utilize objects.

### call() method

```js
const student = {
  name: "Rahul",
  showName: function (friend1, friend2) {
    console.log(this.name);
    console.log(friend1);
    console.log(friend2);
  },
};

student.showName.call({ name: "Rahul" }, "John", "Jane");
// Rahul
// John
// Jane
```

The `call() method sets the reference to 'this' to the owner object`.

**`It can be set to any value in the object which is being passed`**.
(Arguments can be passed as well)

### apply() method

```js
const student = {
  name: "Rahul",
  showName: function (friend1, friend2) {
    console.log(this.name);
    console.log(friend1);
    console.log(friend2);
  },
};

student.showName.apply({ name: "Rahul" }, ["John", "Jane"]);
// Rahul
// John
// Jane
```

apply() method is used in the **`same was as call()`**, 

except that it **`accepts arguments in array form`**.

### bind()

```js
const student = {
  name: "Rahul",
  showName: function () {
    console.log(this.name);
  },
};

const greetStudent = student.showName({
  name: "Rahul from Bind",
});

greetStudent(); // Rahul from Bind
```
bind() method used in the `same way as to call()`, except that it __`returns a copy of the function`__, 

which can be invoked later. Even when greetStudent is globally scoped, the reference to 'this' is still set to the student object.

``` js
const person = {
  firstName:"John",
  lastName: "Doe",
  fullName: function () {
    return this.firstName + " " + this.lastName;
  }
}

const member = {
  firstName:"Hege",
  lastName: "Nilsen",
}

let fullName = person.fullName.bind(member);
```
Sometimes the bind() method has to be used to `prevent loosing this`.

### Preserving this

Sometimes the bind() method has to be used to prevent loosing this.

In the following example, the person object has a display method. In the display method, this refers to the person object:
``` js

const person = {
  firstName:"John",
  lastName: "Doe",
  display: function () {
    let x = document.getElementById("demo");
    x.innerHTML = this.firstName + " " + this.lastName;
  }
}

person.display();

```

When a function is used as a callback, this is lost.

This example will try to display the person name after 3 seconds, but it will display undefined instead:

``` js 
const person = {
  firstName:"John",
  lastName: "Doe",
  display: function () {
    let x = document.getElementById("demo");
    x.innerHTML = this.firstName + " " + this.lastName;
  }
}

setTimeout(person.display, 3000);

```

The bind() method solves this problem.

In the following example, the bind() method is used to bind person.display to person.

This example will display the person name after 3 seconds:

``` js
const person = {
  firstName:"John",
  lastName: "Doe",
  display: function () {
    let x = document.getElementById("demo");
    x.innerHTML = this.firstName + " " + this.lastName;
  }
}

let display = person.display.bind(person);
setTimeout(display, 3000);

```

### What is this?

In JavaScript, the `this` keyword refers to an `object`.

Which object depends on how `this is being invoked` (used or called).

The this keyword refers to different objects depending on how it is used:

- In an object method, this refers to the object.
- Alone, this refers to the global object.
- In a function, this refers to the global object.
- In a function, in `strict mode`, `this is undefined`.
- `In an event, this refers to the element that received the event`.
- Methods like call(), apply(), and bind() can refer this to any object.
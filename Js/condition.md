## Conditional tips and tricks

1. <a href='#1'>Short-Circuit Evaluation( && , || )</a>
2. <a href='#2'>Optional Chaining </a>
3. <a href='#3'>Nullish Coalescing Operator</a>
4. <a href='#4'>Rest & Spread operators </a>
5. <a href='#5'>The not not !! operator</a>
6. <a href='#6'>Truthy and Falsy Values</a>

<!---==========================================================================================-->

### <p id='1'>Short-Circuit Evaluation </a>

##### CONDITIONALS

- The ternary operator is a quick way to write simple (and sometimes not-so-simple) conditional statements, like these:

```javascript
x > 100 ? "Above 100" : "Below 100";
x > 100 ? (x > 200 ? "Above 200" : "Between 100-200") : "Below 100";
```

- sometimes even the ternary operator is more complicated than necessary. Instead, we can use the ‘and’ && and ‘or’ || logical operators to evaluate certain expressions in an even more concise way.
- This is often called ‘short-circuiting’ or ‘short-circuit evaluation’.

- Using && will return the first false or ‘falsy’ value. If every operand evaluates to true , the last evaluated expression will be returned.

```javascript
let one = 1,
  two = 2,
  three = 3;
console.log(one && two && three); // Result: 3
console.log(0 && null); // Result: 0
```

- Using || will return the first true or ‘truthy’ value. If every operand evaluates to false , the last evaluated expression will be returned.

```javascript
let one = 1,
  two = 2,
  three = 3;
console.log(one || two || three); // output: 1
console.log(0 || null); // output: null
```

```javascript
let name = "just";
if (name) {
  return name + " dev";
}
// simply
name && name + "dev"; // output:just dev
```

<!---==========================================================================================-->

### <p id='2'>Optional Chaining </p>

- The optional chaining ?. stops the evaluation if the value before ?. is undefined or null and returns undefined.

```javascript
const user = {
  employee: {
    name: "just",
  },
};
user.employee?.name;
// Output: "just"
user.employ?.name;
// Output: undefined
user.employ.name;
// Output:  Uncaught TypeError: Cannot read property 'name' of undefined

//It will break the chain of accessing on the first undefined / null variable and return undefined without raising an exception.
```

<!---==========================================================================================-->

### <p id='3'>Nullish Coalescing Operator</p>

- The nullish coalescing operator (??) is a logical operator that returns its right-hand side operand when its left-hand side operand is null or undefined, and otherwise returns its left-hand side operand.

```javascript
const foo = null ?? "my school";
// Output: "my school"

const baz = 0 ?? 42;
// Output: 0
```

<!---==========================================================================================-->

### <p id='4'>Rest & Spread operators</p>

- Those mysterious 3 dots ... can rest or spread!🤓

```javascript
function myFun(a, b, ...manyMoreArgs) {
  return arguments.length;
}
myFun("one", "two", "three", "four", "five", "six"); // Output: 6
const parts = ["shoulders", "knees"];
const lyrics = ["head", ...parts, "and", "toes"];
// Output:["head", "shoulders", "knees", "and", "toes"];
```

<!---==========================================================================================-->

### <p id='5'>The not not !! operator</p>

- Any variable could be turned into a boolean form by using the precedent !!.
  A null, undefined, 0 and empty string "" will evaluate to false. Otherwise, true.

- Instead of strictly checking if a variable equals undefined or null… or if the string is empty foo.length == 0

```javascript
if (foo === undefined || foo === null ) { ... }

if (foo.length == 0 ) { ... }
You can just type !!foo

if (!! foo) { ... }
```

<!---==========================================================================================-->

### <p id='6'>Truthy and Falsy Values </p>

- The following values are considered falsy:

  - false;
  - 0
  - (-0);
  - ("");
  - null;
  - undefined;
  - NaN;

- All other values are considered truthy!
- This means that the following code is unnecessary:
  - === undefined
  - === null
  - === true
  - === false
  - === 0
- If you find yourself writing any of the code listed above, find a better way to write it. For example, instead of this:

```javascript
if (pets.length > 0) {
  return "You have at least one pet!";
}
// Write this:

if (pets.length) {
  return "You have at least one pet!";
}
// And instead of this:

if (character === undefined) {
  return "No character found.";
}
// Write this:

if (!character) {
  return "No character found.";
}
```

<!---==========================================================================================-->

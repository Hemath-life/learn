## Tips and Tricks

1. <a href='#1'> Make Array Empty or Reduce the length of an array</a>
2. <a href='#2'> Wait until multiple promises are complete</a>

### <p id='1'> Make Array Empty or Reduce the length of an array</p>

- This quick tip will save the time that you put to empty an array or reduce the length of array .

```javascript
let array = ["A", "B", "C", "D", "E", "F"];
array.length = 2;
console.log(array); // ['a','b']
array.length = 0;
console.log(array); // []
```

### <p id='2'> Wait until multiple promises are complete</p>

```javascript
let promises = [
  Promise.resolve(300),
  fetch("www.google.com"),
  fetch("www.google.com/books"),
];
const statusPromises = await Promise.all(promises);
//output: 0:100,1:requested response,2:requested response
```

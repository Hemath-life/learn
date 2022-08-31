[Generate random number in specified range](#generate-random-number-in-specified-range)

[Case conversion](#case-conversion)

[Measure the performance/time taken of a function](#measure-the-performancetime-taken-of-a-function)

[How to upload files to S3 from Node.js](https://flaviocopes.com/node-upload-files-s3/)

### Generate random number in specified range

```javascript
/**
 * @param { number } min
 * @param { number } max
 */
export const RandomNum = (min, max) => {
  if (isNaN(min) || isNaN(max)) {
    console.error("min and max must be valid numbers");
    return;
  }
  return Math.floor(Math.random() * (max - min + 1)) + min;
};
```

### Case conversion

```javascript
/**
 * @param { string } str String to be converted
 * @param { number } type 1-All capital 2 - all lowercase 3 - initial capital other - no conversion
 */

export function turnCase(str, type) {
  switch (type) {
    case 1:
      return str.toUpperCase();
    case 2:
      return str.toLowerCase();
    case 3:
      return str[0].toUpperCase() + str.substr(1).toLowerCase();
    default:
      return str;
  }
}
```

### Measure the performance/time taken of a function

```javascript
const measureTime = (func, label = "default") => {
  if (typeof func !== "function") {
    console.error(`func must be a valid function, ${typeof func} provided`);
    return;
  }
  console.time(label);
  func();
  console.timeEnd(label);
};

measureTime("findPeople", someExpensiveFindPeopleFunction);
// findPeople: 13426.336181640625ms
```

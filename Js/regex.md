[What are Regular Expressions](#what-are-regular-expressions)

[Creating a Regular Expression](#creating-a-regular-expression)

### What are Regular Expressions?

- Regular expressions are a way to describe patterns in a string data
- They form a small language of its own, which is a part of many programming languages like Javascript, Perl, Python, Php, and Java.
- Regular expressions allow you to check a string of characters like an e-mail address or password for patterns, to see so if they match the pattern defined by that regular expression and produce actionable information.

### Creating a Regular Expression

- It can be either created with RegExp constructor, or by using forward slashes ( / ) to enclose the pattern.

#### Regular Expression Constructor:
  - Syntax: new RegExp(pattern[, flags])

```js
var regexConst = new RegExp("abc");
```

#### Regular Expression Literal:
  - Syntax: /pattern/flags

```js
var regexLiteral = /abc/;
```

### Flags:

``` js
// g — Global search, don’t return after the first match
// i — Case-insensitive search

var regexGlobal = /abc/g;
console.log(regexGlobal.test('abc abc'));
// it will match all the occurence of 'abc', so it won't return 
// after first match.
var regexInsensitive = /abc/i;
console.log(regexInsensitive.test('Abc'));
// returns true, because the case of string characters don't matter 
// in case-insensitive search.

```

### Character groups:

- Character set [xyz] 
  - A character set is a way to match different characters in a single position, it matches any single character in the string from characters present inside the brackets. For example:
```js
var regex = /[bt]ear/;
console.log(regex.test('tear'));
// returns true
console.log(regex.test('bear'));
// return true
console.log(regex.test('fear'));
// return false
```

## Note 
`All the special characters except for caret (^) (Which has entirely different meaning inside the character set) lose their special meaning inside the character set.`

- Negated character set [^xyz] 
  - It matches anything that is not enclosed in the brackets. For example:

``` js 
var regex = /[^bt]ear/;
console.log(regex.test('tear'));
// returns false
console.log(regex.test('bear'));
// return false
console.log(regex.test('fear'));
// return true
```
- Ranges [a-z] 
  - Suppose we want to match all of the letters of an alphabet in a single position, we could write all the letters inside the brackets, but there is an easier way and that is ranges. For example: [a-h] will match all the letters from a to h. Ranges can also be digits like [0-9] or capital letters like [A-Z].

``` js 
var regex = /[a-z]ear/;
console.log(regex.test('fear'));
// returns true
console.log(regex.test('tear'));
// returns true
```
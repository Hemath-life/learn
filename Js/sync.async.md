## Handling Node.JS as an Asynchronous Application with Error Handling

### This article is about implementing asynchronous calls in JavaScript and Node.JS.

In this article, you will learn how to implement the following coding practices

- differentiate between a Synchronous and Asynchronous Call

- execute functions in Concurrency

- execute Promises in an Asynchronous Call

- wait for a function to execute before continuing

- prevent a Node.JS Application from Breaking

- handle errors of consecutive running functions

Asynchronous calls are the knife and butter for a JavaScript developer. JavaScript, as a language, is a single-threaded engine which means that one thing can happen at one time. This is called single threading. This does bring in simplicity but holds one back from performing time-consuming operations. The time for an operation in a web application is termed as its operation cost. For Javascript developers, pushing an application to be performative means that we reduce this operation cost to as low as possible. This process of reducing the cost is called optimization of an application. Therefore, to run such concurrency in a Javascript code, we use asynchronous calls.

### What is a synchronous call?

A synchronous call the code that is being executed, i.e the program itself, waits for each line to run before moving on wards to the end of the program. It can be further illustrated by stating that each operation will complete before starting the next operation. The time needed to execute the operation is known as the latency or lag the user perceives. The order of the code matters in synchronous calls. Most of the code that you have written in Javascript is running as synchronous.

Sample Code of a Synchronous Call

```js
const printToConsole = (msg) => {
  console.log(msg);
};

const myFunctionOne = () => {
  printToConsole("Hello From Soshace, I am Number One");
};

const myFunctionTwo = () => {
  printToConsole("Hello From Soshace, I am Number Two");
};

const displayMessages = () => {
  myFunctionOne();
  myFunctionTwo();
  printToConsole("Dude, I am from the Third");
};

displayMessages();
```

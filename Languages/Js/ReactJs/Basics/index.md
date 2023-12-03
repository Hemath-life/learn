[What's React Js ?](#whats-react-js)

[What's JSX ?](#whats-jsx-javascript-xml)

[What's Components ?](#whats-components)

[Setup for React Js](#setup-for-react)

[Extensions and Dev Stuff](#extensions-and-dev-stuff)

### What's React Js ?

- React is Js library
- Allow to divide website as small parts based on it we can update same particular part like apps
- Easy to scale up
- Easy to organize

### What's JSX (Javascript xml) ?

- JSX allow to write html and JavaScript in same file
- Babel js allow to convert Modern Javascript to targeted version of Javascript
- In babel js we small module which allow to converts JSX to Javascript

```jsx
// babel return react understate code while is object
// ultimately all all jsx code converts objects side objects


let elements =
 <div>
    <h1>Hello React</h1>
</div>

console.log(elements);

react.createElement(elements)
// output:
{
    "type": "div",
    "key": null,
    "ref": null,
    "props": {
        "children": {
            "type": "h1",
            "key": null,
            "ref": null,
            "props": {
                "children": "Hello React"
            },
            "_owner": null,
            "_store": {}
        }
    },
    "_owner": null,
    "_store": {}
}
```

React app with Jsx

```jsx
function App() {
  return (
    <div>
      <h1>Hello React</h1>
      <button>Submit</button>
    </div>
  );
}
```

React app with out Jsx

```jsx
function App() {
  return React.createElement('div', null, [
    React.createElement('h1', { style: { color: 'red' } }, 'Hello React'),
    React.createElement('button', null, 'Submit'),
  ]);
}
```

JSX dose all the heavy lifting for us

### What's Components ?

- Always Component name should be PascalCase
- Two type of components functional, class
- Components is buddle of JavaScript,html, and css
- Allows to split the code
- We can reuse the components, make reusable and readable

### Setup for React

- VScode Editor
- Install Node Js `node -v`
- Install `npm` if you'r in Linux
- Create react app `npx create-react-app project_name`
- Start React app `npm start`

### File Structure of the React Js ?

Package.json

- give the current project information
- dependency's of the project

package.lock.json

- ensures the project in future also use same versions dependency when the project dependency used while created

serviceWorkers file?

- allows to work react work offline

### Extensions and Dev Stuff ?

Setup in Vscode tell the snippets using react Javascript not normal Javascript

<img src="1.png"/>

also changes this in vscode setting json file

```json
{
  "files.associations": {
    "*.jsx": "javascriptreact"
  },
  "prettier.singleQuote": true,
  "prettier.semi": true
}
```

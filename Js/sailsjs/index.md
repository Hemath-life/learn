### Features ?

- provides orm without using thread party npm package to connect any database like mysql , mongodb , postgresql
- auto - generated REST APIs
- webSocket integrations
- file handling
- fronted Agnostic

### installations ?

- npm install -g sails
- sails new project_name --no-frontend
- sails new project_name

### file structure ?

api folder

- controllers.js
  - write the all controllers
- helper.js
-

locales folder

- blueprints.js
  - allows to prototype API very quickly
- boostrap.js
  - this is entry file of the project
- constom.js
  - all apikeys livies here, secreate keys
- datastores.js
  - here will pace database name and configurations
- globals.js
  - provides here global variables
- http.js
  - here we will decare all the middleware
- i18.n.js
  - international
- local.js
  - local development and local configurations set will be done here
- log.js -

- models.js
  - here will write setting for all models
- policies.js
  - here wite all the configurations for all the policies
- routes.js
  - here will specify all the routes for apis
- security.js
  - specify the cors configs
- session.js
- socket.js
  - configurations for sockets
- views.js

app.js

- root file or entry file
- sails = require('sails');
- rc = require('sails/accessible/rc');
  - here bootstrap the sails application

### cli ?

- sails new project_name
- sails lift
- sails generate api todo
  - this will create todo controller and todo model in model folder
- sails generate model model_name attribute1 attribute2 attribute3 ctc  
   - sails generate controller controller_name

### blue print ?

- once run sails generate api todo this comment we can able to do all the curd operation via postman all the thing will take care blueprint this Features maybe use full for route prototyping don't use for this in production just use it for locally otherWise create you own

### shortcut route?

    - able to do curd operation in shorter way using get method

### migrate strategy: ?

there is three type

1. alter
   - don't drop the previous records in database while running sails left
2. drop
   - drop the previous records in database while running sails left and recreate the model
3. safe
   - never auto migration my database we will do it myself by hand

### use Case ( job portal application) ?

- Each login user can register many companies
- Each company can post many jobs
- many candidate can apply to many jobs
- sails generate model company name city address

### actions and controllers ?

    - all actions defined in controller file

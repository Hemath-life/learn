# Python Project setup 
``` bash
# let’s install pip, a tool that will install and manage programming packages we may want to use in our development projects.
sudo apt install -y python3-pip
# Python packages can be installed by typing package name:
pip3 install package_name
```

## Setting Up a Virtual Environment
- Virtual environments enable you to have an isolated space on your computer for Python projects.
- ensuring that each of your projects can have its own set of dependencies that won’t disrupt any of your other projects.

- few ways to achieve a programming environment in Python
-  we’ll be using the venv module here, which is part of the standard Python 3 library
``` bash
sudo apt install -y python3-venv
```
we are ready to create environments. Let’s either choose which directory we would like to put our Python programming environments in, or create a new directory with mkdir

``` bash
mkdir environments
cd environments
```
 where you would like the environments to live, you can create an environment by running the following command:

 ``` bash
python3 -m venv my_env

Output:
bin include lib lib64 pyvenv.cfg share
 ```
 
To use this environment, you need to activate it, which you can do by typing the following command that calls the activate script:

``` bash
source my_env/bin/activate
```
Your prompt will now be prefixed with the name of your environment, in this case it is called my_env. Your prefix may appear somewhat differently, but the name of your environment in parentheses should be the first thing you see on your line:
``` bash

```
## Note:
```
Within the virtual environment, you can use the command python instead of python3, and pip instead of pip3 if you would prefer. If you use Python 3 on your machine outside of an environment, you will need to use the python3 and pip3 commands exclusively.
 ```

## Requirement.txt file
Save all the packages in the file with 
```bash
pip freeze > requirements.txt
```
 Keep in mind that in this case, requirements.txt file will list all packages that have been installed in a virtual environment, regardless of where they came from

## Install project dependencies
When if you’re going to share the project with the rest of the world you will need to install dependencies by running 
```bash
pip install -r requirements.txt
```
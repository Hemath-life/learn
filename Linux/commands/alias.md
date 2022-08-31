`alias` command, used to create a shortcut to another command

It’s common to always run a program with a set of options you like using.

For example, take the ls command. By default it prints very little information:

You can create a new command, for example I like to call it ll, that is an alias to ls -al

You do it in this way:

```bash
alias ll='ls -al'
```

Once you do, you can call `ll` just like it was a regular UNIX command:

The alias will work until the terminal session is closed.

To make it permanent, you need to add it to the shell configuration, which could be ~/.bashrc or ~/.profile or ~/.bash_profile if you use the Bash shell, depending on the use case.

Be careful with quotes if you have variables in the command: using double quotes the variable is resolved at definition time, using single quotes it’s resolved at invocation time. Those 2 are different:

```bash
alias lsthis="ls $PWD"
alias lscurrent='ls $PWD'
```

to get input form terminal

```bash
alias run="npm run start:dev"

dev(){
    npx nodemon $1
}
```

```bash
# to run those alias in terminal
dev app.js
run
nodemon
```

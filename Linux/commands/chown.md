`chown` command, used to change the owner of a file

Every file/directory in an Operating System like Linux or macOS (and every UNIX systems in general) has an owner

The owner (and the root user) can change the owner to another user, too, using the chown command:

```bash
chown <owner> <file>
```

Like this:

```bash
chown hemath test.txt
```

It’s rather common to have the need to change the ownership of a directory, and recursively all the files contained, plus all the subdirectories and the files contained in them, too.

You can do so using the -R flag:

```bash
chown -R <owner> <file>
```

Files/directories don’t just have an owner, they also have a group. Through this command you can change that simultaneously while you change the owner:

```bash
chown <owner>:<group> <file>
# Example:
chown flavio:users test.txt
```

You can also just change the group of a file using the chgrp command:

```bash
chgrp <group> <filename>
```

```
The `chown` command works on Linux, macOS, WSL, and anywhere you have a UNIX environmentalias nodemon="npx nodemon index.js"alias nodemon="npx nodemon index.js"alias nodemon="npx nodemon index.js"alias nodemon="npx nodemon index.js"alias nodemon="npx nodemon index.js"
```

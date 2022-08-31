- [Services running in Ubuntu](#see-all-services-running-in-ubuntu)
- [Grep](#global-regular-expression-grep)

<!---==========================================================================================-->

### See all services running in Ubuntu

```bash
service --status-all
# On the list the + indicates the service is running,
# indicates service is not running
# indicates the service state cannot be determined.

sudo service --status-all | grep postgres (or) grep post*
#  You can get list of all services and select by color one of them with 'grep':

sudo service postgresql status
#   Maybe what you want is the ps command
#   ps -ef
#   will show you all processes running. Then if you have an idea of what you're looking for use grep to filter;
ps -ef | grep postgres
```

<!---==========================================================================================-->

### Global regular expression (grep)

Syntax:

```bash
    # grep 'word' filename

    # fgrep 'word-to-search' file.txt

    # grep 'word' file1 file2 file3

    # grep 'string1 string2'  filename

    # cat otherfile | grep 'something'

    # command | grep 'something'

    # command option1 | grep 'data'

    # grep --color 'data' fileName

    # grep [-options] pattern filename

    # fgrep [-options] words file
```

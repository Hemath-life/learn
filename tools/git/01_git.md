`(fname-- represent file name)`

### How the git work?

- Version Control system- which track each and every change.
- if problem in one Version we'll back come to old version.
- compare previous code to present code.
- git comes with own terminally and language.

### Stages 4 type.

- Working directory - folder or dir when we initialize the git.
- Staging area - git add to push working directory file to staging area.intermediated staging area
  - you may not want add all file. some file you'll add. some file you may not add . here we can choose.
- Local Repository - here gonna save all commit file.
  - collect all changes.
  - when you want to compare the git Working dir and Local Repository gonna compare.
  - if we messed up the working directory no problem using git -checkout-
  - command help us to get last commit file on working directory.
- Remote repository - github,bitbucket,gitlab

3. Version Control using Git and the command line.

```bash
git init
# initialize the git
git status
# to see staging area which file are add and not added.

git add fname
# to add staging area file , (.) represent all file.
git commit -m 'commit message'
# helps you keep track what change you made in these code.
# message could be present better.
git logs
# we can see what commit we made so far.
git diff fname
#  it compare working dir and last committed file.to know which changes we made.
git checkout fname
# to roll back the file last version on current working directory (or)restored the old version.
```

5. What is GitHub?

- GitHub is Repository to store git file.

6. Local repo to github repository.

```bash
git remote add origin URL
#  I created remote Repository some where. origin is simply name of you
# remote and we can change the name.
#  remote is created

git branch -M master
git push -u origin master
#  __ these line push our Local Repository to remote Repository
#  origin __ name of the remote.
#  master __ name of the branch default or main branch.
#  -u __ pushing local to remote  repository using u flag.

```

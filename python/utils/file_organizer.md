```python
import os
import shutil

path = './'
files = os.listdir(path)
files.remove('main.py')
print(files)

for file in files:
    filename, extension = os.path.splitext(file) # splits filename and extension
    extension = extension[1:]

    if os.path.exists(path+'/'+extension): # if dir existing then move file to there
        shutil.move(path+'/'+file, path+'/'+extension+'/'+file)
    else:
        os.makedirs(path+'/'+extension) # create the dir and move file to there
        shutil.move(path+'/'+file, path+'/'+extension+'/'+file)
```

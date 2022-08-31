## Create cron job in linux:
    syntax: 
        - expression bin executionFile
        - example:
            '* * * * * python3 main.py >> outputFile.txt'

    step1: crontab -e
    step2: '* * * * * python3 main.py >> outputFile.txt'
    step3: 'esc:wq'

    example:
        # step1: "* * * * * " "/home/bitcot/Music/py/main.sh","/home/bitcot/Videos/main.sh" ,"/home/bitcot/Downloads/main.sh"


## linux and mac:
    - crontab -l 
        - list the cron jobs in the system 

## Crontab:
    tag:
        - l  list 
        - e  edit mode 


## reference:
<a href="https://crontab.guru/"> https://crontab.guru/</a>
<!---==========================================================================================-->

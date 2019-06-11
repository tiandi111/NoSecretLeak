# NoSecretLeak
How many times have you pushed your code to github but found that there's personal information, secret key, password and so on
that should not be exposed to public? You have to delete them from your code and push again or even worse, somebody stolen it before
you found out!

## Here's the solution!
NoSecretLeak is the safeguard that helps you discover your secrets in your files that you don't want to show to others.

## So how do I use it?
### Step1: Install Go 
    Make sure you installed go and build the workspace and GOPATH correctly.
### Step2: Download NoSecretLeak
    go get github.com/tiandi111/NoSecretLeak
### Step3: Install NoSecretLeak
    cd ~/go/src/github.com/tiandi111/NoSecretLeak
    go install
## Now you are ready to scan your files!
### Step1: Write your secret list
    secret,password,weight,
An example of secret list is shown above. The comma is used to separate each secret, you can also use other separator as you like.
**Ps: The only format that is supported by the current version is .txt file.**
### Step2: Run NoSecretLeak
    cd ~/dir  // cd to the directory you want to scan
    NoSecretLeak -d -s=secret -sep=, // if -d is set, secret file will be deleted after scanning
Write secret path after -s flag and your separator after -sep flag.
### Step3: Get your report 
    Warning! Secrets found!
    Secret Report:
    File path   |   Position   |   Secret
    /Users/tiandi/go/src/github.com/tiandi111/NoSecretLeak/main.go | 70:22 | string
    /Users/tiandi/go/src/github.com/tiandi111/NoSecretLeak/secret | 1:1 | string
## Or you can make it a git plugin using the bash script attached
The bash script is named as git-safepush which can be recognized by git. 
### Step1: Download and put the bash script in your envrioment
I recommend you to put the script in ~/go/bin. You can also put it anywhere as long as it is set to enviroment variable.
### Step2: Customized your script
Inside the script, you will see this line where you can customized your NoSecretLeak command.

    # Set file path for secret file and separator
    # if -d is set, secret file will be deleted after scanning
    NoSecretLeak -s=/Users/tiandi/secret -sep=,
    
### Step3: Push your code by 'git safepush'
    git safepush origin master
## Erase your secrets and showcase your work!
After scanning, NoSecretLeak automatically delete secret list (if it doesn't, you will get warnings on your terminal). So don't worry and feel free to exhibit your work to the world!
    :)

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
### Step4: Write your secret list
    secret,password,weight,
An example of secret list is shown above. The comma is used to separate each secret, you can also use other separator as you like.
**Ps: The only format that is supported by the current version is .txt file.**
### Step5: Run NoSecretLeak
    cd ~/dir  // cd to the directory you want to scan
    NoSecretLeak -s=secret -sep=,
Write secret path after -s flag and your separator after -sep flag.
### Step6: Get your report 
    Warning! Secrets found!
    Secret Report:
    File path   |   Position   |   Secret
    /Users/tiandi/go/src/github.com/tiandi111/NoSecretLeak/main.go | 70:22 | string
    /Users/tiandi/go/src/github.com/tiandi111/NoSecretLeak/secret | 1:1 | string
### Step7: Erase your secrets and showcase your work!
After scanning, NoSecretLeak automatically delete secret list (if it doesn't, you will get warnings on your terminal). So don't worry and feel free to exhibit your work to the world!
    :)

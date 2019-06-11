Language: [English](#NoSecretLeak), [中文](#NoSecretLeak防止秘密泄露)
# NoSecretLeak防止秘密泄露
这是一个基于Go语言的关键字查找小工具。在使用GitHub等开源代码托管平台时，开发者可能因疏忽误将个人信息，密钥，密码等关键信息上传到平台上，造成自身及公司利益受损。NoSecretLeak帮助使用者检测和定位这些隐私信息，预防隐私泄露、秘密泄露等危险情况。该工具还提供一个bash脚本，使用者可以配置该脚本并使用一个新的git命令'git safepush'，该命令在执行时会先运行NoSecretLeak对待上传代码进行检测，若有关键信息被检测出，则阻止此次push直到没有关键信息存在于代码中。

## 安装
### 步骤1: 安装GO
    Make sure you installed go and build the workspace and GOPATH correctly.
### 步骤2: 下载NoSecretLeak
    go get github.com/tiandi111/NoSecretLeak
### 步骤3: 安装NoSecretLeak
    cd ~/go/src/github.com/tiandi111/NoSecretLeak
    go install
    
## 使用
### 步骤一: 生成wordlist
    vim secret
    
    secret,password,weight,
上述是wordlist的一个范例，使用了逗号作为分隔符，具体使用时可以通过设置-sep标志位自定义分隔符。
### 步骤二：运行NoSecretLeak
    cd ~/dir  // cd to the directory you want to scan
    NoSecretLeak -d -s=secret -sep=, // if -d is set, secret file will be deleted after scanning
-s标志位后填写关键信息wordlist的路径

-d标志位用来设置是否在检测完成后自动删除wordlist文件

-sep标志位用来自定义分隔符
### 步骤三：输出报告
    Warning! Secrets found!
    Secret Report:
    File path   |   Position   |   Secret
    /Users/tiandi/go/src/github.com/tiandi111/NoSecretLeak/main.go | 70:22 | string
    /Users/tiandi/go/src/github.com/tiandi111/NoSecretLeak/secret | 1:1 | string

## 作为GIT插件使用
项目中命名为git-safepush的脚本可作为git插件使用。
### 步骤一：下载插件，设置环境变量，赋予权限
建议直接将其放在GO的Workspace下的bin文件中（~/go/bin），省去设置环境变量的步骤。
进入脚本所在文件夹，使用如下命令设置权限：

    $ chmod u+x safe-push
### 步骤二：自定义脚本
在脚本中有如下代码，可以通过编辑这行代码设置运行NoSecretLeak的命令。

    # Set file path for secret file and separator
    # if -d is set, secret file will be deleted after scanning
    NoSecretLeak -s=/Users/tiandi/secret -sep=,
    
### 步骤三：使用‘git safepush’命令
使用该命令时，除了在push前会检测关键信息外，与‘git push’命令没有任何区别
    $ git safepush origin master    
如果有关键信息被检测到，则会看到如下输出：

    Warning! Secrets found!
    Secret Report: 
    File path   |   Position   |   Secret
    /Users/tiandi/go/src/github.com/tiandi111/NoSecretLeak/.git/hooks/fsmonitor-watchman.sample | 12:18 | string
    /Users/tiandi/go/src/github.com/tiandi111/NoSecretLeak/.git/logs/HEAD | 5:172 | string
    /Users/tiandi/go/src/github.com/tiandi111/NoSecretLeak/.git/logs/refs/heads/master | 5:172 | string
    Secret found in your code, 'git push' is aborted!

可以看到因为有关键信息被检测出来，此次push被取消。

如果没有关键信息被检测到，会看到如下输出：
    
    No secret found, your code is safe to release!
    Everything up-to-date


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

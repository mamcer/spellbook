# jenkins

## home dir

```
/var/lib/jenkins
```

## docker

```bash
docker pull jenkins/jenkins:lts
docker run --name jenkins -d -v jenkins_home:/home/mario/jenkins_home -p 8080:8080 -p 50000:50000 --restart=on-failure jenkins/jenkins:lts
docker exec -it jenkins /bin/bash
cat /var/jenkins_home/secrets/initialAdminPassword
```

## bare metal

follow the instructions

[https://www.jenkins.io/doc/book/installing/linux/#debianubuntu](https://www.jenkins.io/doc/book/installing/linux/#debianubuntu)

start as a service

```bash
sudo systemctl enable jenkins
sudo systemctl start jenkins
sudo systemctl status jenkins
```

[https://www.jenkins.io/doc/book/pipeline/docker/](https://www.jenkins.io/doc/book/pipeline/docker/)  

## after installation

install blue ocean plugin!

configure timezone: [user] > Account > Time zone (America/Argentina/Buenos_Aires)

## configure github ssh access

setup ssh on build server

copy keys to jenkins directory

```bash
sudo cp ~/.ssh/* /var/lib/jenkins/.ssh
sudo chown -R jenkins:jenkins /var/lib/jenkins/.ssh/
```

go to Manage Jenkins > Credentials  
add SSH Username with private key  
configure a username and add Private Key (enter directly)

## configure jenkins email

Manage Jenkins >  System  
E-mail Notification  
SMTP server: smtp.gmail.com  
Use SMTP Authentication  
User Name: ci.mamcer@gmail.com  
Password: [app-password]   
Use SSL  
SMTP port: 465  
> since 2022 yo need to create an app password for this to work [app_password](https://support.google.com/mail/answer/185833?hl=en)
> Test configuration by sending test e-mail
> Fix the "address not configured yet<ci.mamcer@gmail.com>" issue
> go to Manage Jenkins > Jenkins Location section and specify a valid 'System Admin e-mail address'

## configure golang

install Go Plugin (1.4 at the moment of this write)  

manage jenkins > tools > Go installations:  
define a version and add a [name] (the name defined here is the same that should be referenced on a pipeline)
    
```
tools {
    go '[name]'
} 
```

## jenkins file

Jenkinsfile

```yaml
pipeline {
    agent any
    tools {
        go '1.23.4'
    }
    stages {
        stage('build') {
            steps {
                sh 'go -C ./cmd/api build -o ../../bin main.go'
            }
        }
    }
    post {
        success {
            emailext(
                        to: '<email-address>',
                        subject: "SUCCESS: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]'",
                        body: """<p>Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' succeeded.</p><p>Check console output at <a href='${env.BUILD_URL}'>${env.BUILD_URL}</a></p>""",
                        mimeType: 'text/html'
                    )
        }
        failure {
            emailext(
                        to: '<email-address>',
                        subject: "FAILURE: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]'",
                        body: """<p>Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' failed.</p><p>Check console output at <a href='${env.BUILD_URL}'>${env.BUILD_URL}</a></p>""",
                        mimeType: 'text/html'
                    )
        }
        always {
            cleanWs()
        }
    }
}
```

## jenkins plan

Source Code Management

```
repository: https://github.com/mamcer/ufiles.git
branch: */develop
Build Triggers
Poll SCM
H/5 * * * *
Build 
[Execute shell]
    dotnet restore
[Execute shell]
    dotnet build --configuration release
[Execute shell]
    rm -fr test/UFiles.Core.Test/TestResults
    rm -fr test/UFiles.Util.Test/TestResults
[Execute shell]
    dotnet test --configuration release --logger trx --results-directory ./TestResults
Process xUnit test result report
report type: MSTest-Version N/A (default)
Pattern: **/TestResults/*.trx
Post-build Actions
[Editable Email Notification]
project recipient list: <email-address>
project reply-to list: $DEFAULT_REPLYTO
content-type: Default Content Type
default subject: $DEFAULT_SUBJECT
default content: $DEFAULT_CONTENT
```
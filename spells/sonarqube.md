# sonarqube

## install

server

[https://docs.sonarsource.com/sonarqube-community-build/try-out-sonarqube/](https://docs.sonarsource.com/sonarqube-community-build/try-out-sonarqube/)

```bash
/opt/sonarqube/bin/linux-x86-64/sonar.sh console
```

docker 

```bash
docker pull sonarqube
docker run -d --name sonarqube -p 9000:9000 [image-name]
```

login as `admin/admin` and create a project

[https://docs.sonarqube.org/latest/analysis/scan/sonarscanner/](https://docs.sonarqube.org/latest/analysis/scan/sonarscanner/)

scanner

```bash
dotnet tool install --global dotnet-sonarscanner
```

dotnet tool install --global dotnet-reportgenerator-globaltool


## run dotnet scanner
    
```bash
dotnet sonarscanner begin /k:"nNostalgia" /d:sonar.host.url="http://192.168.100.100:9000" /d:sonar.cs.dotcover.reportsPaths=coverage.html  /d:sonar.token="$SONAR_TOKEN"   

dotnet build --no-incremental

~/.dotnet/tools/dotCover cover-dotnet --output=coverage.html --reporttype=HTML --filters="-:*.Test;-:module=*;class=*.Migrations.*;-:Spectre.Console.Cli;-:testhost" -- test Nostalgia.sln

dotnet sonarscanner end /d:sonar.token="$SONAR_TOKEN"
```

## bash script .net 8 solution

```bash
#!/bin/bash

SONAR_PROJECT_KEY="[project-key]"
SONAR_HOST="http://localhost:9000"
SONAR_TOKEN="[token]"
TEST_PROJECT="[solution-path]"

dotnet sonarscanner begin /k:"$SONAR_PROJECT_KEY" /d:sonar.host.url="$SONAR_HOST" /d:sonar.token="$SONAR_TOKEN" /d:sonar.coverageReportPaths="coverage/SonarQube.xml"

dotnet build 

dotnet test $TEST_PROJECT --collect:"XPlat Code Coverage" --results-directory ./coverage

reportgenerator "-reports:./coverage/*/coverage.cobertura.xml" "-targetdir:coverage" "-reporttypes:SonarQube"

dotnet sonarscanner end /d:sonar.token="$SONAR_TOKEN"
```

## golang (might be obsolete and/or deprecated... this was long time ago in a galaxy far away)

## sonarqube project properties

`sonar-project.properties`
    
```
sonar.projectKey=[project-key]
sonar.projectName=[project-name]

sonar.sources=.
sonar.exclusions=**/*_test.go,**/vendor/**

sonar.tests=.
sonar.test.inclusions=**/*_test.go
sonar.test.exclusions=**/vendor/**
```

## sonarqube scanner

download sonar scanner

```bash
    # by default the sonar-scanner configuration is configured to localhost:9000
    sonar-scanner-4.2.0.1873-macosx/bin/sonar-scanner
```

run sonar scanner

```bash
~/Downloads/sonar-scanner-4.2.0.1873-macosx/bin/sonar-scanner -Dproject.settings=sonar-project.properties
```

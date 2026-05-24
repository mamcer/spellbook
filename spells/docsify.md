# docsify

[https://docsify.js.org/#/quickstart](https://docsify.js.org/#/quickstart)

## install 

```bash
npm install -g docsify-cli@latest
```

## init

```bash
docsify init ./docs
```

you should see /docs:

```
index.html
README.md
.nojekyll
```

## preview

```bash
docsify serve docs
```
## docker

```bash
vim Dockerfile
```

```docker
FROM node:latest
LABEL description="A demo Dockerfile for build Docsify."
WORKDIR /docs
RUN npm install -g docsify-cli@latest
EXPOSE 3000/tcp
ENTRYPOINT docsify serve .
```
```bash
docker build -f Dockerfile -t spellbook .
```

```bash
docker run -d -p 3000:3000 --name=spellbookc -v $(pwd):/docs --rm spellbook
```

## optionally connect to container

```bash
docker container exec -it spellbookc /bin/bash
```
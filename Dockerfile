FROM node:latest
LABEL description="spellbook"
WORKDIR /docs
RUN npm install -g docsify-cli@4.4.4
COPY . /docs
EXPOSE 3000/tcp
ENTRYPOINT ["docsify", "serve", "."]
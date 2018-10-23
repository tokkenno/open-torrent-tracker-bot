FROM node:10

COPY ./ /usr/src/app
WORKDIR /usr/src/app

RUN npm install

CMD node src/index.js
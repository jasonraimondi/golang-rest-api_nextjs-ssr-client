FROM node:alpine
WORKDIR /app

COPY *.json /app/
COPY .babelrc /app/
RUN npm install

COPY . /app/

RUN npm run build
CMD ["npm", "run", "start"]

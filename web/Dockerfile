FROM node:alpine
WORKDIR /app

COPY *.json /app/
COPY .babelrc /app/
RUN npm install

COPY components /app/components/
COPY lib /app/lib/
COPY pages /app/pages/

RUN npm run build
CMD ["npm", "run", "start"]

FROM node:alpine
WORKDIR /app
COPY package* /app/
RUN npm install
COPY tsconfig.json /app/
COPY components /app/components/
COPY lib /app/lib/
COPY pages /app/pages/
COPY .babelrc /app/

RUN npm run build
CMD ["npm", "run", "start"]

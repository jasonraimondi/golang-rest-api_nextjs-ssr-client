FROM node:12-alpine
WORKDIR /app
COPY . /app/
RUN npm install && npm run build && npm prune --production
CMD ["npm", "run", "start"]

FROM node:14.16.1

RUN mkdir -p /frontend
COPY . /frontend

WORKDIR /frontend
RUN npm install

EXPOSE 3000

CMD [ "npm", "start" ]
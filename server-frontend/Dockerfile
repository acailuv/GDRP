FROM node:14.16.1

WORKDIR /server-frontend

RUN npm install -g localtunnel

CMD lt -h ${WEB_HOST} -p 3000 --subdomain ${WEB_SUBDOMAIN} --local-host frontend
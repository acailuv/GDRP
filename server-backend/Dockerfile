FROM node:14.16.1

WORKDIR /server-backend
RUN mkdir output

RUN npm install -g localtunnel

CMD lt -h ${API_HOST} -p 5000 --subdomain ${API_SUBDOMAIN} --local-host api | tee ./output/backend-link.txt
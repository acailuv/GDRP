FROM node:14.16.1

RUN mkdir -p /frontend
COPY . /frontend

WORKDIR /frontend
RUN npm install

EXPOSE 3000

CMD if [ ${ENV} = production ];         \
	then                                \
        npm install -g http-server &&   \
        npm run build &&                \
        cd build &&                     \
        hs -p 3000;                     \
	else                                \
	    npm run start;                  \
	fi
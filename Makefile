start:
	@if [ -d "./frontend/node_modules" ]; then\
		docker-compose up;\
	else\
		cd ./frontend;\
		npm install;\
		cd ../;\
		docker-compose up --build;\
	fi

generate-migration:
	@migrate create -ext sql -dir ./backend/database/migrations -seq ${name}

run-migration-up:
	@migrate -database ${database} -path ./backend/database/migrations up

run-migration-down:
	@migrate -database ${database} -path ./backend/database/migrations down
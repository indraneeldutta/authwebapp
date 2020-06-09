deps:
	go get -u github.com/gorilla/mux
	go get -u github.com/joho/godotenv
	go get github.com/markbates/goth
	go get github.com/gorilla/pat
	go get -u github.com/gorilla/sessions
	go get -u github.com/lib/pq
	
build:
	go build -o bin/main

run:
	./bin/main
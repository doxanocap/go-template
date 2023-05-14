run:
	go run cmd/main.go
up:
	go run cmd/main.go up
down:
	go run cmd/main.go down
redo: 
	make down && make up
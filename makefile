api:
	go get
	xdg-open http://localhost:8080/v1/doctors
	go run main.go

test:
	go get
	cd tests && go test
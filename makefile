build:
	cd cmd &&go build -o ../task-tracker main.go

test:
	go test -p=1 -count=1 ./commands ./models
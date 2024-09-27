build:
	cd cmd && go build -o ../tasker main.go

test:
	go test -p=1 -count=1 ./commands ./models
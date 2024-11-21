build:
	cd api && go build -o temperature_logger .
	cd api && env GOOS=windows GOARCH=amd64 go build -o temperature_logger.exe .
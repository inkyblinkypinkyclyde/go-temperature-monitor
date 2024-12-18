build:
	cd api && go build -o temperature_logger_mac_arm .
	cd api && env GOOS=windows GOARCH=amd64 go build -o temperature_logger_win_x64.exe .
	cd api && env GOOS=linux GOARCH=arm go build -o temperature_logger_linux_arm .
	cd api && env GOOS=linux GOARCH=amd64 go build -o temperature_logger_linux_x64 .
	cd api && env GOOS=darwin GOARCH=amd64 go build -o temperature_logger_mac_intel .

test:
	cd api && go test -v ./...
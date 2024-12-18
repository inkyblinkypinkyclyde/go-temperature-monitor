build-app:
	cd app && go build -o ../build/temperature_logger_mac_arm .
	cd app && env GOOS=windows GOARCH=amd64 go build -o ../build/temperature_logger_win_x64.exe .
	cd app && env GOOS=linux GOARCH=arm go build -o ../build/temperature_logger_linux_arm .
	cd app && env GOOS=linux GOARCH=amd64 go build -o ../build/temperature_logger_linux_x64 .
	cd app && env GOOS=darwin GOARCH=amd64 go build -o ../build/temperature_logger_mac_intel .

test:
	cd api && go test -v ./...
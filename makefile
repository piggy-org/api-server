all:
	GOOS=linux GOARCH=amd64 go build main.go -o piggy-market
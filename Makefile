build: clean
	GOOS=linux GOARCH=amd64 go build

clean:
	rm -f portplz
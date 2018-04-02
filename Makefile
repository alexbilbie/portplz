build: clean
	GOOS=linux GOARCH=amd64 go build

test:
	docker build -t portplz .
	docker run --rm portplz 443

clean:
	rm -f portplz
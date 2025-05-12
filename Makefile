install:
	(cd pdf-generator && go mod download && go mod tidy)

dev:
	(cd pdf-generator && go run tests/tests.go)

build:
	cd pdf-generator/cmd && \
	GOARCH=arm64 GOOS=linux go build -o bootstrap main.go && \
	zip bootstrap.zip bootstrap && \
	rm bootstrap && \
	echo "Build completed successfully. Created bootstrap.zip"

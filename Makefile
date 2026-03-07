install:
	(cd pdf-generator && go mod download && go mod tidy)

dev-label:
	(cd pdf-generator && go run tests/label/label.go)

dev-carta:
	(cd pdf-generator && go run tests/carta/carta.go)

build:
	cd pdf-generator/cmd && \
	GOARCH=arm64 GOOS=linux go build -o bootstrap main.go && \
	zip bootstrap.zip bootstrap && \
	rm bootstrap && \
	echo "Build completed successfully. Created bootstrap.zip"

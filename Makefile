API_PACKAGES=./...

build:
	go build ${API_PACKAGES}

unit:
	go test ${API_PACKAGES}

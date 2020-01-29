coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

guard-%:
	@ if [ "${${*}}" = "" ]; then \
		echo "Environment variable $* not set"; \
		exit 1; \
	fi

test:
	go test ./... -cover
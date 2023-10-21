build:
	@ go build -o bin/bolt cmd/main.go

# delete previous build
remove:
	@ rm bin/bolt
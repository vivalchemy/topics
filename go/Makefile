PROJECT_NAME=fun_with_go

dev: ./main.go
	@air

build: ./main.go
	@ echo "Building app.."
	@ go build -o ./bin/${PROJECT_NAME}
	@ echo "App is built"

preview: ./bin/${PROJECT_NAME}
	./bin/${PROJECT_NAME}

clean: ./bin/
	rm -f ./bin/*


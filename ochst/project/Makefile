TEST_BINARY=ochstTest
TEST_DEPLOY=testadmin@172.16.145.128:/home/testadmin/testoch/

build_test:
	@echo "Building test binary"
	cd ../ && GOARCH=arm64 GOOS=linux go build -o ${TEST_BINARY} ./main.go
	@echo "Built ${TEST_BINARY}"
deploy_test:
	@echo "Delpoying test binary"
	cd ../ && scp ./${TEST_BINARY} ${TEST_DEPLOY}
	@echo "Delpoyed"
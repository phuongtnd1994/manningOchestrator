TEST_WORKER=ochstTest
TEST_DEPLOY_WORKER=testadmin@172.16.145.128:/home/testadmin/testoch/
TEST_MANAGER=ochstManager
TEST_DEPLOY_MANAGER=testadmin@172.16.145.128:/home/testadmin/testoch/

build_test:
	@echo "Building worker binary"
	cd ../ && GOARCH=arm64 GOOS=linux go build -o ${TEST_WORKER} ./main.go
	@echo "Built ${TEST_WORKER}"
deploy_test:
	@echo "Delpoying worker binary"
	cd ../ && scp ./${TEST_WORKER} ${TEST_DEPLOY_WORKER}
	@echo "Delpoyed"
	
build_deploy_test: build_test deploy_test

build_manager:
	@echo "Building manager binary"
	cd ../ && GOARCH=arm64 GOOS=linux go build -o ${TEST_MANAGER} ./main.go
	@echo "Built ${TEST_MANAGER}"

deploy_manager_test:
	@echo "Delpoying manager binary"
	cd ../ && scp ./${TEST_MANAGER} ${TEST_DEPLOY_MANAGER}
	@echo "Delpoyed"

build_deploy_manager: build_manager deploy_manager_test

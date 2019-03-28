#### Project generation setup
PROJECT_NAME=github.com/oracle/oci-go-sdk
PROJECT_PATH=$(GOPATH)/src/$(PROJECT_NAME)
REMOVE_AFTER_GENERATE=audit/audit_waiters.go objectstorage/objectstorage_waiters.go

#### Versions
#### If you are doing a release, do not forget to increment this versions
VER_MAJOR=1
## where the number is the week number
VER_MINOR=25
VER_TAG=preview
###################

##### Integ test setup
INTEGTEST_DIR = integtest
TEST_HELPERS = test_helpers.go test_service_deps_helpers.go
TEST_FILES = $(notdir $(wildcard $(INTEGTEST_DIR)/*_integ_test.go))
TARGETS = $(patsubst %_client_integ_test.go, %, $(TEST_FILES))
TEST_TARGETS = $(patsubst %_client_integ_test.go, test-%, $(TEST_FILES))

#### AutoTest setup
AUTOTEST_DIR = autotest
AUTOTEST_HELPERS = client.go configurations.go main_test.go
AUTOTEST_FILES = $(notdir $(wildcard $(AUTOTEST_DIR)/*client_auto_test.go))
AUTOTEST_TARGETS = $(patsubst %_client_auto_test.go, autotest-%, $(AUTOTEST_FILES))

define HELP_MESSAGE
make -f MakefileDevelopment.mk build: builds and generate the sdk
make -f MakefileDevelopment.mk build-sdk: to build the sdk only
make -f MakefileDevelopment.mk generate: to generate the sdk
make -f MakefileDevelopment.mk list-test-services to list all integ tests
make -f MakefileDevelopment.mk test-[name of the package] to integ test a package
make -f MakefileDevelopment.mk test-all runs all integ tests
make -f MakefileDevelopment.mk list-autotest to list all autogenerated test targets at the service level
make -f MakefileDevelopment.mk autotest-[name of the package] to run autogenerated tests
make -f MakefileDevelopment.mk autotest-all runs all autogenerated tests
endef


.PHONY: help

export HELP_MESSAGE
help:
	@echo "$$HELP_MESSAGE"

list-test:
	@echo $(TEST_TARGETS)

list-autotest-services:
	@echo $(AUTOTEST_TARGETS)

test-all: build-sdk build-autotest test-sdk-only $(TEST_TARGETS)

autotest-all: build-sdk test-sdk-only $(AUTOTEST_TARGETS)

autotest: build-autotest
	go test -v -run $(TEST_NAME) -count 1 -timeout 3h github.com/oracle/oci-go-sdk/autotest


$(TEST_TARGETS): test-%:%
	@echo Testing $(INTEGTEST_DIR)/$<_client_integ_test.go
	@(cd $(INTEGTEST_DIR) && go test -v $(TEST_HELPERS) $<_client_integ_test.go)

$(AUTOTEST_TARGETS): autotest-%:%
	@echo Testing $(AUTOTEST_DIR)/$<_client_auto_test.go
	@(cd $(AUTOTEST_DIR) && go test -v $(AUTOTEST_HELPERS) $<_client_auto_test.go)


$(TARGETS): %:integtest/%_client_integ_test.go

generate:
	@echo "Cleaning and generating sdk"
	@(cd $(PROJECT_PATH) && make clean-generate)
	@echo "Cleaning autotest files"
	@find autotest -name \*_auto_test.go|xargs rm -f
	PROJECT_NAME=$(PROJECT_NAME) mvn clean install
	@(cd $(PROJECT_PATH) && rm -f $(REMOVE_AFTER_GENERATE))

build-autotest:
	@echo "building autotests"
	@(cd $(AUTOTEST_DIR) && gofmt -s -w . && go test -c)

build-sdk:
	@echo "Building sdk"
	@(cd $(PROJECT_PATH) && make build)

test-sdk-only:
	@echo "Testing sdk common"
	@(cd $(PROJECT_PATH) && make test)



release-sdk:
	@echo "Building oci-go-sdk with major:$(VER_MAJOR) minor:$(VER_MINOR) patch:$(VER_PATCH) tag:$(VER_TAG)"
	@(cd $(PROJECT_PATH) && VER_MAJOR=$(VER_MAJOR) VER_MINOR=$(VER_MINOR) VER_PATCH=$(VER_PATCH) VER_TAG=$(VER_TAG) make release)

build: generate build-sdk build-autotest
	@(cd $(PROJECT_PATH) && make pre-doc)

release: generate release-sdk build-autotest

generate-pipeline: build-sdk build-autotest test-sdk-only

#### Project generation setup
PROJECT_NAME=github.com/oracle/oci-go-sdk
PROJECT_PATH=$(GOPATH)/src/$(PROJECT_NAME)
REMOVE_AFTER_GENERATE=audit/audit_waiters.go objectstorage/objectstorage_waiters.go
DOC_SERVER_URL_DEV=https:\/\/docs.cloud.oracle.com

#### Versions
#### If you are doing a release, do not forget to increment this versions
VER_MAJOR=27
## where the number is the week number
VER_MINOR=2
VER_TAG=preview
###################

#### AutoTest setup
AUTOTEST_DIR = autotest
AUTOTEST_HELPERS = client.go configurations.go main_test.go
AUTOTEST_FILES = $(notdir $(wildcard $(AUTOTEST_DIR)/*client_auto_test.go))
AUTOTEST_TARGETS = $(patsubst %_client_auto_test.go, autotest-%, $(AUTOTEST_FILES))

define HELP_MESSAGE
make -f MakefileDevelopment.mk build: builds and generate the sdk
make -f MakefileDevelopment.mk build-sdk: to build the sdk only
make -f MakefileDevelopment.mk generate: to generate the sdk
make -f MakefileDevelopment.mk list-autotest-services to list all autogenerated test targets at the service level
make -f MakefileDevelopment.mk autotest-[name of the package] to run autogenerated tests
make -f MakefileDevelopment.mk autotest-all runs all autogenerated tests
endef


.PHONY: help

export HELP_MESSAGE
help:
	@echo "$$HELP_MESSAGE"

list-autotest-services:
	@echo $(AUTOTEST_TARGETS)

test-all: build-sdk build-autotest test-sdk-only test-integ-test

autotest-all: build-sdk test-sdk-only $(AUTOTEST_TARGETS)

autotest: build-autotest
	go test -v -run $(TEST_NAME) -count 1 -timeout 3h github.com/oracle/oci-go-sdk/autotest

$(AUTOTEST_TARGETS): autotest-%:%
	@echo Testing $(AUTOTEST_DIR)/$<_client_auto_test.go
	@(cd $(AUTOTEST_DIR) && go test -v $(AUTOTEST_HELPERS) $<_client_auto_test.go)

generate:
	@echo "Cleaning and generating sdk"
	@(cd $(PROJECT_PATH) && make clean-generate)
	@echo "Cleaning autotest files"
	@find autotest -name \*_auto_test.go|xargs rm -f
	PROJECT_NAME=$(PROJECT_NAME) mvn clean install
	@(cd $(PROJECT_PATH) && rm -f $(REMOVE_AFTER_GENERATE))
	find . -name \*.go |xargs sed -i "" "s#\"$(PROJECT_NAME)/\(v[0-9]*/\)*#\"$(PROJECT_NAME)/v$(VER_MAJOR)/#g"

build-autotest:
	@echo "building autotests"
	@(cd $(AUTOTEST_DIR) && gofmt -s -w . && go test -c)

build-sdk:
	@echo "Building sdk"
	@(cd $(PROJECT_PATH) && make build)

test-sdk-only:
	@echo "Testing sdk common"
	@(cd $(PROJECT_PATH) && make test)

test-integ-test:
	@echo "Testing sdk integ test"
	@(cd $(PROJECT_PATH) && make test-integ)

release-sdk:
	@echo "Building oci-go-sdk with major:$(VER_MAJOR) minor:$(VER_MINOR) patch:$(VER_PATCH) tag:$(VER_TAG)"
	@(cd $(PROJECT_PATH))
	find . -name \*.go |xargs sed -i "" "s#\"$(PROJECT_NAME)/\(v[0-9]*/\)*#\"$(PROJECT_NAME)/v$(VER_MAJOR)/#g"
	@(VER_MAJOR=$(VER_MAJOR) VER_MINOR=$(VER_MINOR) VER_PATCH=$(VER_PATCH) VER_TAG=$(VER_TAG) make release)

build: generate build-sdk build-autotest
	@(cd $(PROJECT_PATH) && make pre-doc)

release: generate release-sdk build-autotest

# command used by self service pipeline to clean all generated files
clean-pipeline:
	@echo "Cleaning generated files"
	@(cd $(PROJECT_PATH) && make clean-generate)
	@echo "Cleaning autotest files"
	@find autotest -name \*_auto_test.go|xargs rm -f
	@(cd $(PROJECT_PATH) && rm -f $(REMOVE_AFTER_GENERATE))

# doing build and lint for generated code in self-service pipeline
lint-pipeline: update-import build-autotest test-sdk-only
	@echo "Rendering doc server to ${DOC_SERVER_URL_DEV}"
	find . -name \*.go |xargs sed -i 's/{{DOC_SERVER_URL}}/${DOC_SERVER_URL_DEV}/g'
	find . -name \*.go |xargs sed -i 's/https:\/\/docs.us-phoenix-1.oraclecloud.com/${DOC_SERVER_URL_DEV}/g'

# update all imports to match latest major version
update-import:
	find . -name \*.go |xargs sed -i "s#\"$(PROJECT_NAME)/\(v[0-9]*/\)*#\"$(PROJECT_NAME)/v$(VER_MAJOR)/#g"
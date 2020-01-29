BRANCH = "master"
GO_BUILDER_IMAGE = "vidsyhq/go-builder"
PATH_BASE = "/go/src/github.com/revett"
REPONAME = "go-yt-comments"
TEST_PACKAGES = "./ytc"
VERSION = $(shell cat ./VERSION)

check-version:
	@echo "=> Checking if VERSION exists as Git tag..."
	(! git rev-list ${VERSION})

push-tag:
	git checkout ${BRANCH}
	git pull origin ${BRANCH}
	git tag ${VERSION}
	git push origin ${BRANCH} --tags

test:
	@API_KEY=${API_KEY} go test "${TEST_PACKAGES}" -cover
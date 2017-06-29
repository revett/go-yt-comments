BRANCH = "master"
GO_BUILDER_IMAGE = "vidsyhq/go-builder"
PATH_BASE = "/go/src/github.com/revett"
REPONAME = "go-yt-comments"
VERSION = $(shell cat ./VERSION)

check-version:
	@echo "=> Checking if VERSION exists as Git tag..."
	(! git rev-list ${VERSION})

install-ci:
	@docker run \
	-e BUILD=false \
	-v "${CURDIR}":${PATH_BASE}/${REPONAME} \
	-w ${PATH_BASE}/${REPONAME} \
	${GO_BUILDER_IMAGE}

push-tag:
	git checkout ${BRANCH}
	git pull origin ${BRANCH}
	git tag ${VERSION}
	git push origin ${BRANCH} --tags

test:
	@go test . -cover

test-ci:
	@docker run \
	-v "${CURDIR}":${PATH_BASE}/${REPONAME} \
	-w ${PATH_BASE}/${REPONAME} \
	--entrypoint=go \
	${GO_BUILDER_IMAGE} test . -cover

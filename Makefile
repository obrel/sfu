SHELL		= 	/bin/sh
APP_NAME	= 	obrel-sfu
VERSION		:=	$(shell git describe --always --tags)
GIT_COMMIT	=	$(shell git rev-parse HEAD)
GIT_DIRTY	=	$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE	=	$(shell date '+%Y-%m-%d %H:%M:%S')
CGO_ENABLED	=	0
GOARCH		=	amd64
GOOS		=	$(shell uname -s)

.PHONY: default
default: help

.PHONY: help
help:
	@echo 'Management commands for ${APP_NAME}'
	@echo
	@echo 'Usage:'
	@echo '	make lint           Run static linter.'
	@echo '	make test           Run tests.'
	@echo '	make sec            Run security checks.'
	@echo '	make cov            Run coverage detail'
	@echo '	make run ARGS=      Run with supplied arguments.'
	@echo '	make build          Compile application.'
	@echo '	make clean          Clean directory tree.'
	@echo '	make prepare        Prepare application before making PR.'
	@echo

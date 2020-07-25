# try refresh envvar from .env
ifneq (,$(wildcard ./.env))
	include .env
	export
endif

# aux commands

envvar-exists-%:
	@if [ -z '${${*}}' ]; then echo 'ERROR: variable $* not set' && exit 1; fi

cmd-exists-%:
	@hash $(*) > /dev/null 2>&1 || \
		(echo "ERROR: '$(*)' must be installed and available on your PATH."; exit 1)

.PHONY: envvar-exists-% cmd-exists-%

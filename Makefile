.PHONY: _ all compose run start stop
.SILENT: _ all compose run start stop

all:
	$(MAKE) -p _ | grep -P '^[^#.%_]\w*?:' | grep -iv make | sed 's/:.*//' | sort

compose_params ?= --compatibility
compose_command ?=
compose:
	docker-compose $(compose_params) $(compose_command)

cmd ?= up
params ?= --force-recreate --build
run:
	$(MAKE) compose compose_command="$(cmd) $(params)"

start:
	$(MAKE) run params='--build --force-recreate -d'

stop:
	$(MAKE) run cmd='down' params=''



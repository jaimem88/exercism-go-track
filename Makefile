# Use richgo for pretty test output if you have it installed
# https://github.com/kyoh86/richgo
# if you have richgo installed set this to 'richgo'
export RICHGO := $(if $(RICHGO),$(RICHGO),richgo)

run:
	go run ./cmd/exercism/main.go

test-one:
	$(RICHGO) test -v --bench --benchmem -count 1 ${ARGS}

test:
	$(RICHGO) test -v --bench --benchmem -count 1 ./...



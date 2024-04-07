
PWD = $(shell pwd)

# Location to install dependencies to
LOCALBIN ?= $(PWD)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

# Tool Binaries
SWAG ?= $(LOCALBIN)/swag

# Tool Versions
SWAG_VERSION ?= latest

# Install swag
.PHONY: swag
swag: $(SWAG)
$(SWAG): $(LOCALBIN)
	GOBIN=$(LOCALBIN) go install github.com/swaggo/swag/cmd/swag@$(SWAG_VERSION)

# Generate swagger documanetation
.PHONY: swag_init
swag_init: swag
	$(SWAG) init

# Format swag comments
.PHONY: swag_fmt
swag_fmt: swag
	$(SWAG) fmt

# Run all swagger targets
.PHONE: swag_gen
swag_gen: swag_init swag_fmt

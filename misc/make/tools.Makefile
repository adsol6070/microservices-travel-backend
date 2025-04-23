ifeq ($(ARCH),x86_64)
	ARCH := amd64
else ifeq ($(ARCH),aarch64)
	ARCH := arm64 
endif

define github_url
	https://github.com/$(GITHUB)/releases/download/v$(VERSION)/$(ARCHIVE)
endef

# creates a directory bin
bin:
	@ mkdir -p $@

# Tools
MIGRATE := $(shell command -v migrate || echo "bin/migrate")
migrate: bin/migrate

bin/migrate: VERSION := 4.18.2
bin/migrate: GITHUB  := golang-migrate/migrate
bin/migrate: ARCHIVE := migrate.$(OSTYPE)-$(ARCH).tar.gz
bin/migrate: bin
	@ printf "Installing migrate... "
	@ curl -Ls $(call github_url) -o bin/$(ARCHIVE)
	@ tar -zxf bin/$(ARCHIVE) -C bin
	@ chmod +x bin/migrate
	@ rm bin/$(ARCHIVE)  # Remove the tar.gz file after extraction
	@ ./bin/migrate --version
	@ echo "done."

AIR := $(shell command -v air || echo "bin/air")
air: bin/air ## Installs air (go file watcher)

bin/air: VERSION := 1.61.7
bin/air: GITHUB  := air-verse/air
bin/air: ARCHIVE  := air_$(VERSION)_$(OSTYPE)_$(ARCH).tar.gz
bin/air: bin
	@ printf "Installing air... "
	@ curl -Ls $(call github_url) -o bin/$(ARCHIVE)
	@ tar -zxf bin/$(ARCHIVE) -C bin
	@ chmod +x bin/air
	@ rm bin/$(ARCHIVE)  # Remove the tar.gz file after extraction
	@ ./bin/air -v
	@ echo "done."
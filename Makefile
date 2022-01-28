# Intro for beggining the paradocs install process
intro:
	$(info "Beginning paradocs installation process....")
	$(info "Post 'go install' exe will be added to path and can be called with 'paradocs'")
	@sleep 1

# Build go executable locally
build:
	@echo "Starting build process for paradocs...."
	@go build -o paradocs
	@chmod +x ./paradocs
	@mv ./paradocs /usr/local/bin
	@echo 'alias paradocs='paradocs'' >> ~/.bashrc
	@echo "Build process finished for paradocs...."

# Install go executable in path
install:
	@echo "Beginning install process for paradocs...."
	@go install
	@echo "Install process for paradocs finished...."

# Run all make commands
all: intro build install
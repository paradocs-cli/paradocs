intro:
	$(info "Beginning paradocs installation process....")
	$(info "Post 'go install' exe will be added to path and can be called with 'paradocs'")
	@sleep 3
install:
	@go install
all: intro install

# Deploy the app
deploy:
	$(info Deploy the app)
	$(warning WIP)

# Undeploy the app
undeploy:
	$(info Undeploy the app)
	$(warning WIP)

# Build the golang dependancies
build:
	$(info Building golang dependancies)
	cd view && go build
	cd controller && go build
	cd model && go build

# Tests
test:
	$(info Tests)
	$(warning WIP)

# Run the cli (for test purpose)
run:
	$(info Run the cli (for test purpose))
	cd server && go run cli.go

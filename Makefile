APP_NAME := tttm-go
MAIN_FILE := cmd/server/main.go

dev:
	@echo "Running $(APP_NAME) in development mode..."
	air
run:
	@echo "Running $(APP_NAME)..."
	go run $(MAIN_FILE)
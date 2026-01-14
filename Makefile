APP_NAME := latlongapi
MAIN_PKG := .
PORT ?= 8080

.PHONY: run dev build clean stop

build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(APP_NAME) $(MAIN_PKG)

run: build
	@echo "Starting $(APP_NAME) on port $(PORT)..."
	@PORT=$(PORT) ./$(APP_NAME)

stop:
	@echo "Attempting to stop existing server on port $(PORT)..."
	@if command -v lsof >/dev/null 2>&1; then \
		pids=$$(lsof -ti tcp:$(PORT) 2>/dev/null || true); \
		if [ -n "$$pids" ]; then \
			echo "Killing process IDs: $$pids"; \
			kill $$pids || true; \
		else \
			echo "No process found listening on port $(PORT)."; \
		fi; \
	else \
		echo "lsof not found; skipping port-based shutdown. Consider installing lsof for full dev workflow support."; \
	fi

dev:
	@$(MAKE) stop PORT=$(PORT)
	@echo "Starting $(APP_NAME) in dev mode..."
	@set -a; \
	if [ -f .env ]; then \
		echo "Loading environment from .env"; \
		. ./.env; \
	else \
		echo "No .env file found, starting with current shell environment."; \
	fi; \
	set +a; \
	PORT=$(PORT) go run $(MAIN_PKG)

clean:
	@echo "Cleaning build artifacts..."
	@rm -f $(APP_NAME)



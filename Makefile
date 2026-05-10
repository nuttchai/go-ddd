.PHONY: generate dev dev-local

# Regenerate Google Wire injector files.
generate:
	go run github.com/google/wire/cmd/wire@latest ./internal/http/client/config

# Run HTTP service locally from source.
dev:
	APP_ENV=local ENV_PATH=.env.local go run ./cmd/http

# Start local Docker DB and run HTTP service with .env.local.
dev-local:
	docker compose up -d db
	APP_ENV=local ENV_PATH=.env.local go run ./cmd/http

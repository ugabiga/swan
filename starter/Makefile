dev-b:
	air

dev-w:
	cd web && yarn dev & cd web && yarn openapi

dev:
	@make -j2 dev-b dev-w

ent-new:
	@go run -mod=mod entgo.io/ent/cmd/ent new $(name) --target ./internal/ent/schema

ent-gen:
	@go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert ./internal/ent/schema
	@go generate ./internal/ent/


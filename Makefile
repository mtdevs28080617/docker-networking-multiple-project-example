.PHONY: boot
boot:
	@cd app1 && docker compose up -d
	@cd app2 && docker compose up -d
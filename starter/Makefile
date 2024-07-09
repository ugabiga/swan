dev:
	@tmux new-session \; \
	split-window -h \; \
	send-keys 'make dev-w' C-m \; \
	select-pane -t 0 \; \
	send-keys 'air' C-m

dev-w:
	cd web && yarn dev & cd web && yarn openapi

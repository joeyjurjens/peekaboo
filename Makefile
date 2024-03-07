.PHONY: watch install install_templ_cli install_goose_cli

watch:
	@wgo -file=.go -file=.templ -xfile=_templ.go templ generate -p pkg/templates/ :: go run cmd/web/main.go

install: install_templ_cli install_goose_cli

install_templ_cli:
	@go install github.com/a-h/templ/cmd/templ@latest

install_goose_cli:
	@go install github.com/pressly/goose/v3/cmd/goose@latest

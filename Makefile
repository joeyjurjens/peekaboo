.PHONY: watch install install_templ_cli

watch:
	@wgo -file=.go -file=.templ -xfile=_templ.go templ generate :: go run .

install:
	@$(MAKE) install_templ_cli

install_templ_cli:
	go install github.com/a-h/templ/cmd/templ@latest

.PHONY: start
start:
	go run ./cmd/down-to-meet-server --port=9000

define FORMAT
	go list -f ' \
		{{- range .GoFiles}}{{printf "%s/%s\n" $$.Dir .}}{{end -}} \
		{{- range .TestGoFiles}}{{printf "%s/%s\n" $$.Dir .}}{{end -}} \
		{{- range .XTestGoFiles}}{{printf "%s/%s\n" $$.Dir .}}{{end -}} \
	' ./... | xargs go run golang.org/x/tools/cmd/goimports -format-only -l -w -local go.timothygu.me/downtomeet
endef

.PHONY: format
format:
	$(FORMAT)

.PHONY: generate
generate:
	go generate -x ./...
	$(FORMAT)

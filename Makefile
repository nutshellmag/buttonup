GOCMD=/opt/homebrew/bin/go
CADDYCMD= /opt/homebrew/bin/caddy# Optional

build:
	$(GOCMD) build -o buttonup

test:
	$(CADDYCMD) run --config devel.Caddyfile --adapter caddyfile

clean:
	$(GOCMD) clean
	rm buttonup
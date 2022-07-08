GOCMD=/opt/homebrew/bin/go
CADDYCMD= /opt/homebrew/bin/caddy# Optional

build:
	$(GOCMD) build -o buttonup

test:
	./buttonup
	$(CADDYCMD) run --config=devel.Caddyfile

clean:
	$(GOCMD) clean
	rm buttonup
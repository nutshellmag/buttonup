GOCMD=/opt/homebrew/bin/go

build:
	$(GOCMD) build -o buttonup

clean:
	$(GOCMD) clean
	rm buttonup
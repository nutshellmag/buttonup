GOCMD: $(which go)

build:
	$(GOCMD) build -o buttonup

clean:
	$(GOCMD) clean
	rm buttonup
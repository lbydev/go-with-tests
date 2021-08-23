.PHONY: mock
mock:
	rm -rf test/mocks
	mkdir -p test/mocks/blog
	mockgen -package=mblog -source=pkg/blog/blog.go > test/mocks/blog/blog.go

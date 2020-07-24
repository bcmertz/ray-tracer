run: clean build
	./ray-tracer && feh test.ppm
build:
	go build -o ray-tracer
clean:
	rm -f test.ppm ray-tracer

run: clean build
	./ray-tracer && sxiv test.ppm
build:
	go build -o ray-tracer
clean:
	rm -f test.ppm ray-tracer
convert:
	convert test.ppm test.png

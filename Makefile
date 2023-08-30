.PHONY = all

all: clean
	@go build ./cmd/abigen

run-example: all
	./abigen abc ./utils/testdata/0x03db66c8da0e47340f65030aa4c076c5c4864923b54dcc9ee13d87559c1c0b96.json

clean:
	-@rm abc.go abigen

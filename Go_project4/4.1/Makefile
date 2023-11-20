TARGET=myapp

SOURCES=$(wildcard *.go)

GO=go

GOFLAGS=-v

all: $(TARGET)

$(TARGET): $(SOURCES)
	$(GO) build $(GOFLAGS) -o $(TARGET)

ast:
	$(GO) tool compile -W $(SOURCES) > ast.txt

obj:
	$(GO) tool compile -o $(TARGET).o $(SOURCES)

clean:
	rm -f $(TARGET) $(TARGET).o ast.txt

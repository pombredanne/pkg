build:
	go build github.com/ruda/pkg

install:
	go install github.com/ruda/pkg

fmt:
	go fmt *.go

clean:
	rm -rf *~ pkg

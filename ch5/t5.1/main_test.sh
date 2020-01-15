wget google.com \
 && go run main.go < index.html \
 && go test -v *.go \
 && go test -v -bench ./ \
 && rm -f index.html.*

rm -f test.xml \
 && wget -O test.xml www.w3.org/TR/2006/REC-xml11-20060816 \
 && go run main.go < ./test.xml \
 && go test -v *.go \
 && go test -v -bench ./

go run main.go https://www.google.com/ "mngb" "gbar" \
 && go test -v *.go \
 && go test -v -bench ./

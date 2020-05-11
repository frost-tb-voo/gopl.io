# go tool cover
# go help testflag
go get -u gopl.io/ch7/eval

echo
go test -v -run=Coverage -coverprofile=eval.out \
 gopl.io/ch7/eval
echo
go tool cover -html=eval.out -o eval.html
echo
go tool cover -func=eval.out

echo
go test -v -run=Coverage -covermode=count -coverprofile=eval.out \
 gopl.io/ch7/eval
echo
go tool cover -html=eval.out -o eval.html
echo
go tool cover -func=eval.out


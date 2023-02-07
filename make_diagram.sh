rm -r go-diagrams
go run .
cd go-diagrams
dot -Tpng app.dot > D.png

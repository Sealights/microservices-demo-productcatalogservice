go mod download
go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o productcatalogservice .

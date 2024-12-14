package example

//go:generate wrap-mockgen.sh -line=$GOLINE -source=$GOFILE -package=$GOPACKAGE
type uniqueIdentificatorGenerator interface {
    GenerateID() (string, error)
}

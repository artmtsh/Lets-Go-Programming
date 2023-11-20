package generate

// export PATH=$PATH:$(go env GOPATH)/bin
// source ~/.zshrc
//
//go:generate mockgen -destination=mock_foo.go -package=generate . Doer
type Doer interface {
	DoSomething(int, string) error
}

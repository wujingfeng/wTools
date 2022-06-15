package core

type server interface {
	ListenAndServe() error
}

package transport

type ListeningPoint interface {
    Read()
    Build()
}
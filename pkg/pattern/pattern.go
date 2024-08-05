package pattern

type Pattern interface {
    Matches(target string) bool
}

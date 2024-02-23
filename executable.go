package parallelize

type Executable interface {
	Execute() error
}

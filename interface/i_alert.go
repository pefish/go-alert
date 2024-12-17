package interface_

type IAlert interface {
	Infof(format string, a ...any) error
	Warnf(format string, a ...any) error
	Errorf(format string, a ...any) error
}

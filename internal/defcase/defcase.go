package defcase

type DefCase interface {
	Convert(tag, pkgPath, name string) string
	Converter(tag, pkgPath string) func(name string) string
}

var defCase DefCase = noCase{}

func SetDefCase(dcase DefCase) {
	defCase = dcase
}

type noCase struct{}

func (c noCase) Convert(tag, pkgPath, name string) string {
	return name
}

func (c noCase) Converter(tag, pkgPath string) func(name string) string {
	return func(name string) string { return name }
}

func Convert(tag, pkgPath, name string) string {
	return defCase.Convert(tag, pkgPath, name)
}

func Converter(tag, pkgPath string) func(name string) string {
	return defCase.Converter(tag, pkgPath)
}

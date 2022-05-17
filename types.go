package substrate

import (
	"github.com/yspk/scale.go/source"
	"github.com/yspk/scale.go/types"
)

func RegCustomTypes(sourceCode []byte) {
	types.RegCustomTypes(source.LoadTypeRegistry(sourceCode))
}

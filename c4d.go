package c4d

import (
	"github.com/rzrbld/puml-c4-to-object-go/encode"
	"github.com/rzrbld/puml-c4-to-object-go/types"
)

func Parse(pumlc4String string) *types.EncodedObj {
	var encObj = &types.EncodedObj{}
	readedNodes, readedRels := encode.ReadStrings(pumlc4String)
	encObj.Nodes = readedNodes
	encObj.Rels = readedRels
	return encObj
}

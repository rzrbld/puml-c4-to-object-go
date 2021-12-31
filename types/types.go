package types

type GenericC4Type struct {
	Alias string `json:"alias"`
	Label string `json:"label"`
	Techn string `json:"techn"`
	Descr string `json:"descr"`
	Type  string `json:"type"`
	GType string `json:"gtype"`
	Index string `json:"index"`
	From  string `json:"from"`
	To    string `json:"to"`
}

type ParserGenericType struct {
	Object        map[string]interface{} `json:"Object"`
	BoundaryAlias string                 `json:"BoundaryAlias"`
	IsRelation    bool                   `json:"IsRelation"`
}

type EncodedObj struct {
	Nodes []*ParserGenericType `json:"Nodes"`
	Rels  []*ParserGenericType `json:"Rels"`
}

type CompCont struct {
	Alias string `json:"alias"`
	Label string `json:"label"`
	Techn string `json:"techn"`
	Descr string `json:"descr"`
	GType string `json:"gtype"`
}

type PersSystem struct {
	Alias string `json:"alias"`
	Label string `json:"label"`
	Descr string `json:"descr"`
	GType string `json:"gtype"`
}

type Boundary struct {
	Alias string `json:"alias"`
	Label string `json:"label"`
	Type  string `json:"type"`
	GType string `json:"gtype"`
	Descr string `json:"descr"` //fake it
	Techn string `json:"techn"` //fake it
}

type Node struct {
	Alias string `json:"alias"`
	Label string `json:"label"`
	Type  string `json:"type"`
	Descr string `json:"descr"`
	GType string `json:"gtype"`
}

type Rel struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Label string `json:"label"`
	Techn string `json:"techn"`
	Descr string `json:"descr"`
	GType string `json:"gtype"`
}

type RelIndex struct {
	Index string `json:"index"`
	From  string `json:"from"`
	To    string `json:"to"`
	Label string `json:"label"`
	Techn string `json:"techn"`
	Descr string `json:"descr"`
	GType string `json:"gtype"`
}

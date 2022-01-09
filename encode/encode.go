package encode

import (
	"regexp"
	"strings"

	"github.com/fatih/structs"
	"github.com/rzrbld/puml-c4-to-object-go/types"
	log "github.com/sirupsen/logrus"
)

func ReadStrings(pumlc4String string) ([]*types.ParserGenericType, []*types.ParserGenericType) {

	frNodes := []*types.ParserGenericType{}
	frRels := []*types.ParserGenericType{}
	var obj = &types.ParserGenericType{}
	str := string(pumlc4String)

	reBoundary := regexp.MustCompile(`(?m)(.*)(\(.*\)).*\{((.|\n)*)}`)
	strBoundary := reBoundary.FindAllString(str, -1)

	log.Traceln("-------------- BOUNDARIES ---------------------")

	if len(strBoundary) > 0 {
		bName := "1"
		for i, match := range reBoundary.FindAllString(str, -1) {
			log.Traceln("Boundary:", match, " found at index:", i)
			str = strings.ReplaceAll(str, string(match), "")
			log.Traceln("-------------- NODES IN BOUNDARIES ---------------------")
			re := regexp.MustCompile(`(?m)(.*)\((.*),(.*)\)`)
			for g, match2 := range re.FindAllString(match, -1) {
				if g == 0 {
					bName = GetAliasName(match2)
					obj = ParseMatch(match2, true, "")

					if len(obj.Object) != 0 {
						if obj.IsRelation {
							frRels = append(frRels, obj)
						} else {
							frNodes = append(frNodes, obj)
						}
					}
				} else {
					obj = ParseMatch(match2, true, bName)
					if len(obj.Object) != 0 {
						if obj.IsRelation {
							frRels = append(frRels, obj)
						} else {
							frNodes = append(frNodes, obj)
						}
					}
				}
			}
		}
	} else {
		log.Warnln("boundary not found")
	}

	log.Traceln("-------------- NODES and RELS ---------------------")

	re := regexp.MustCompile(`(?m)(.*)\((.*),(.*)\)`)
	for _, match := range re.FindAllString(str, -1) {
		obj = ParseMatch(match, false, "")
		if len(obj.Object) != 0 {
			if obj.IsRelation {
				frRels = append(frRels, obj)
			} else {
				frNodes = append(frNodes, obj)
			}
		}
	}

	return frNodes, frRels
}

func ParseMatch(match string, isBoundary bool, bAlias string) *types.ParserGenericType {

	var obj = make(map[string]interface{})
	var boundaryAlias string
	var isRel bool
	result := &types.ParserGenericType{}
	// 0 is always boundary
	reGetAttrib := regexp.MustCompile(`\((.*)\)`)
	reAttr := reGetAttrib.FindAllString(match, -1)
	strAttr := ""
	trimmedAttrString := ""

	if len(reAttr) > 0 {
		trimmedAttrString = strings.TrimSpace(reAttr[0])
		strAttr = trimBrackets(trimmedAttrString)
	} else {
		log.Errorln("Error empty attributes")
	}

	log.Traceln("newAttrString", trimmedAttrString, "Len:", len(reAttr))

	getType := strings.Split(match, "(")
	log.Traceln("getType", getType)

	nodeType := strings.TrimSpace(getType[0])
	log.Traceln("nodeType", nodeType)

	attrSlice := SplitAtCommas(strAttr)
	log.Traceln("attrSlice", attrSlice)

	// attrSlice[0] is always alias name - trim it
	attrSlice[0] = strings.Trim(attrSlice[0], " ")

	if isBoundary {
		log.Traceln("NODE TYPE b >>", nodeType, "NODE ATTR b str >>", strAttr, "NODE ATTR b str >>", attrSlice)
		if bAlias != "" {
			log.Traceln("NODE Bound relation>>", bAlias)
		}
		obj, boundaryAlias, isRel = matchToTypes(nodeType, attrSlice[0], attrSlice, bAlias)

	} else {
		log.Traceln("NODE TYPE >>", nodeType, "NODE ATTR str >>", strAttr, "NODE ATTR str >>", attrSlice)
		obj, boundaryAlias, isRel = matchToTypes(nodeType, attrSlice[0], attrSlice, bAlias)
	}

	result.BoundaryAlias = boundaryAlias
	result.Object = obj
	result.IsRelation = isRel

	return result
}

func trimBrackets(s string) string {
	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	s = s[:0]

	return s
}

func matchToTypes(nodeType string, nodeAlias string, nodeAttr []string, boundaryAlias string) (map[string]interface{}, string, bool) {
	var obj = make(map[string]interface{})
	relFlag := false
	switch ntype := nodeType; ntype {

	case "Component", "ComponentDb", "ComponentQueue", "Component_Ext", "ComponentDb_Ext", "ComponentQueue_Ext", "Container", "ContainerDb", "ContainerQueue", "Container_Ext", "ContainerDb_Ext", "ContainerQueue_Ext":
		// ($alias, $label, $techn, $descr="", $sprite="", $tags="", $link="")
		nodeAttr = NormalizeArr(nodeAttr, 4)
		node := &types.CompCont{}
		node.GType = ntype
		for i, val := range nodeAttr {
			switch i {
			case 0:
				node.Alias = strings.TrimSpace(val)
			case 1:
				node.Label = strings.TrimSpace(val)
			case 2:
				node.Techn = strings.TrimSpace(val)
			case 3:
				node.Descr = strings.TrimSpace(val)
			}
		}
		obj = structs.Map(node)

	case "Person", "Person_Ext", "System", "System_Ext", "SystemDb", "SystemQueue", "SystemDb_Ext", "SystemQueue_Ext", "Enterprise": //Enterprise is fake
		// $alias, $label, $descr="", $sprite="", $tags="", $link=""
		nodeAttr = NormalizeArr(nodeAttr, 3)
		node := &types.PersSystem{}
		node.GType = ntype
		for i, val := range nodeAttr {
			switch i {
			case 0:
				node.Alias = strings.TrimSpace(val)
			case 1:
				node.Label = strings.TrimSpace(val)
			case 2:
				node.Descr = strings.TrimSpace(val)
			}
		}
		obj = structs.Map(node)

	case "Enterprise_Boundary", "System_Boundary", "Container_Boundary":
		// Boundary($alias, $label, "Enterprise", $tags, $link)
		nodeAttr = NormalizeArr(nodeAttr, 2)
		node := &types.Boundary{}

		node.Type = strings.ReplaceAll(ntype, "_Boundary", "")
		node.GType = node.Type
		node.Descr = ""
		node.Techn = ""
		for i, val := range nodeAttr {
			switch i {
			case 0:
				node.Alias = strings.TrimSpace(val)
			case 1:
				node.Label = strings.TrimSpace(val)
			}
		}
		obj = structs.Map(node)

	case "Deployment_Node", "Deployment_Node_L", "Deployment_Node_R", "Node", "Node_L", "Node_R":
		// $alias, $label, $type="", $descr="", $sprite="", $tags="", $link=""
		nodeAttr = NormalizeArr(nodeAttr, 4)
		node := &types.Node{}
		node.GType = ntype
		for i, val := range nodeAttr {
			switch i {
			case 0:
				node.Alias = strings.TrimSpace(val)
			case 1:
				node.Label = strings.TrimSpace(val)
			case 2:
				node.Type = strings.TrimSpace(val)
			case 3:
				node.Descr = strings.TrimSpace(val)
			}
		}

		obj = structs.Map(node)

	case "Rel", "Rel_Back", "Rel_Neighbor", "Rel_Back_Neighbor", "Rel_D", "Rel_Down", "Rel_U", "Rel_Up", "Rel_L", "Rel_Left", "Rel_R", "Rel_Right":
		// $from, $to, $label, $techn="", $descr="", $sprite="", $tags="", $link=""
		nodeAttr = NormalizeArr(nodeAttr, 5)
		rel := &types.Rel{}
		rel.GType = ntype
		for i, val := range nodeAttr {
			switch i {
			case 0:
				rel.From = strings.TrimSpace(val)
			case 1:
				rel.To = strings.TrimSpace(val)
			case 2:
				rel.Label = strings.TrimSpace(val)
			case 3:
				rel.Techn = strings.TrimSpace(val)
			case 4:
				rel.Descr = strings.TrimSpace(val)
			}
		}

		obj = structs.Map(rel)
		relFlag = true

	case "RelIndex", "RelIndex_Back", "RelIndex_Neighbor", "RelIndex_Back_Neighbor", "RelIndex_D", "RelIndex_Down", "RelIndex_U", "RelIndex_Up", "RelIndex_L", "RelIndex_Left", "RelIndex_R", "RelIndex_Right":
		// $e_index, $from, $to, $label, $techn="", $descr="", $sprite="", $tags="", $link=""
		nodeAttr = NormalizeArr(nodeAttr, 6)
		rel := &types.RelIndex{}
		rel.GType = ntype
		for i, val := range nodeAttr {
			switch i {
			case 0:
				rel.Index = strings.TrimSpace(val)
			case 1:
				rel.From = strings.TrimSpace(val)
			case 2:
				rel.To = strings.TrimSpace(val)
			case 3:
				rel.Label = strings.TrimSpace(val)
			case 4:
				rel.Techn = strings.TrimSpace(val)
			case 5:
				rel.Descr = strings.TrimSpace(val)
			}
		}

		obj = structs.Map(rel)
		relFlag = true

	default:
		log.Warnln("Unknown Node Type:", nodeType)
	}

	return obj, boundaryAlias, relFlag
}

func GetAliasName(match string) string {
	strParts := strings.Split(match, "(")
	strAttr := strings.Split(strParts[1], ")")[0]
	attrSlice := strings.Split(strAttr, ",")
	// attrSlice[0] is always alias name - trim it
	attrSlice[0] = strings.Trim(attrSlice[0], " ")
	return attrSlice[0]
}

func SplitAtCommas(s string) []string {
	res := []string{}
	var beg int
	var inString bool

	for i := 0; i < len(s); i++ {
		if s[i] == ',' && !inString {
			res = append(res, s[beg:i])
			beg = i + 1
		} else if s[i] == '"' {
			if !inString {
				inString = true
			} else if i > 0 && s[i-1] != '\\' {
				inString = false
			}
		}
	}
	return append(res, s[beg:])
}

func NormalizeArr(slice []string, targetSize int) []string {
	if len(slice) < targetSize {
		for i := 0; i < targetSize-len(slice); i++ {
			slice = append(slice, "Undefined")
		}
	}
	return slice
}

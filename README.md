# PlantUML C4 to object parser

Turns puml C4 notation in to structurized object

## QuckStart guide:

Complete example located at `example/main.go`

### 1. Prepare puml c4 content
For example one string:
``` 
System(taxamo, "Taxamo", "Calculates local tax (for EU B2B customers) and acts as a front-end for Braintree Payments.")

```

### 2. Add parser
```go
    import (
		...
		pc42obj "github.com/rzrbld/puml-c4-to-object-go"
		"github.com/rzrbld/puml-c4-to-object-go/types"
		...
	)
    ...
    pumlC4Str := `System(taxamo, "Taxamo", "Calculates local tax (for EU B2B customers) and acts as a front-end for Braintree Payments.")`

    // do Parse stuff
	var testObj = &types.EncodedObj{}
	testObj = pc42obj.Parse(pumlC4Str)

	// marshal to json
	jsonMarshaled, _ := json.MarshalIndent(testObj, "", "\t")

	// print json output
	fmt.Println(string(jsonMarshaled))

    ...

```

### 3. Compile and run

```json

{
	"Nodes": [
		{
			"Object": {
				"Alias": "taxamo",
				"Descr": "\"Calculates local tax (for EU B2B customers) and acts as a front-end for Braintree Payments.\"",
				"GType": "System",
				"Label": "\"Taxamo\""
			},
			"BoundaryAlias": "",
			"IsRelation": false
		}
	],
	"Rels": []
}

```

## Little more complex example

```go

import (
		...
		pc42obj "github.com/rzrbld/puml-c4-to-object-go"
		"github.com/rzrbld/puml-c4-to-object-go/types"
		...
	)

func main() {
    //...
    pumlC4Str := `System(taxamo, "Taxamo", "Calculates local tax (for EU B2B customers) and acts as a front-end for Braintree Payments.")`

    // do Parse stuff
	var testObj = &types.EncodedObj{}
	testObj = pc42obj.Parse(pumlC4Str)

    foreachObjects(testObj.Nodes)
    //...
}

func foreachObjects(objMap []*types.ParserGenericType) {
	for index, elem := range objMap {

		fmt.Println("---", index, "---")
		fmt.Println(" Object: ", elem.Object, " BoundaryAlias:", elem.BoundaryAlias, " IsRelation:", elem.IsRelation)

		var node types.GenericC4Type
		err := mapstructure.Decode(elem.Object, &node)
		if err != nil {
			log.Errorln("Kind of error. ", err)
		}

		fmt.Println("Alias", node.Alias,
			"GType", node.GType,
			"Label", node.Label,
			"Techn", node.Techn,
			"Descr", node.Descr,
			"Type", node.Type,
			"Index", node.Index,
			"From", node.From,
			"To", node.To)
	}
}

```

### Compile and run

```
--- 0 ---
 Object:  map[Alias:taxamo Descr:"Calculates local tax (for EU B2B customers) and acts as a front-end for Braintree Payments." GType:System Label:"Taxamo"]  BoundaryAlias:   IsRelation: false

Alias taxamo
GType System
Label "Taxamo"
Techn
Descr "Calculates local tax (for EU B2B customers) and acts as a front-end for Braintree Payments."
Type
Index
From
To

```

## Supported elements and attributes

| Element | Supported attributes | Unsupported attrinbutes |
| ------- | -------------------- | -------------------- | 
| "Component", "ComponentDb", "ComponentQueue", "Component_Ext", "ComponentDb_Ext", "ComponentQueue_Ext", "Container", "ContainerDb", "ContainerQueue", "Container_Ext", "ContainerDb_Ext", "ContainerQueue_Ext"   | alias, label, techn, descr  |  sprite, tags, link |
| "Person", "Person_Ext", "System", "System_Ext", "SystemDb", "SystemQueue", "SystemDb_Ext", "SystemQueue_Ext"  | alias, label, descr | sprite, tags, link |
| "Enterprise_Boundary", "System_Boundary", "Container_Boundary" | alias, label | sprite, tags, link |
| "Deployment_Node", "Deployment_Node_L", "Deployment_Node_R", "Node", "Node_L", "Node_R" | alias, label, type, descr | sprite, tags, link |
| "Rel", "Rel_Back", "Rel_Neighbor", "Rel_Back_Neighbor", "Rel_D", "Rel_Down", "Rel_U", "Rel_Up", "Rel_L", "Rel_Left", "Rel_R", "Rel_Right" | from, to, label, techn, descr, | sprite, tags, link |
| "RelIndex", "RelIndex_Back", "RelIndex_Neighbor", "RelIndex_Back_Neighbor", "RelIndex_D", "RelIndex_Down", "RelIndex_U", "RelIndex_Up", "RelIndex_L", "RelIndex_Left", "RelIndex_R", "RelIndex_Right" | e_index, from, to, label, techn, descr | sprite, tags, link |



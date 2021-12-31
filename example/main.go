package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
	pc42obj "github.com/rzrbld/puml-c4-to-object-go"
	"github.com/rzrbld/puml-c4-to-object-go/types"
	log "github.com/sirupsen/logrus"
)

func main() {
	// set loglevel to Error
	log.SetLevel(log.ErrorLevel)

	// put plantUML C4 file content in to this variable
	pumlC4Str := `@startuml "enterprise"
	!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Context.puml
	' uncomment the following line and comment the first to use locally
	' !include C4_Context.puml
	
	LAYOUT_TOP_DOWN()
	'LAYOUT_AS_SKETCH()
	LAYOUT_WITH_LEGEND()
	
	Person(customer, "Customer", "A customer of Widgets Limited.")
	
	Enterprise_Boundary(c0, "Widgets Limited") {
		Person(csa, "Customer Service Agent", "Deals with customer enquiries.")
	
		System(ecommerce, "E-commerce System", "Allows customers to buy widgts online via the widgets.com website.")
	
		System(fulfilment, "Fulfilment System", "Responsible for processing and shipping of customer orders.")
	}
	
	System(taxamo, "Taxamo", "Calculates local tax (for EU B2B customers) and acts as a front-end for Braintree Payments.")
	
	System(braintree, "Braintree Payments", "Processes credit card payments on behalf of Widgets Limited.")
	
	System(post, "Jersey Post", "Calculates worldwide shipping costs for packages.")
	
	Rel_R(customer, csa, "Asks questions to", "Telephone")
	
	Rel_R(customer, ecommerce, "Places orders for widgets using")
	
	Rel(csa, ecommerce, "Looks up order information using")
	
	Rel_R(ecommerce, fulfilment, "Sends order information to")
	
	Rel_D(fulfilment, post, "Gets shipping charges from")
	
	Rel_D(ecommerce, taxamo, "Delegates credit card processing to")
	
	Rel_L(taxamo, braintree, "Uses for credit card processing")
	
	Lay_D(customer, braintree)
	
	@enduml
	`
	// do Parse stuff
	var testObj = &types.EncodedObj{}
	testObj = pc42obj.Parse(pumlC4Str)

	// marshal to json
	jsonMarshaled, _ := json.MarshalIndent(testObj, "", "\t")

	// print json output
	fmt.Println(string(jsonMarshaled))

	// foreach Nodes as generic type
	fmt.Println("------------------------------------------------------- Nodes ------------------------------------------------------")
	foreachObjects(testObj.Nodes)

	// foreach Relations as generic type
	fmt.Println("----------------------------------------------------- Relations -----------------------------------------------------")
	foreachObjects(testObj.Rels)

}

func foreachObjects(objMap []*types.ParserGenericType) {
	for index, elem := range objMap {
		fmt.Println("-----------------------------------------------------", index, "-----------------------------------------------------")
		fmt.Println(" Object: ", elem.Object, " BoundaryAlias:", elem.BoundaryAlias, " IsRelation:", elem.IsRelation)
		var node types.GenericC4Type
		err := mapstructure.Decode(elem.Object, &node)
		if err != nil {
			log.Errorln("Some went wrong on map structure. ", err)
		}

		fmt.Println("\nAlias", node.Alias,
			"\nGType", node.GType,
			"\nLabel", node.Label,
			"\nTechn", node.Techn,
			"\nDescr", node.Descr,
			"\nType", node.Type,
			"\nIndex", node.Index,
			"\nFrom", node.From,
			"\nTo", node.To)

		fmt.Println("")

	}
}

# PlantUML C4 to object parser

Turns puml C4 notation in to structurized object

## QuckStart guide:
look at `example/main.go`

```golang

package main

import (
	"encoding/json"
	"fmt"

	pc42obj "github.com/rzrbld/puml-c4-to-object-go"
	"github.com/rzrbld/puml-c4-to-object-go/types"
)

func main() {
	pumlC4Str := `
	@startuml
	!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml
	
	SHOW_PERSON_OUTLINE()
	AddElementTag("backend container", $fontColor=$ELEMENT_FONT_COLOR, $bgColor="#335DA5", $shape=EightSidedShape())
	AddRelTag("async", $textColor=$ARROW_COLOR, $lineColor=$ARROW_COLOR, $lineStyle=DashedLine())
	AddRelTag("sync/async", $textColor=$ARROW_COLOR, $lineColor=$ARROW_COLOR, $lineStyle=DottedLine())
	
	title Container diagram for Internet Banking System
	
	Person(customer, Customer, "A customer of the bank, with personal bank accounts")
	
	System_Boundary(c1, "Internet Banking") {
		Container(web_app, "Web Application", "Java, Spring MVC", "Delivers the static content and the Internet banking SPA")
		Container(spa, "Single-Page App", "JavaScript, Angular", "Provides all the Internet banking functionality to cutomers via their web browser")
		Container(mobile_app, "Mobile App", "C#, Xamarin", "Provides a limited subset of the Internet banking functionality to customers via their mobile device")
		ContainerDb(database, "Database", "SQL Database", "Stores user registration information, hashed auth credentials, access logs, etc.")
		Container(backend_api, "API Application", "Java, Docker Container", "Provides Internet banking functionality via API", $tags="backend container")
	}
	
	System_Ext(email_system, "E-Mail System", "The internal Microsoft Exchange system")
	System_Ext(banking_system, "Mainframe Banking System", "Stores all of the core banking information about customers, accounts, transactions, etc.")
	
	Rel(customer, web_app, "Uses", "HTTPS")
	Rel(customer, spa, "Uses", "HTTPS")
	Rel(customer, mobile_app, "Uses")
	
	Rel_Neighbor(web_app, spa, "Delivers")
	Rel(spa, backend_api, "Uses", "async, JSON/HTTPS", $tags="async")
	Rel(mobile_app, backend_api, "Uses", "async, JSON/HTTPS", $tags="async")
	Rel_Back_Neighbor(database, backend_api, "Reads from and writes to", "sync, JDBC")
	
	Rel_Back(customer, email_system, "Sends e-mails to")
	Rel_Back(email_system, backend_api, "Sends e-mails using", "sync, SMTP")
	Rel_Neighbor(backend_api, banking_system, "Uses", "sync/async, XML/HTTPS", $tags="sync/async")
	
	SHOW_LEGEND()
	@enduml
	`
	var test = &types.EncodedObj{}
	test = pc42obj.Parse(pumlC4Str)
	foo_marshalled, _ := json.Marshal(test)
	fmt.Println(string(foo_marshalled)) // write response to ResponseWriter (w)

	// fmt.Println("OUTPUT", b)

}

```

let's suggest you have puml C4 like this:


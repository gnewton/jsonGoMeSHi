package main

import (
        "net/http"
        "log"
	"flag"
//        "os"
        "runtime"
        "fmt"
        "github.com/ant0ine/go-json-rest/rest"
)

var URL_HOST *string
var PORT *string
var PATH *string
var BASE_URL string

const MAX_PAGE_SIZE = 100


// Change these to match the MeSH files you download from http://www.nlm.nih.gov/mesh/ist.html
// The XML files are desc2014, supp2014, qual2014
// You can compress these with either gz or bz2 & this app will transparently uncompress them. Or you can leave them as-is.
//
var DESCRIPTOR_XML_FILE *string
var QUALIFIER_XML_FILE *string
var SUPPLEMENTAL_XML_FILE *string
var PHARMACOLOGICAL_XML_FILE *string

func main() {

	flags()
	
	//URL_HOST,_ := os.Hostname()
	BASE_URL ="http://" + *URL_HOST + ":" + *PORT + *PATH
	log.Println("Base URL:", BASE_URL)
	runtime.GOMAXPROCS(runtime.NumCPU())
	err := loadData()
	if err != nil{
		fmt.Println(err)
		return 
	}
        handler := rest.ResourceHandler{ 
		EnableGzip: true,
	}
        handler.SetRoutes(
		&rest.Route{"GET", *PATH, GetAll},

		&rest.Route{"GET", *PATH + "/" + DESCRIPTOR, GetAllDescriptors},
		&rest.Route{"GET", *PATH + "/" + DESCRIPTOR + "/", GetAllDescriptors},
                &rest.Route{"GET", *PATH + "/" + DESCRIPTOR + "/:id", GetDescriptor},

		&rest.Route{"GET", *PATH + "/" + SUPPLEMENTAL + "/:id", GetSupplemental},
		&rest.Route{"GET", *PATH + "/" + SUPPLEMENTAL, GetAllSupplementals},
		&rest.Route{"GET", *PATH + "/" + SUPPLEMENTAL + "/", GetAllSupplementals},

		&rest.Route{"GET", *PATH + "/" + QUALIFIER + "/:id", GetQualifier},
		&rest.Route{"GET", *PATH + "/" + QUALIFIER + "/", GetAllQualifiers},
		&rest.Route{"GET", *PATH + "/" + QUALIFIER, GetAllQualifiers},

		&rest.Route{"GET", *PATH + "/" + PHARMACOLOGICAL + "/:id", GetPharmacological},
		&rest.Route{"GET", *PATH + "/" + PHARMACOLOGICAL + "/", GetAllPharmacologicals},
		&rest.Route{"GET", *PATH + "/" + PHARMACOLOGICAL, GetAllPharmacologicals},


		&rest.Route{"GET", *PATH + "/" + TREE, GetTrees},
		&rest.Route{"GET", *PATH + "/" + TREE + "/", GetTrees},

		&rest.Route{"GET", *PATH + "/" + TREE+ "/:a", GetTree},
		&rest.Route{"GET", *PATH + "/" + TREE+ "/:a/", GetTree},
		&rest.Route{"GET", *PATH + "/" + TREE+ "/:a.:b", GetTree},
		&rest.Route{"GET", *PATH + "/" + TREE+ "/:a.:b.:c", GetTree},
		&rest.Route{"GET", *PATH + "/" + TREE+ "/:a.:b.:c.:d", GetTree},
		&rest.Route{"GET", *PATH + "/" + TREE+ "/:a.:b.:c.:d.:e", GetTree},
		&rest.Route{"GET", *PATH + "/" + TREE+ "/:a.:b.:c.:d.:e.:f", GetTree},
		&rest.Route{"GET", *PATH + "/" + TREE+ "/:a.:b.:c.:d.:e.:f.:g", GetTree},
		&rest.Route{"GET", *PATH + "/" + TREE+ "/:a.:b.:c.:d.:e.:f.:g.:h", GetTree},
        )
        http.ListenAndServe(":" + *PORT, &handler)
}


// flags
// -h host
// -p port
// -t path
// -D descriptorXML
// -Q qualifierXml
// -S supplementalXml
// -P PHARMACOLOGICAL_XML_FILE
func flags(){
	URL_HOST = flag.String("h", "localhost", "Host name for web service")
	PORT = flag.String("p", "8080", "Port name for web service")
	PATH = flag.String("b", "/mesh", "Base path for web service URL")

	DESCRIPTOR_XML_FILE = flag.String("D", "testData/desc2014_29records.xml.bz2", "Full path to descriptor XML file")
	QUALIFIER_XML_FILE = flag.String("Q", "testData/qual2014_8records.xml.bz2", "Full path to qualifier XML file")
	SUPPLEMENTAL_XML_FILE = flag.String("S", "testData/supp2014_4records.xml", "Full path to supplemental XML file")
	PHARMACOLOGICAL_XML_FILE = flag.String("P", "testData/pa2014_8records.xml", "Full path to pharmacological supplemental XML file")
	flag.Parse()
}
package main

import (
        "github.com/ant0ine/go-json-rest/rest"
        "net/http"
        "log"
        "os"
        "runtime"
        "fmt"
)

var URL_HOST string
const PORT = "8080"
const PATH = "/mesh"

var BASE_URL string

const NLM_TERMS_URL="https://www.nlm.nih.gov/mesh/termscon.html"
const MESH_OWNER="U.S. National Library of Medicine is the creator, maintainer, and provider of this data"
const MESH_VERSION="2014 MeSH"
const MESH_COPYRIGHT="Copyright U.S. National Library of Medicine"

const MAX_PAGE_SIZE = 100

// Change these to match the MeSH files you download from http://www.nlm.nih.gov/mesh/ist.html
// The XML files are desc2014, supp2014, qual2014
// You can compress these with either gz or bz2 & this app will transparently uncompress them. Or you can leave them as-is.
//
const DESCRIPTOR_XML_FILE = "testData/desc2014_29records.xml.bz2"
//const DESCRIPTOR_XML_FILE = "/home/newtong/2014/mesh/desc2014.xml.bz2"

const QUALIFIER_XML_FILE = "testData/qual2014_8records.xml.bz2"
//const QUALIFIER_XML_FILE = "/home/newtong/2014/mesh/qual2014.xml.bz2"

const SUPPLEMENTAL_XML_FILE = "testData/supp2014_4records.xml"
//const SUPPLEMENTAL_XML_FILE = "/home/newtong/2014/mesh/supp2014.xml.bz2"

const PHARMACOLOGICAL_XML_FILE = "testData/pa2014_8records.xml"
//const PHARMACOLOGICAL_XML_FILE = "/home/newtong/2014/mesh/pa2014.xml"

func main() {
	URL_HOST,_ := os.Hostname()
	BASE_URL ="http://" + URL_HOST + ":" + PORT + PATH
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
		&rest.Route{"GET", PATH, GetAll},

		&rest.Route{"GET", PATH + "/" + DESCRIPTOR, GetAllDescriptors},
		&rest.Route{"GET", PATH + "/" + DESCRIPTOR + "/", GetAllDescriptors},
                &rest.Route{"GET", PATH + "/" + DESCRIPTOR + "/:id", GetDescriptor},

		&rest.Route{"GET", PATH + "/" + SUPPLEMENTAL + "/:id", GetSupplemental},
		&rest.Route{"GET", PATH + "/" + SUPPLEMENTAL, GetAllSupplementals},
		&rest.Route{"GET", PATH + "/" + SUPPLEMENTAL + "/", GetAllSupplementals},

		&rest.Route{"GET", PATH + "/" + QUALIFIER + "/:id", GetQualifier},
		&rest.Route{"GET", PATH + "/" + QUALIFIER + "/", GetAllQualifiers},
		&rest.Route{"GET", PATH + "/" + QUALIFIER, GetAllQualifiers},

		&rest.Route{"GET", PATH + "/" + PHARMACOLOGICAL + "/:id", GetPharmacological},
		&rest.Route{"GET", PATH + "/" + PHARMACOLOGICAL + "/", GetAllPharmacologicals},
		&rest.Route{"GET", PATH + "/" + PHARMACOLOGICAL, GetAllPharmacologicals},


		&rest.Route{"GET", PATH + "/" + TREE, GetTrees},
		&rest.Route{"GET", PATH + "/" + TREE + "/", GetTrees},

		&rest.Route{"GET", PATH + "/" + TREE+ "/:a", GetTree},
		&rest.Route{"GET", PATH + "/" + TREE+ "/:a/", GetTree},
		&rest.Route{"GET", PATH + "/" + TREE+ "/:a.:b", GetTree},
		&rest.Route{"GET", PATH + "/" + TREE+ "/:a.:b.:c", GetTree},
		&rest.Route{"GET", PATH + "/" + TREE+ "/:a.:b.:c.:d", GetTree},
		&rest.Route{"GET", PATH + "/" + TREE+ "/:a.:b.:c.:d.:e", GetTree},
		&rest.Route{"GET", PATH + "/" + TREE+ "/:a.:b.:c.:d.:e.:f", GetTree},
		&rest.Route{"GET", PATH + "/" + TREE+ "/:a.:b.:c.:d.:e.:f.:g", GetTree},
		&rest.Route{"GET", PATH + "/" + TREE+ "/:a.:b.:c.:d.:e.:f.:g.:h", GetTree},
        )
        http.ListenAndServe(":" + PORT, &handler)
}

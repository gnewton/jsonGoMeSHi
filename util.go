package main

import (
	"github.com/gnewton/jianGoMeSHi"
        "log"
        "sort"
//        "strings"
)


const DESCRIPTOR = "descriptor"
const QUALIFIER = "qualifier"
const SUPPLEMENTAL = "supplemental"
const TREE = "tree"

var NOUNS = []string{DESCRIPTOR, QUALIFIER, SUPPLEMENTAL, TREE}
var allNouns []jianGoMeSHi.IdEntry

var descMap map[string]*jianGoMeSHi.DescriptorRecord
var descMap2 map[string]*LocalDesc
var descSlice  []*jianGoMeSHi.IdEntry


var suppMap map[string]*jianGoMeSHi.SupplementalRecord
var suppSlice  []*jianGoMeSHi.IdEntry

var qualMap map[string]*jianGoMeSHi.QualifierRecord
var qualSlice  []*jianGoMeSHi.IdEntry

var root *jianGoMeSHi.Node
var treeMap map[string]*jianGoMeSHi.Node

type LocalDesc jianGoMeSHi.DescriptorRecord

type Paging struct{
	Limit int  `json:"limit"`
	Offset int `json:"offset"`
	Count int  `json:"count"`
	NextPageUrl string `json:"nextPageUrl,omitempty"`
	PrevioustPageUrl string `json:"previousPageUrl,omitempty"`
}


func (desc *LocalDesc) setTreeNumberUrls(baseUrl string){
	if desc.TreeNumberList.TreeNumber != nil{
		for i:=0; i<len(desc.TreeNumberList.TreeNumber); i++{
			tn := &(desc.TreeNumberList.TreeNumber[i])
			tn.Url = baseUrl + "/" + TREE + "/" + tn.TreeNumber
		}
	}
}


func (desc *LocalDesc) setDescUrls(baseUrl string){
	if desc.PharmacologicalActionList.PharmacologicalAction != nil{
		for i:=0; i<len(desc.PharmacologicalActionList.PharmacologicalAction); i++{
			ref := &(desc.PharmacologicalActionList.PharmacologicalAction[i])
			ref.DescriptorReferredTo.Url = baseUrl + "/" + DESCRIPTOR + "/" + ref.DescriptorReferredTo.DescriptorUI
		}
	}

	if desc.SeeRelatedList.SeeRelatedDescriptor != nil{
		for i:=0; i<len(desc.SeeRelatedList.SeeRelatedDescriptor); i++{
			ref := &(desc.SeeRelatedList.SeeRelatedDescriptor[i])
			ref.DescriptorReferredTo.Url = baseUrl + "/" + DESCRIPTOR + "/" + ref.DescriptorReferredTo.DescriptorUI
		}
	}
}



func loadData()(error){
	treeMap = make(map[string]*jianGoMeSHi.Node)
	var err error
	log.Println("Start Loading MeSH XML...")

	log.Println("\tLoading Supplemental MeSH XML from file: ", SUPPLEMENTAL_XML_FILE)
	suppMap, err = jianGoMeSHi.SupplementalMapFromFile(SUPPLEMENTAL_XML_FILE)
	if err != nil{
		return err
	}
	index := 0

	suppSlice = make([]*jianGoMeSHi.IdEntry, len(suppMap))

	for supp := range suppMap{
		newEntry := new(jianGoMeSHi.IdEntry)
		newEntry.Id = suppMap[supp].SupplementalRecordUI
		newEntry.Url = BASE_URL + "/" + SUPPLEMENTAL + "/" + newEntry.Id
		suppSlice[index] = newEntry
		index += 1
	}


	log.Println("\tLoading Qualifier MeSH XML from file:", QUALIFIER_XML_FILE)
	qualMap, err = jianGoMeSHi.QualifierMapFromFile(QUALIFIER_XML_FILE)
	if err != nil{
		return err
	}

	qualSlice = make([]*jianGoMeSHi.IdEntry, len(qualMap))
	index = 0
	for qual := range qualMap{
		newEntry := new(jianGoMeSHi.IdEntry)
		newEntry.Id = qualMap[qual].QualifierUI
		newEntry.Url = BASE_URL + "/" + QUALIFIER + "/" + newEntry.Id
		qualSlice[index] = newEntry
		index += 1
	}

	log.Println("\tLoading Descriptor MeSH XML from file: ", DESCRIPTOR_XML_FILE)
	descMap, err = jianGoMeSHi.DescriptorMapFromFile(DESCRIPTOR_XML_FILE)
	if err != nil{
		return err
	}
	log.Println("\tBuilding name map")
	_ = jianGoMeSHi.MeshDescriptorNameMap(descMap)

	descSlice = make([]*jianGoMeSHi.IdEntry, len(descMap))
	index = 0
	descMap2 = make(map[string]*LocalDesc)

	for desc := range descMap{
		newEntry := new(jianGoMeSHi.IdEntry)
		descriptorRecord := descMap[desc]
		var localDesc = (*LocalDesc)(descriptorRecord)
		localDesc.setDescUrls(BASE_URL)
		localDesc.setTreeNumberUrls(BASE_URL)
		
		descMap2[desc] = localDesc
		newEntry.Id = descMap[desc].DescriptorUI
		newEntry.Url = BASE_URL + "/" + DESCRIPTOR + "/" + newEntry.Id
		descSlice[index] = newEntry
		index += 1
	}

	sort.Sort(ById(descSlice))
	sort.Sort(ById(qualSlice))
	sort.Sort(ById(suppSlice))

	root = jianGoMeSHi.MakeTree(descMap)
	root.Traverse(0, AddUrlInfo)
	sort.Sort(ByIdX(root.Children))
	
	log.Println("Done Loading MeSH XML...")

	allNouns = make([]jianGoMeSHi.IdEntry, len(NOUNS))
	for i,noun := range NOUNS{
		allNouns[i].Id = "/" + noun
		allNouns[i].Url = BASE_URL + "/" + noun
	}


	return nil
}

func AddUrlInfo(node *jianGoMeSHi.Node){
	//fmt.Println("AddUrlInfo", node.TreeNumber)
	treeMap[node.TreeNumber] = node
	if node.Children == nil{
		node.Children = make([]jianGoMeSHi.IdEntry, len(node.ChildrenMap))
		if node.Descriptor != nil{
			node.DescriptorUrl = BASE_URL + "/" + DESCRIPTOR + "/" + node.Descriptor.DescriptorUI
		}
	}
	i :=0
	for _,childNode := range node.ChildrenMap{
		node.Children[i].Id = childNode.TreeNumber
		node.Children[i].Url = BASE_URL + "/" + TREE + "/" + childNode.TreeNumber
		node.Children[i].Label = childNode.Name
		i++
	}
}


//sort slices

type ByIdX []jianGoMeSHi.IdEntry

type ById []*jianGoMeSHi.IdEntry

func (a ById) Len() int           { return len(a) }
func (a ByIdX) Len() int           { return len(a) }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByIdX) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ById) Less(i, j int) bool { return a[i].Id < a[j].Id }
func (a ByIdX) Less(i, j int) bool { return a[i].Id < a[j].Id }

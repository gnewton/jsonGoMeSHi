package main

import (
        "github.com/ant0ine/go-json-rest/rest"
        "log"
)

var COPYRIGHT_AND_LEGAL = MESH_COPYRIGHT + ";  " + MESH_OWNER

func GetAll(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	w.WriteJson(newEnvelope(allNouns))
}


func GetAllDescriptors(w rest.ResponseWriter, req *rest.Request) {
	log.Println("-------- ", req.PathParam("start"))
	if req.Request.Method != "GET"{
		return
	}
	w.WriteJson(newEnvelope(descSlice))
}



func GetDescriptor(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	descriptorUI := req.PathParam("id")

	//descriptor, ok := descMap[descriptorUI]
	_, ok := descMap2[descriptorUI]
	if ok{
		w.WriteJson(newEnvelope(descMap2[descriptorUI]))
	}else{
		rest.NotFound(w,req)
	}
}


func GetAllSupplementals(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	w.WriteJson(newEnvelope(suppSlice))
}


func GetSupplemental(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	supplementalUI := req.PathParam("id")
	
	supplemental, ok := suppMap[supplementalUI]
	if ok{
		w.WriteJson(newEnvelope(supplemental))
	}else{
		rest.NotFound(w,req)
	}

}


func GetAllQualifiers(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
		w.WriteJson(newEnvelope(qualSlice))
}


func GetQualifier(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	qualifierUI := req.PathParam("id")
	
	qualifier, ok := qualMap[qualifierUI]
	if ok{
		w.WriteJson(newEnvelope(qualifier))
	}else{
		rest.NotFound(w,req)
	}
}

func GetAllPharmacologicals(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	w.WriteJson(newEnvelope(pharmSlice))
}


func GetPharmacological(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	pharmUI := req.PathParam("id")
	
	pharm, ok := pharmMap[pharmUI]
	if ok{
		w.WriteJson(newEnvelope(pharm))
	}else{
		rest.NotFound(w,req)
	}
}


func GetTrees(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	//var nd *jianGoMeSHi.Node
	//child := nd.Init()
	//w.WriteJson(root)
	//log.Println(root)
	log.Printf("%+v\n", root)
	w.WriteJson(newEnvelope(root.Children))
	//log.Println(root.Children["D02"].Children["705"])
	//w.WriteJson("hello")
}


func GetTree(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	treeNumber := req.PathParam("a")
	
	if req.PathParam("b") != ""{
		treeNumber = treeNumber + "." + req.PathParam("b")
	}
	if req.PathParam("c") != ""{
		treeNumber = treeNumber + "." + req.PathParam("c")
	}
	if req.PathParam("d") != ""{
		treeNumber = treeNumber + "." + req.PathParam("d")
	}
	if req.PathParam("e") != ""{
		treeNumber = treeNumber + "." + req.PathParam("e")
	}
	if req.PathParam("f") != ""{
		treeNumber = treeNumber + "." + req.PathParam("f")
	}
	if req.PathParam("g") != ""{
		treeNumber = treeNumber + "." + req.PathParam("g")
	}
	if req.PathParam("h") != ""{
		treeNumber = treeNumber + "." + req.PathParam("h")
	}

	
	node, ok := treeMap[treeNumber]
	if ok{
		w.WriteJson(newEnvelope(node))
	}else{
		rest.NotFound(w,req)
	}
}

var meta =  Meta{
		CopyrightAndLegal: COPYRIGHT_AND_LEGAL,
		CopyrightAndLegalUrl: NLM_TERMS_URL,
		MeSH_Version: MESH_VERSION,
	}

func newEnvelope(record interface{}) *Envelope {
	env := new(Envelope)
	env.Meta = meta
	env.Data = record
	return env;
}
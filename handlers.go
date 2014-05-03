package main

import (
        "github.com/ant0ine/go-json-rest/rest"
        "log"
)

func GetAll(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	w.WriteJson(allNouns)
}


func GetAllDescriptors(w rest.ResponseWriter, req *rest.Request) {
	log.Println("-------- ", req.PathParam("start"))
	if req.Request.Method != "GET"{
		return
	}
	w.WriteJson(descSlice)
}


func GetDescriptor(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	descriptorUI := req.PathParam("id")

	//descriptor, ok := descMap[descriptorUI]
	_, ok := descMap[descriptorUI]
	if ok{
		//var descriptorWithUrl jianGoMeSHi.DescriptorRecord
		
		//descriptorWithUrl = *descriptor
		//w.WriteJson(populateWithUrl(descriptor, "http://localhost:8080/descriptor"))
		w.WriteJson(descMap2[descriptorUI])
	}else{
		rest.NotFound(w,req)
	}
}


func GetAllSupplementals(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	w.WriteJson(suppSlice)
}


func GetSupplemental(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	supplementalUI := req.PathParam("id")
	
	supplemental, ok := suppMap[supplementalUI]
	if ok{
		w.WriteJson(supplemental)
	}else{
		rest.NotFound(w,req)
	}

}


func GetAllQualifiers(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	w.WriteJson(qualSlice)
}


func GetQualifier(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	qualifierUI := req.PathParam("id")
	
	qualifier, ok := qualMap[qualifierUI]
	if ok{
		w.WriteJson(qualifier)
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
	w.WriteJson(root.Children)
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
	
	node, ok := treeMap[treeNumber]
	if ok{
		w.WriteJson(node)
	}else{
		rest.NotFound(w,req)
	}
}

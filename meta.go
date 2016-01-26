package main

type Meta struct {
	MeSH_Version         string
	CopyrightAndLegal    string
	CopyrightAndLegalUrl string
	Notes                string `json:",omitempty"`
}

type Envelope struct {
	Meta Meta
	Data interface{}
}

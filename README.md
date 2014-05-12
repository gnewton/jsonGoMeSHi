jsonGoMeSHi
===========

JSON REST Interface for MeSH data based on [jianGoMeSHi][https://github.com/gnewton/jainGoMeSHi]


Example JSON
============

* Base URL
    {
    
        "Meta": {
            "MeSH_Version": "2014 MeSH",
            "CopyrightAndLegal": "Copyright U.S. National Library of Medicine; U.S. National Library of Medicine is the creator, maintainer, and provider of this data",
            "CopyrightAndLegalUrl": "https://www.nlm.nih.gov/mesh/termscon.html"
        },
        "Data": [
            {
                "Id": "/descriptor",
                "Url": "http://s2.semanticscience.org:8080/mesh/descriptor"
            },
            {
                "Id": "/qualifier",
                "Url": "http://s2.semanticscience.org:8080/mesh/qualifier"
            },
            {
                "Id": "/supplemental",
                "Url": "http://s2.semanticscience.org:8080/mesh/supplemental"
            },
            {
                "Id": "/tree",
                "Url": "http://s2.semanticscience.org:8080/mesh/tree"
            },
            {
                "Id": "/pharmacological",
                "Url": "http://s2.semanticscience.org:8080/mesh/pharmacological"
            }
        ]
    
    }
* Descriptor Record


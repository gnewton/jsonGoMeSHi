jsonGoMeSHi
===========

JSON REST Interface for MEDLINE/PubMed Medical Subject Headings (MeSH) XML data based on [jianGoMeSHi](https://github.com/gnewton/jianGoMeSHi)


Web Service API
============
* Base URL:
  Lists all REST nouns: `descriptor`, `qualifier`, `supplemental`, `pharmacological`, `tree`
* For each of noun: `descriptor`, `qualifier`, `supplemental`, `pharmacological`, behaviour is as follows:
  * `http://hostname:port/mesh/noun` - list all records for `noun`
  * `http://hostname:port/mesh/noun/ID` - list record ID for `noun`
* Note that descriptor and supplemental record lists are quite large: there is a [bug to add paging](https://github.com/gnewton/jsonGoMeSHi/issues/1) to these API calls

Example JSON
============

Note all URLs are examples only: the base URL is settable in jsonGoMeSHi

* Base URL  http://hostname:port/mesh

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
  * See [here](https://github.com/gnewton/jsonGoMeSHi/blob/master/exampleJson/descriptor.json)

* Qualifier Record
  * See [here](https://github.com/gnewton/jsonGoMeSHi/blob/master/exampleJson/qualifier.json)

* Supplemental Record
  * See [here](https://github.com/gnewton/jsonGoMeSHi/blob/master/exampleJson/supplemental.json)

* Pharmacological action Record
```
     {
         "Meta": {
             "MeSH_Version": "2014 MeSH",
             "CopyrightAndLegal": "Copyright U.S. National Library of Medicine; U.S. National Library of Medicine is the creator, maintainer, and provider of this data",
             "CopyrightAndLegalUrl": "https://www.nlm.nih.gov/mesh/termscon.html"
         },
         "Data": {
             "DescriptorReferredTo": {
                 "DescriptorUI": "D000020",
                 "DescriptorName": "Abortifacient Agents, Nonsteroidal",
                 "Url": "http://s2.semanticscience.org:8080/mesh/descriptor/D000020"
             },
             "PharmacologicalActionSubstanceList": {
                 "Substance": [
                     {
                         "RecordUI": "C030266",
                         "SupplementalUrl": "http://s2.semanticscience.org:8080/mesh/supplemental/C030266",
                         "RecordName": "3-(2-ethylphenyl)-5-(3-methoxyphenyl)-1H-1,2,4-triazole"
                     },
                     {
                         "RecordUI": "D002260",
                         "DescriptorUrl": "http://s2.semanticscience.org:8080/mesh/descriptor/D002260",
                         "RecordName": "Carboprost"
                     },
                     {
                         "RecordUI": "D015237",
                         "DescriptorUrl": "http://s2.semanticscience.org:8080/mesh/descriptor/D015237",
                         "RecordName": "Dinoprost"
                     },
                     {
                         "RecordUI": "C010714",
                         "SupplementalUrl": "http://s2.semanticscience.org:8080/mesh/supplemental/C010714",
                         "RecordName": "dinoprost tromethamine"
                     },
                     {
                         "RecordUI": "C039153",
                         "SupplementalUrl": "http://s2.semanticscience.org:8080/mesh/supplemental/C039153",
                         "RecordName": "fenprostalene"
                     },
                     {
                         "RecordUI": "C021182",
                         "SupplementalUrl": "http://s2.semanticscience.org:8080/mesh/supplemental/C021182",
                         "RecordName": "gemeprost"
                     },
                     {
                         "RecordUI": "C025505",
                         "SupplementalUrl": "http://s2.semanticscience.org:8080/mesh/supplemental/C025505",
                         "RecordName": "meteneprost"
                     },
                     {
                         "RecordUI": "D008727",
                         "DescriptorUrl": "http://s2.semanticscience.org:8080/mesh/descriptor/D008727",
                         "RecordName": "Methotrexate"
                     },
                     {
                         "RecordUI": "D016595",
                         "DescriptorUrl": "http://s2.semanticscience.org:8080/mesh/descriptor/D016595",
                         "RecordName": "Misoprostol"
                     },
                     {
                         "RecordUI": "C039582",
                         "SupplementalUrl": "http://s2.semanticscience.org:8080/mesh/supplemental/C039582",
                         "RecordName": "MMC protein, Momordica charantia"
                     },
                     {
                         "RecordUI": "C002443",
                         "SupplementalUrl": "http://s2.semanticscience.org:8080/mesh/supplemental/C002443",
                         "RecordName": "neem oil"
                     },
                     {
                         "RecordUI": "C016767",
                         "SupplementalUrl": "http://s2.semanticscience.org:8080/mesh/supplemental/C016767",
                         "RecordName": "sulprostone"
                     },
                     {
                         "RecordUI": "D015978",
                         "DescriptorUrl": "http://s2.semanticscience.org:8080/mesh/descriptor/D015978",
                         "RecordName": "Trichosanthin"
                     }
                 ]
             }
         }
     
     }
```


Acknowledgement
=============
This work is a by-product of my graduate work at Carleton Univerity at [Dumontier Lab](http://dumontierlab.com/)
import threading
import requests
import json

def Snake2BigGolangCase(raw_str):
    # saving first and rest using split()
    init, *temp = raw_str.split('_')
    # using map() to get all words other than 1st
    # and titlecasing them
    res = ''.join([init[0:1].upper(), init.lower()[1:], *map(str.title, temp)])

    total = len(res)
    if total >= 2:
        if (res[total-2:]).lower() == "id":
            res = res[:total-2] + "ID"
    return res

typeMapping = {
    "string": "string",
    "string[]": "[]string",
    "object": "interface{}", # sub
    "object[]" : "[]interface{}",
    "int" : "int",
    "int[]" : "[]int",
    "boolean" : "bool",
    "timestamp" : "int",
    "float" : "float64",
    "file" : "*multipart.FileHeader",
    "list": "" # sub
}


# store global object name to share
globalStruct = []

# store global interface usage 
globalInterface = []

def ParseObjectToStruct(structName, raw):
    globalStruct.append(structName)

    # store another structs
    dependenciesStructRaw = ""
    
    # parent structs
    structRaw = ("""
//=======================================================
// Object Raw Type - %s
//=======================================================
type %s struct {""" % (structName, structName))
            
    for item in raw["children"]:               
        # omitempty
        omemptyStr = ""
        try:
            if item["required"] == "False":
                omemptyStr = ",omitempty"
        except:
            print("this object does not support required field, skipped...\n")

        
        # float convert to string
        toStrstr = ""
        if item["type"] == "flaot":
            toStrstr = ",string"
        
        if item["type"] == "file":
            structRaw += """
// %s is a filetype, should parse by http.request
// File *multipart.FileHeader `form:"file" binding:"required"`
// Using `ShouldBind`
// --------------------
// var form Form
// _ := c.ShouldBind(&form)

// Get raw file bytes - no reader method
// openedFile, _ := form.File.Open()
// file, _ := ioutil.ReadAll(openedFile)

// Upload to disk
// `form.File` has io.reader method
// c.SaveUploadedFile(form.File, path)
// --------------------

// Using `FormFile`
// --------------------
// formFile, _ := c.FormFile("file")

// Get raw file bytes - no reader method
// openedFile, _ := formFile.Open()
// file, _ := ioutil.ReadAll(openedFile)

// Upload to disk
// `formFile` has io.reader method
// c.SaveUploadedFile(formFile, path)
// --------------------
""" % (item["name"])
            continue

        if item["type"] == "list" or item["type"] == "object":
            subObjName = Snake2BigGolangCase(item["name"])
            if not subObjName in globalStruct:
                globalStruct.append(objName)
                dependenciesStructRaw += ParseObjectToStruct(subObjName, item)
            
            if item["type"] == "list":
                subObjName = '[]' + subObjName

            structRaw += """
// %s is %s
%s %s `json:\"%s%s%s\"`""" % (
                item["name"],
                item["description"],
                Snake2BigGolangCase(item["name"]),
                subObjName,
                item["name"],
                omemptyStr,
                toStrstr
            )
            continue

        structRaw += """
// %s is %s
%s %s `json:\"%s,omitempty%s\"`""" % (
            item["name"],
            item["description"],
            Snake2BigGolangCase(item["name"]),
            TypeMappingGolangType(item["type"]),
            item["name"],
            toStrstr
        )
    structRaw += ("""
}""")

    return dependenciesStructRaw + "\n\n" + structRaw


def TypeMappingGolangType(type_name):
    if not type_name in typeMapping:
        print(type_name + " is not exists in type mapping, plz manually add it")

    # if list
    return typeMapping[type_name]

if __name__ == "__main__":
    manualResp = requests.get("https://open.shopee.com/api/v1/doc/module/?version=2&SPC_CDS=423b402d-921f-401c-844b-98a3e724bf8a&SPC_CDS_VER=2")
    manual = manualResp.json()

    # Skip Overview
    skip_module = [ 87 ]

    # Path Mapping
    availablePaths = []
    
    for module in manual['modules']:
        if module["module_id"] in skip_module:
            continue

        filename = "model-%s.go" % (module["module_name"].lower())
        f = open(filename, "a", encoding="utf-8")
        f.write("package shopeego")

        for item in module["items"]:
            apiManualResp = requests.get("https://open.shopee.com/api/v1/doc/api/?SPC_CDS=2cd9fdcd-1b7e-4708-891d-a7b5206b160c&SPC_CDS_VER=2&version=2&api_name=%s" % (item["name"]))
            apiManual = apiManualResp.json()

            apiPaths = apiManual["api_name"].split(".")
            apiName = Snake2BigGolangCase(apiPaths[len(apiPaths)-2] + "_" + apiPaths[len(apiPaths)-1])
            
            availablePaths.append({
                "name" : apiName,
                "path" : apiManual["path"]
            })

            apiParams = json.loads(apiManual["params"])

            # add to global interface list:
            globalInterface.append(apiName)
            
            reqStr = ("""
//=======================================================
// %sRequest
//=======================================================
type %sRequest struct {""" % (apiName, apiName))

            for reqItem in apiParams["request_params"]:
                # omitempty
                omemptyStr = ""
                if reqItem["required"] == "False":
                    omemptyStr = ",omitempty"
                
                # float convert to string
                toStrstr = ""
                if reqItem["type"] == "flaot":
                    toStrstr = ",string"
                
                if reqItem["type"] == "file":
                    resStr += """
// %s is a filetype, should parse by http.request

""" % (resItem["name"])
                    continue
                    
                # DRY:1 
                if reqItem["type"] == "list" or reqItem["type"] == "object":
                    objName = Snake2BigGolangCase(reqItem["name"])
                    if not objName in globalStruct:
                        globalStruct.append(objName)
                        f.write(ParseObjectToStruct(objName, reqItem))
                                       
                    reqStr += """
    // %s is %s
    %s %s `json:\"%s%s%s\"`""" % (
                    reqItem["name"],
                    reqItem["description"],
                    Snake2BigGolangCase(reqItem["name"]),
                    objName,
                    reqItem["name"],
                    omemptyStr,
                    toStrstr
                )

                    continue #DRY END

                reqStr += """
    // %s is %s
    %s %s `json:\"%s%s%s\"`""" % (
                    reqItem["name"],
                    reqItem["description"],
                    Snake2BigGolangCase(reqItem["name"]),
                    TypeMappingGolangType(reqItem["type"]),
                    reqItem["name"],
                    omemptyStr,
                    toStrstr
                )
            reqStr += ("""
}""")
            resStr = ("""
//=======================================================
// %sResponse
//=======================================================
type %sResponse struct {""" % (apiName, apiName))
            
            for resItem in apiParams["response_params"]:               
                # float convert to string
                toStrstr = ""
                if resItem["type"] == "flaot":
                    toStrstr = ",string"
                
                if resItem["type"] == "file":
                    resStr += """
    // %s is a filetype, should parse by http.request
""" % (resItem["name"])
                    continue
                    
                # DRY:2 
                if resItem["type"] == "list" or resItem["type"] == "object":
                    objName = Snake2BigGolangCase(resItem["name"])
                    if not objName in globalStruct:
                        globalStruct.append(objName)
                        f.write(ParseObjectToStruct(objName, resItem))
                                       
                    resStr += """
    // %s is %s
    %s %s `json:\"%s%s%s\"`""" % (
                    resItem["name"],
                    resItem["description"],
                    Snake2BigGolangCase(resItem["name"]),
                    objName,
                    resItem["name"],
                    omemptyStr,
                    toStrstr
                )
                    continue #DRY END

                resStr += """
    // %s is %s
    %s %s `json:\"%s,omitempty%s\"`""" % (
                    resItem["name"],
                    resItem["description"],
                    Snake2BigGolangCase(resItem["name"]),
                    TypeMappingGolangType(resItem["type"]),
                    resItem["name"],
                    toStrstr
                )
            resStr += ("""
}""")

            print(reqStr)
            f.write(reqStr)
            print(resStr)
            f.write(resStr)
        f.close()
    
    # urls.go
    urlStrRaw = """
package shopeego
var availablePaths map[string]string = map[string]string{
"""

    for urlPair in availablePaths:
        urlStrRaw += """
    "%s": "%s",
""" % (urlPair["name"], urlPair["path"])
    
    urlStrRaw += """
}
"""
    f = open("urls.go", "a", encoding="utf-8")
    f.write(urlStrRaw)
    f.close()

    # generate interface and implements
    globalInterfaceStr = """
package shopeego

type V2I interface {
"""
    globalImplementStr = """
package shopeego

import "encoding/json"
"""

    for item in globalInterface:
        globalInterfaceStr += """
    %s(*%sRequest) (*%sResponse, error)
""" % (item, item, item)

        globalImplementStr += """

func (s *ShopeeClient) %s(req *%sRequest) (resp *%sResponse, err error) {
	b, err := s.post("%s", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}
""" % (item, item, item, item)

    globalInterfaceStr += """
}
"""
    f = open("v2i.go", "a", encoding="utf-8")
    f.write(globalInterfaceStr)
    f.close()

    f = open("implv2.go", "a", encoding="utf-8")
    f.write(globalImplementStr)
    f.close()
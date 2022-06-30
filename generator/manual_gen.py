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
    "object": "interface{}",
    "object[]" : "[]interface{}",
    "int" : "int",
    "int[]" : "[]int",
    "boolean" : "bool",
    "timestamp" : "int",
    "float" : "float64",
}

def TypeMappingGolangType(type_name):
    if not type_name in typeMapping:
        print(type_name + " is not exists in type mapping, plz manually add it")
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
            




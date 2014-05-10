package main

import (
   "github.com/codegangsta/martini"
   "net/http"
   "strings"
   "encoding/json"
)

//fake data for now

type Attribute struct {
  Names string "json:name"
  DataType string "json:type"
  Description string "json:description"
  Required bool "json:required"
}

type ResourceAttributes struct {
  ResourceName string "json: resourceName"
  Attributes []Attribute "json: attributes"
}

type ErrorMsg struct {
  Msg string "json: msg"
}

//will maybe return a json data

func (ra ResourceAttributes) String() (s string) {
  jsonObj, err := json.Marshal(ra)
  if err != nil {
    s = ""
  }else{
    s = string(jsonObj)
  }
  return
}

type jsonConvertible interface { }

func JsonString( obj jsonConvertible ) (s string) {
  jsonObj, err := json.Marshal( obj )

  if err != nil {
    s = ""
  } else {
    s = string( jsonObj )
  }

  return
}

func main() {
  m := martini.Classic()
  //just for fun
  m.Get("/", func() string {
    return "where is all za people?"
  })

  //get attr for a given resources
  m.Get("/attr/:resources", func( params martini.Params, writer http.ResponseWriter) (int, string) {
    resources := strings.ToLower(params["resources"])
    writer.Header().Set("Content-Type", "application/json")
    if resources == "tv" {
      resourceAttrs := ResourceAttributes{"tv", make([]Attribute, 1)}
      resourceAttrs.Attributes[0] = Attribute{"Location","string", "What facility is the TV located in.", true}
      return http.StatusOK, JsonString( resourceAttrs )
    }else{
      return http.StatusNotFound, JsonString( ErrorMsg{"Resource not found: " + resources} )
    }
  })
  m.Run()
}

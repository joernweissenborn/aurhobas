package src

import "github.com/joernweissenborn/aursir4go/appkey"

var ItemAppkey = appkey.AppKeyFromYaml(`

applicationkeyname: aurhobas.item

functions:

  - name: GetId
    output:
      - name: Id
        type: 1

 - name: GetName
    output:
      - name: Name
        type: 1

  - name: GetFiles
    output:
      - name: Files
        type: 10

  - name: GetFileDetails
    input:
      - name: File
        type: 1
    output:
      - name: Name
        type: 1
      - name: Type
        type: 1
      - name: Tags
        type: 10

`)

type StoreItemRequest struct {
	Name string
	Type string
	Tags []string
	Data string
}

type StoreItemReply struct {
	Id string
}

var ItemAppkey = appkey.AppKeyFromYaml(`

applicationkeyname: AurHobas.Item

functions:

  - name: GetId
    output:
      - name: Id
        type: 1
  - name: GetMeta
    output:
      - name: Name
        type: 1
      - name: Type
        type: 1
      - name: Tags
        type: 10
  - name: GetData
    input:
      - name: Id
        type: 1
    output:
      - name: Data
        type: 1


`)

type GetIdReply struct {
	Id string
}

type GetMetaReply struct {
	Name string
	Type string
	Tags []string
}
type GetDataReply struct {
	Data string
}

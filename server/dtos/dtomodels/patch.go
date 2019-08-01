package dtomodels

type Patch struct {
	Operation string `json:"op,required"`
	Path      string `json:"path,required"`
	Value     string `json:"value"`
}

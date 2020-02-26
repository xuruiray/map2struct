# map2struct
map2struct

map convert to struct

for example
```
  map[string]string{
    "num": "1",
    "str": "asd",
    "arr": "[123,123,123]",
  }
  
  type Temp struct {
    Num int    `json:"num"`
    Str string `json:"str"`
    Arr []int  `json:"arr"`
  }
```


# map2struct
map2struct

map convert to struct
bind by json tag and map key

for example
```
  map[string]string{
    "num": "1",
    "str": "asd",
    "arr": "[123,123,123]",
    "str_arr": "[\"123\",\"123\"]",
  }
  
  type Temp struct {
    Num int    `json:"num"`
    Str string `json:"str"`
    Arr []int  `json:"arr"`
    StrArr []string `json:"str_arr"`
  }
```


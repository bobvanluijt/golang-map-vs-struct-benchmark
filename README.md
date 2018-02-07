# Generate json files

Run: `node --max_old_space_size=8183 generateJson.js`
Output: `Id to find:  p_xxx` (replace xxx with the actual number)

# Test 

Test maps: `go run ./map/speedTest.go p_xxx`

Test structs: `go run ./struct/speedTest.go p_xxx`

# Struct:

```
type PropertiesA []struct {
	History []struct {
		Timestamp int64  `json:"timestamp"`
		Value     string `json:"value"`
	} `json:"history"`
	Property string `json:"property"`
}
```

# Map:

```
type propertiesB map[string]map[string]string
```

# Results

General balance history and items (resp: 125000, max 1100):

```
MAP QUERY FOUND:  low4vvu5abdo2sh2ps5bwewmi
MAP QUERY TOOK:  4955000

STRUCT QUERY FOUND:  low4vvu5abdo2sh2ps5bwewmi
STRUCT QUERY TOOK:  1222000 (WIN)
```

Many items, few history items (resp: 500000, max 15):

```
MAP QUERY FOUND:  oe2fmjop1rql0etxj2974aemi
MAP QUERY TOOK:  34858000

STRUCT QUERY FOUND:  oe2fmjop1rql0etxj2974aemi
STRUCT QUERY TOOK:  5360000 (WIN)
```

Few items, many history items (resp: 5, +/- 525960)

```
MAP QUERY FOUND:  acipbh9dozyqrnjeam485gsyvi
MAP QUERY TOOK:  53729000

STRUCT QUERY FOUND:  acipbh9dozyqrnjeam485gsyvi
STRUCT QUERY TOOK:  1107000 (WIN)
```
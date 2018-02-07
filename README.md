# Generate json files

Run: `node --max_old_space_size=8183 generateJson.js`
Output: `Id to find:  prop_xxx` (replace xxx with the actual number)

# Test 

Test: `go run speedTest.go prop_xxx`

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
STRUCT QUERY FOUND: r4qcg9kzfzrexaygedxdvaemi
STRUCT QUERY TOOK:  1012000 (WIN)
MAP QUERY FOUND:  r4qcg9kzfzrexaygedxdvaemi
MAP QUERY TOOK:  1288000
```

Many items, few history items (resp: 100000, max 15):

```
STRUCT QUERY FOUND: by5a38ok0pw6c6vhw2n265hfr
STRUCT QUERY TOOK:  12353000
MAP QUERY FOUND:  by5a38ok0pw6c6vhw2n265hfr
MAP QUERY TOOK:  11555000 (WIN)
```

Few items, many history items (resp: 5, +/- 525960)

```
STRUCT QUERY FOUND: dr1h0ok5nj8dxxibbr5mjf9a4i
STRUCT QUERY TOOK:  974000 (WIN)
MAP QUERY FOUND:  dr1h0ok5nj8dxxibbr5mjf9a4i
MAP QUERY TOOK:  1984000
```
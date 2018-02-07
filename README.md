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

```
STRUCT QUERY FOUND:  3lpes4oqj63f9826vs3sdcxr
STRUCT QUERY TOOK:  1217000
INTERFACE QUERY FOUND:  3lpes4oqj63f9826vs3sdcxr
INTERFACE QUERY TOOK:  1170000
```
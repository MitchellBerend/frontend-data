package files

import _ "embed"

// This file contains constants that are byte representations of data files

//go:embed index.json
var IndexJson []byte

//go:embed over-engineering.json
var OverEngineeringJson []byte

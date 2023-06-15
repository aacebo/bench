package container

import (
	"encoding/base64"
	"fmt"
)

func userData(id string, lang string, problem string) *string {
	v := fmt.Sprintf(`
	#!/bin/bash

	export BENCH_ID="%s"
	export BENCH_LANG="%s"
	export BENCH_PROBLEM="%s"
	`, id, lang, problem)

	b64 := base64.StdEncoding.EncodeToString([]byte(v))
	return &b64
}

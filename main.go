package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

type FeatureFlags map[string]string

func(f FeatureFlags) ToSlice() []string {
	if len(f) == 0 {
		return nil
	}
	result := make([]string, 0, len(f))
	for k, v := range f {
		result = append(result, fmt.Sprintf("%s=%s", k, v))
	}
	return result
}

func main() {
	var fmap map[string]string
	flag.StringToStringVar(&fmap, "feature-flag", fmap, "list of feature flags")
	flag.Parse()

	ff := FeatureFlags(fmap)
	fmt.Println(ff.ToSlice())
}

package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
	"gomodules.xyz/errors"
	"gomodules.xyz/sets"
)

var knownFlags = sets.NewString("DisableAnalytics")

type FeatureFlags map[string]string

func (f FeatureFlags) IsValid() error {
	var errs []error
	for k := range f {
		errs = append(errs, fmt.Errorf("unknown feature flag %q", k))
	}
	return errors.NewAggregate(errs)
}

func (f FeatureFlags) ToSlice() []string {
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
	if err := ff.IsValid(); err != nil {
		panic(err)
	}
	fmt.Println(ff.ToSlice())
}

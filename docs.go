// Package testament provides convinience tools in testing.
//
// AssertInError
//
// If you have a deeply nested error you want to check whether a particular
// error is in the haystack you can call this function. It will upwrap
// everything made from any of or any combination of fmt.Errorf,
// hashicorp/go-multierror, and pkg/errors helpers. It will fallback to the
// string representation of the error being in the haystack string
// representation.
package testament

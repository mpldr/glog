// glog aims to be the "one for all" logging package. It has a wide variety of
// features.
// To "handle errors gracefully" most times when a Setting-Function is called,
// no error is returned. If the specified level does not exist, execution is
// just silently aborted.
//
// An important note is that *all* log-level ranges are inclusive. This means
// that "below …" actually means "below or equal to …".
package glog

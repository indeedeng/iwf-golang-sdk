package iwf

/**
 * return a struct pointer of the input type
 * this is needed for deserializing data into the input object for execute/decide API before invoking them
 * e.g. for an integer type, implementation could be
 *   {
 *       var i int
 *       return &i
 *   }
 * use nil for empty input(no need to implement this)
 * TODO: think of a better way to do it in Golang, maybe using generic?
 */
type NewTypePtr func() interface{}
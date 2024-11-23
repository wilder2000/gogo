package server

//func RegMapping[M any](c sm.HTTPController[M]) {
//	ctrl := newController(c)
//	mappings[c.UrlPath()] = ctrl.Prepare
//}
//func newController[M any](c sm.HTTPController[M]) *sm.AbstractController[M] {
//	ty := reflect.ValueOf(c)
//	fi := ty.Elem().FieldByName("AbstractController")
//	if fi.Type().ConvertibleTo(reflect.TypeOf(sm.AbstractController[M]{})) {
//		cc := fi.Interface().(sm.AbstractController[M])
//		cc.HTTPController = c
//		return &cc
//	}
//	return nil
//}

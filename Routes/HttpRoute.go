package Routes

/*
HTTP的消息过来时，每一个请求都是一个新的协程;
所以如果要走actor模块，就要把消息按一定规则转到指定的协程上运行；
但有时候，也可以直接处理，只要做好协程安全就可以了；
**/

// func init() {
// 	HTTPRoutelist = make(map[int]event.HTTPcall)
// 	HTTPRoutelist[ActionCode.Http_User_Test] = HttpEvents.DemoEvent
// }

// var (
// 	HTTPRoutelist map[int]event.HTTPcall
// )

// func HTTPInit() {
// 	Service.HTTPEx.RouteFun = HTTPRoute
// 	Service.HTTPEx.GetIPFun = GetIP
// }

// func HTTPRoute(code int) event.HTTPcall {
// 	f, ok := HTTPRoutelist[code]
// 	if ok {
// 		return f
// 	}
// 	return nil
// }

// func GetIP(w http.ResponseWriter, req *http.Request) string {
// 	//这个方法是用来拿IP的，因为会被https代理，所以RemoteAddr不一定拿到客户的IP；
// 	//所以与你自己的运营沟通一下看看在哪里可以拿到IP；
// 	if ips, ok := req.Header["X-Forwarded-For"]; ok {
// 		if len(ips) > 0 {
// 			return ips[0]
// 		}
// 	}
// 	return req.RemoteAddr
// }

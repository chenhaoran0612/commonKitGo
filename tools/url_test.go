package tools

import "testing"

func TestIsValidUrl(t *testing.T) {
	t.Log(IsValidUrl("http://ww.com"))                                                                                                         // true
	t.Log(IsValidUrl("http://"))                                                                                                               // false
	t.Log(IsValidUrl("https://"))                                                                                                              // false
	t.Log(IsValidUrl("https://w.com"))                                                                                                         // true
	t.Log(IsValidUrl("http://ethapi.test.poolx.io:8081/watchDashboard/subToken=NzMzMTkxMmY3NjUyYThkNmNjNDM4ZDY5NzY4NTlmZWZ8QlRDfEpheFBvb2w=")) // false
	t.Log(IsValidUrl("/foo/bar"))                                                                                                              // false
}

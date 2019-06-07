package main

const (
	TenantID     = "2c"
	ClientID     = "86"
	ClientSecret = "K"
)

func main() {
	var SitesParam = map[string]interface{}{
		//"ms_profile": map[string]interface{}{
		//	"tenant_id":     TenantID,
		//	"client_id":     ClientID,
		//	"client_secret": ClientSecret,
		//},
		"ms_sources": []interface{}{
			"sites",
			"groups",
		},
	}
	p_siteParam := &SitesParam
	println("address of sitesParam: ", p_siteParam)

	b := *p_siteParam
	println("siteParam:", b)
	println("siteParam:", SitesParam)

	//ms_sources:=SitesParam["ms_sources"].([]interface{})
	//ms:=SitesParam["ms_sources"]
	//println("ms_sources:",ms_sources,"ms",ms)
	gr := SitesParam["ms_sources"].([]interface{})[1].(string)

	gr_p := &gr

	p_ms := &SitesParam["ms_sources"]

	*gr_p = "tttt"
	p_ms = "sdf"

}

package main

import (
	"encoding/json"
	"fmt"

	"github.com/huangapple/umeng-openapi-sdk/services/uapp"
)

var (
	apiKey      = "7785319"
	apiSecurity = "8KIcW4Zd37b"
	appKey      = "5fe567c9adb42d58268e0f73"
)

func DumpObj(i interface{}, perfect bool) {

	var bytes []byte
	var err error
	if perfect {
		bytes, err = json.MarshalIndent(i, "", "    ")

	} else {
		bytes, err = json.Marshal(i)
	}

	fmt.Printf("%s %v\n", bytes, err)
}

func main() {
	uAppCli := uapp.NewUAppClient(apiKey, apiSecurity)
	allAppDataResp, err := uAppCli.GetAllAppData()
	if err != nil {
		fmt.Println(err)
		return
	}
	opt, _ := json.MarshalIndent(allAppDataResp, "", "    ")
	fmt.Println(string(opt))

	//newAccountsResp, err := uAppCli.GetNewAccounts(appKey, "2020-05-01", "2020-05-11", "daily", "")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//opt, _ = json.MarshalIndent(newAccountsResp, "", "    ")
	//fmt.Println(string(opt))

	info, err := uAppCli.GetDurations(appKey, "2021-04-09", "daily_per_launch", "", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	DumpObj(info, true)
}

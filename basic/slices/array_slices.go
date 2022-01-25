package main

import "fmt"

type KAWechatInfo struct {
	DmError  int    `json:"dm_error"`
	ErrorMsg string `json:"error_msg"`
	Data     struct {
		KaStatus      int         `json:"ka_status"`
		KaLevel       int         `json:"ka_level"`
		ImgType       int         `json:"img_type"`
		Title         string      `json:"title"`
		Subtitle      string      `json:"subtitle"`
		Code          string      `json:"code"`
		H5URL         string      `json:"h5_url"`
		ImageURL      string      `json:"image_url"`
		BannerURL     string      `json:"banner_url"`
		Message       string      `json:"message"`
		RawCode       string      `json:"raw_code"`
		OfficialCode  string      `json:"official_code"`
		ImageMediaID  string      `json:"image_media_id"`
		WelfwareTitle string      `json:"welfware_title"`
		Welfware      interface{} `json:"welfware"`
		InstantTitle  string      `json:"instant_title"`
		Jump          struct {
			Path string `json:"path"`
			ID   string `json:"id"`
			Type int    `json:"type"`
		} `json:"jump"`
		TaskCenter struct {
			Img  string `json:"img"`
			Text string `json:"text"`
		} `json:"task_center"`
		InstantAward interface{} `json:"instant_award"`
		AddWay       int         `json:"add_way"`
	} `json:"data"`
}

func main() {

	//var arr1 [5]int
	//var slice1 = arr1[1:3]
	//
	//// init
	//for i := 0; i < len(arr1); i++ {
	//	arr1[i] = i
	//}
	//
	//// slice1
	//for i := 0; i < len(slice1); i++ {
	//	fmt.Println(slice1[i])
	//}
	//
	//fmt.Printf("len(arr1): %d\n", len(arr1))
	//fmt.Printf("len(slice1): %d\n", len(slice1))
	//fmt.Printf("cap(slice1): %d\n", cap(slice1))

	sl := make([]int, 2)
	fmt.Println(sl)

	for i := 0; i < len(sl); i++ {
		if i == len(sl)-1 { // 最后一个
			sl = append(sl[:i], sl[i+1:]...)
		}
	}

	fmt.Println(sl)

	var res KAWechatInfo
	fmt.Printf("%+v\n", res)

}

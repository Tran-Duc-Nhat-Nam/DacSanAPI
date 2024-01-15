package routes

// func ThemTinhThanhJson() {
// 	res, err := http.Get("https://provinces.open-api.vn/api/?depth=1")
// 	if err != nil {
// 		return
// 	}
// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return
// 	}
// 	var items []map[string]interface{}
// 	err = json.Unmarshal(body, &items)
// 	if err != nil {
// 		return
// 	}
// 	for _, value := range items {
// 		fmt.Println(value["code"])
// 		fmt.Println(value["name"])
// 		ThemTinhThanhCSDL(int(value["code"].(float64)), value["name"].(string))
// 	}
// 	defer res.Body.Close()
// }

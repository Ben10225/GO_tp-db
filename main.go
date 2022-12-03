package main

import (
	"encoding/json"
	"fmt"
	. "go_db/components"
	. "go_db/data"
	"io/ioutil"
	"strings"
)

func main() {
	insertAll()
}

func insertAll() {
	data, _ := ioutil.ReadFile("data/taipei-attractions.json")

	var result ResultT
	json.Unmarshal([]byte(data), &result)
	// fmt.Println(result.Result.Results)

	details := result.Result.Results

	db := InitDb()

	catMap := map[string]int{}
	cct := 0

	mrtMap := map[string]int{}
	mct := 0
	// var clst []string

	for _, v := range details {
		/*
			  // use struct
				cat := struct{ Category_name string }{v.Category}
				// cat := Cat{Category_name: v.Category}
				db.Table("categories").Create(cat)
		*/

		// category
		if catMap[v.Category] == 0 {
			cct++
			catMap[v.Category] = cct

			db.Table("categories").Create(map[string]interface{}{
				"category_name": v.Category,
			})
		}
		// if !slices.Contains(clst, v.Category) {
		// 	clst = append(clst, v.Category)
		// 	cct++
		// 	catMap[v.Category] = cct
		// }

		// mrt
		if mrtMap[v.Mrt] == 0 {
			mct++
			mrtMap[v.Mrt] = mct

			db.Table("mrts").Create(map[string]interface{}{
				"mrt_name": v.Mrt,
			})
		}

		// attractions
		db.Table("attractions").Create(map[string]interface{}{
			"id":          v.Id,
			"name":        v.Name,
			"category_id": catMap[v.Category],
			"description": v.Description,
			"address":     v.Address,
			"transport":   v.Transport,
			"mrt_id":      mrtMap[v.Mrt],
			"lat":         v.Lat,
			"lng":         v.Lng,
		})

		// images
		imgs := strings.Split(v.Imgs, "https")
		// var imgsNew []string
		for j, imgurl := range imgs {
			if j > 0 {
				if strings.ToLower(imgurl[len(imgurl)-3:]) == "jpg" || strings.ToLower(imgurl[len(imgurl)-3:]) == "png" {
					// imgsNew = append(imgsNew, "https"+imgurl)
					db.Table("images").Create(map[string]interface{}{
						"iid": v.Id,
						"url": "https" + imgurl,
					})
				}
			}
		}
	}
	fmt.Println("新增資料完成")
}

package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Items []Item

func itemList(w http.ResponseWriter, r *http.Request) {
		items := Items{}
		items = append(items, Item{Id: 1, Name: "MU-TON"})
		items = append(items, Item{Id: 2, Name: "SAM"})
		items = append(items, Item{Id: 3, Name: "ID"})

		fmt.Println("Access: /item/list")
		json.NewEncoder(w).Encode(items)
}

func main() {
	http.HandleFunc("/item/list", itemList)
	http.ListenAndServe(":8000", nil)
}
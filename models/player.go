package models

type Player struct {
	Id   int    `json:"id"`
	Ign  string `json:"ign"`
	Name string `json:"name"`
	Age  *int   `json:"age"`
}

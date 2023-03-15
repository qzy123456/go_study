package main

import (
	"fmt"
)

//定义菜单结构体
type Menu struct {
	ID       int
	Name     string
	ParentID int
	Children []*Menu
}

//递归生成菜单列表
func GenerateMenuList(items []*Menu, parentID int) []*Menu {
	var menuList []*Menu
	for _, item := range items {
		if item.ParentID == parentID {
			children := GenerateMenuList(items, item.ID)
			item.Children = children
			menuList = append(menuList, item)
		}
	}
	return menuList
}

func main() {
	//定义菜单项
	menuItems := []*Menu{
		&Menu{ID: 1, Name: "Home", ParentID: 0},
		&Menu{ID: 2, Name: "About", ParentID: 0},
		&Menu{ID: 3, Name: "Contact", ParentID: 0},
		&Menu{ID: 4, Name: "Products", ParentID: 0},
		&Menu{ID: 5, Name: "Clothes", ParentID: 4},
		&Menu{ID: 6, Name: "Shoes", ParentID: 4},
		&Menu{ID: 7, Name: "T-Shirts", ParentID: 5},
		&Menu{ID: 8, Name: "Jeans", ParentID: 5},
		&Menu{ID: 9, Name: "Sport Shoes", ParentID: 6},
		&Menu{ID: 10, Name: "Casual Shoes", ParentID: 6},
	}

	//调用GenerateMenuList递归生成菜单列表
	menuList := GenerateMenuList(menuItems, 0)

	//打印菜单列表
	for _, menu := range menuList {
		fmt.Println(menu.Name)
		if menu.Children != nil {
			for _, child := range menu.Children {
				fmt.Printf("\t%s\n", child.Name)
				if child.Children != nil {
					for _, grandchild := range child.Children {
						fmt.Printf("\t\t%s\n", grandchild.Name)
					}
				}
			}
		}
	}
}

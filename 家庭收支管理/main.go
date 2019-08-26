package main

import (
	"fmt"
	"strconv"
)

//第一阶段 首先完成菜单显示 并能根据用户输入内容给出相关提示 输入退出时退出系统
//第二阶段 实现登记收入 并标注每笔收入备注

func main() {
	//判断是否退出
	loop := true
	//记录用户输入的字母
	key := ""
	//总金额
	money := 0
	//收入备注
	incomeNote := ""
	for {
		fmt.Println("          ————————————————")
		fmt.Println("           家庭收支管理系统")
		fmt.Println("            1.收支明细")
		fmt.Println("            2.登记收入")
		fmt.Println("            3.登记支出")
		fmt.Println("            4.查看金额")
		fmt.Println("            5.退出")
		fmt.Println("请输入选择：")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("收支明细")
		case "2":
			fmt.Println("登记收入")
			for {
				fmt.Println("请输入本次收入金额：")
				jine := ""
				fmt.Scanln(&jine)
				income, err := strconv.Atoi(jine)
				if err != nil {
					fmt.Println("输入金额必须为数字")

				} else {
					money += income
					break
				}
			}
			fmt.Println("请输入本次收入备注：")
			fmt.Scanln(&key)
			incomeNote += key + "\t"
			fmt.Println("收入登记成功")
		case "3":
			for {
				fmt.Println("请输入本次支出金额：")
				jine := ""
				fmt.Scanln(&jine)
				outcome, err := strconv.Atoi(jine)
				if err != nil {
					fmt.Println("输入金额必须为数字")

				} else {
					if money <= outcome {
						fmt.Println("余额不足")
					} else {
						money -= outcome
						break
					}
				}
			}
			fmt.Println("请输入本次支出备注：")
			fmt.Scanln(&key)
			incomeNote += key + "\t"
			fmt.Println("支出登记成功")
		case "4":
			fmt.Printf("账号余额为%v:", money)
			fmt.Println()
		case "5":
			fmt.Println("您确认退出？输入y退出其他无效")
			fmt.Scanln(&key)
			if key == "y" {
				loop = false
			}
		default:
			fmt.Println("输入有误")
		}

		if !loop {
			break
		}
	}
	fmt.Println("您已退出")
}

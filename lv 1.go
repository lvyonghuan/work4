package main

//前排：只实现了ex里的部分功能，保留了一些注释掉的旧版代码

import "fmt"

func main() {
	skillmap := map[int]string{
		1: "尝尝我的厉害吧！",
		2: "让我们打开天窗说亮话吧！",
		3: "你的恶名从爱尔兰到契丹无人不知无人不晓！",
	}
	//fmt.Println("输入你希望选择的技能\n1,龙卷风摧毁停车场；2,乌鸦坐飞机；3，自定义技能")
	var skill string
	var sle int
	fmt.Println("输入你选择的释放模板。1," + skillmap[1] + ";2," + skillmap[2] + ";3," + skillmap[3] + "。")
	_, _ = fmt.Scanln(&sle)
	fmt.Println("请输入技能名。")
	_, _ = fmt.Scanln(&skill)
	switch sle {
	case 1:
		ReleaseSkill(skill, skillmap, 1, releaseSkillFunc)
	case 2:
		ReleaseSkill(skill, skillmap, 2, releaseSkillFunc)
	case 3:
		ReleaseSkill(skill, skillmap, 3, releaseSkillFunc)
	}
}
func releaseSkillFunc(name string, skill string) {
	fmt.Println(name + skill + "!")
}
func ReleaseSkill(skillNames string, skillmap map[int]string, num int, releaseSkillFunc func(string, string)) {
	releaseSkillFunc(skillmap[num], skillNames)
}

//switch sle {
//case 1:
//	ReleaseSkill("龙卷风摧毁停车场", func(name string) {
//		fmt.Println("尝尝我的厉害吧！" + name + "！")
//	})
//case 2:
//	ReleaseSkill("乌鸦坐飞机", func(name string) {
//		fmt.Println("尝尝我的厉害吧！" + name + "！")
//	})
//case 3:
//	fmt.Println("请输入技能名")
//	var temp string
//	fmt.Scanln(&temp)
//	ReleaseSkill(temp, func(name string) {
//		fmt.Println("尝尝我的厉害吧！" + name + "！")
//	})
//}

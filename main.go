package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"

	"github.com/kfpwd/db"
)

const (
	defaultLength  = 12
	lowerCharSet   = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberSet      = "0123456789"
	specialCharSet = "!@#$%^&*=+"
)

// 确保在指定位置添加一个指定类型的字符
func ensureCharType(password []byte, charSet string, position int) []byte {
	if len(password) == 0 || position >= len(password) {
		return password
	}

	// 随机选择一个字符
	charIdx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))

	// 在指定位置放置字符
	password[position] = charSet[charIdx.Int64()]

	return password
}

func generatePassword(length int) string {
	if length < 4 {
		length = 4 // 最小长度为4
	}

	// 包含所有字符集（除了特殊字符）
	charSet := lowerCharSet + upperCharSet + numberSet

	password := make([]byte, length)
	charSetLength := big.NewInt(int64(len(charSet)))

	// 随机生成密码（除了最后一位）
	for i := 0; i < length-1; i++ {
		randomIndex, _ := rand.Int(rand.Reader, charSetLength)
		password[i] = charSet[randomIndex.Int64()]
	}

	// 确保密码包含各种类型的字符
	password = ensureCharType(password, lowerCharSet, 0)
	password = ensureCharType(password, upperCharSet, 1)
	password = ensureCharType(password, numberSet, 2)

	// 确保最后一个字符是特殊字符
	specialCharIdx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(specialCharSet))))
	password[length-1] = specialCharSet[specialCharIdx.Int64()]

	return string(password)
}

// 显示帮助信息
func showHelp() {
	fmt.Println("密码管理工具")
	fmt.Println("用法: kfpwd <command> [options]")
	fmt.Println("\n命令:")
	fmt.Println("  create <length>  生成指定长度的随机密码")
	fmt.Println("  list            显示所有保存的密码")
	fmt.Println("  save <name> <password> [url] 保存密码到数据库")
	fmt.Println("  delete <id>     删除指定ID的密码记录")
	fmt.Println("\n选项:")
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	// 初始化数据库
	dbSqlite, err := db.InitDB()
	if err != nil {
		fmt.Printf("初始化数据库失败: %v\n", err)
		os.Exit(1)
	}
	defer dbSqlite.Close()

	// 解析命令行参数
	if len(os.Args) < 2 {
		showHelp()
	}

	// 根据命令执行相应操作
	switch os.Args[1] {
	case "create":

		length := defaultLength
		if len(os.Args) > 2 {
			fmt.Sscanf(os.Args[2], "%d", &length)
		}
		if length < 6 {
			fmt.Println("错误: 密码长度必须大于6")
			os.Exit(1)
		}
		password := generatePassword(length)
		fmt.Println("生成的密码:", password)

	case "list":
		passwords, err := db.ListPasswords(dbSqlite)
		if err != nil {
			fmt.Printf("获取密码列表失败: %v\n", err)
			os.Exit(1)
		}
		if len(passwords) == 0 {
			fmt.Println("没有保存的密码")
			return
		}
		fmt.Printf("ID\t创建时间\t\t名称\t\t密码\t\t\tURL\n")
		for _, p := range passwords {
			url := p.URL
			if url == "" {
				url = "--"
			}
			fmt.Printf("%d\t%s\t%s\t\t%s\t\t%s\n",
				p.ID,
				p.CreatedAt.Format("2006-01-02 15:04"),
				p.Name,
				p.Value,
				url)
		}

	case "save":
		if len(os.Args) < 4 {
			fmt.Println("错误: 请提供密码名称和密码值")
			os.Exit(1)
		}
		name := os.Args[2]
		password := os.Args[3]
		url := ""
		if len(os.Args) > 4 {
			url = os.Args[4]
		}
		err := db.SavePassword(dbSqlite, name, password, url)
		if err != nil {
			fmt.Printf("保存密码失败: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("密码保存成功")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("错误: 请提供要删除的密码ID")
			os.Exit(1)
		}
		var id int64
		fmt.Sscanf(os.Args[2], "%d", &id)
		err := db.DeletePassword(dbSqlite, id)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Println("密码删除成功")

	default:
		showHelp()
	}
}

package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"
)

const (
	dbName  = "passwords.db"
	dbTable = "passwords"
)

type Password struct {
	ID        int64
	Name      string
	Value     string
	URL       string
	CreatedAt time.Time
}

// InitDB 初始化数据库连接和表结构
func InitDB() (*sql.DB, error) {
	// 获取当前执行文件所在目录
	exePath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("获取执行路径失败: %v", err)
	}
	exeDir := filepath.Dir(exePath)
	dbPath := filepath.Join(exeDir, dbName)

	// 连接数据库
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %v", err)
	}

	// 创建表
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS passwords (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			value TEXT NOT NULL,
			url TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("创建表失败: %v", err)
	}

	return db, nil
}

// SavePassword 保存密码到数据库
func SavePassword(db *sql.DB, name, password, url string) error {
	_, err := db.Exec("INSERT INTO passwords (name, value, url) VALUES (?, ?, ?)", name, password, url)
	if err != nil {
		return fmt.Errorf("保存密码失败: %v", err)
	}
	return nil
}

// ListPasswords 列出所有保存的密码
func ListPasswords(db *sql.DB) ([]Password, error) {
	rows, err := db.Query("SELECT id, name, value, url, created_at FROM passwords ORDER BY created_at DESC")
	if err != nil {
		return nil, fmt.Errorf("查询密码失败: %v", err)
	}
	defer rows.Close()

	var passwords []Password
	for rows.Next() {
		var p Password
		err := rows.Scan(&p.ID, &p.Name, &p.Value, &p.URL, &p.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("读取密码记录失败: %v", err)
		}
		passwords = append(passwords, p)
	}

	return passwords, nil
}

// DeletePassword 根据ID删除密码记录
func DeletePassword(db *sql.DB, id int64) error {
	result, err := db.Exec("DELETE FROM passwords WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("删除密码失败: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("获取影响行数失败: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("未找到ID为%d的密码记录", id)
	}

	return nil
}

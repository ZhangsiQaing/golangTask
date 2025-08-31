package sql

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Account struct {
	ID      uint    `gorm:"primaryKey;autoIncrement;column:id"`
	Balance float64 `gorm:"not null;default:0;column:balance"`
}

type Transaction struct {
	ID            uint    `gorm:"primaryKey;autoIncrement;column:id"`
	FromAccountID uint    `gorm:"not null;column:from_account_id"`
	ToAccountID   uint    `gorm:"not null;column:to_account_id"`
	Amount        float64 `gorm:"not null;column:amount"`
}

func Transfer(db *gorm.DB, fromID, toID uint, amount float64) error {
	tx := db.Begin() // 开启事务
	if tx.Error != nil {
		return tx.Error
	}

	var from Account
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}). // 行级锁
									Where("id = ?", fromID).First(&from).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("查询转出账户失败: %w", err)
	}

	// 检查余额
	if from.Balance < amount {
		tx.Rollback()
		return fmt.Errorf("账户余额不足")
	}

	// 扣钱
	if err := tx.Model(&Account{}).
		Where("id = ?", fromID).
		Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("扣钱失败: %w", err)
	}

	// 加钱
	if err := tx.Model(&Account{}).
		Where("id = ?", toID).
		Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("加钱失败: %w", err)
	}

	// 记录交易
	transaction := Transaction{
		FromAccountID: fromID,
		ToAccountID:   toID,
		Amount:        amount,
	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("写入交易失败: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("提交事务失败: %w", err)
	}

	return nil
}

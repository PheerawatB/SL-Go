package repositories

import "gorm.io/gorm"

type BankAccount struct {
	ID            string
	AccountHolder string
	AccountType   int
	Balance       float64
}

type AccountRepository interface {
	Save(BankAccount BankAccount) error
	Delete(ID string) error
	FindByID(ID string) (BankAccount, error)
	FindAll() ([]BankAccount, error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	db.Table("BankAccount").AutoMigrate(&BankAccount{})
	return accountRepository{db: db}
}

func (r accountRepository) Save(bankAccount BankAccount) error {
	return r.db.Table("BankAccount").Create(&bankAccount).Error
}

func (r accountRepository) Delete(ID string) error {
	return r.db.Table("BankAccount").Where("id=?", ID).Delete(&BankAccount{}).Error
}

func (r accountRepository) FindByID(ID string) (BankAccount, error) {
	var account BankAccount
	err := r.db.Table("BankAccount").Where("id=?", ID).First(&account).Error
	return account, err
}

func (r accountRepository) FindAll() (accounts []BankAccount, err error) {
	if err := r.db.Table("BankAccount").Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

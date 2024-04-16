package entities

type MstUser struct {
	Id        string  `gorm:"column:id;primaryKey" json:"id"`
	Username  string  `gorm:"column:username" json:"username"`
	Email     string  `gorm:"column:email" json:"email"`
	Password  string  `gorm:"column:password" json:"password"`
	Phone     string  `gorm:"column:phone" json:"phone"`
	CreatedAt string  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *string `gorm:"column:deleted_at" json:"deleted_at"`
}


func (MstUser) TableName() string {
	return "mst_user"
}

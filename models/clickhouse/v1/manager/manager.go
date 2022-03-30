package manager

type Manager struct {
	ID      uint8  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	// Phone     string          `json:"phone"`
	// Pwd       string          `json:"-"`
	// CreatedAt models.GormTime `json:"created_at"`
	// UpdatedAt models.GormTime `json:"updated_at"`
}

func (Manager) TableName() string {
	return "manager"
}

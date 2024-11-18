package models

func (Table) TableName() string { return "tables" }

func (MenuGet) TableName() string { return "menus" }

func (OrderCreate) TableName() string { return "orders" }

func (TableCreate) TableName() string { return Table{}.TableName() }

func (TableGet) TableName() string { return Table{}.TableName() }

func (OrderGet) TableName() string { return OrderCreate{}.TableName() }

/*Nếu bạn muốn sử dụng một tên bảng cụ thể, 
bạn có thể định nghĩa tên bảng trong struct 
của bạn bằng cách thêm hàm TableName() 
trong struct model, như sau:

type Table struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string
	Price float64
}

func (Table) TableName() string {
	return "custom_table_name"
}
*/
package v2

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

var (
	once sync.Once
	db   *sql.DB
	err  error
)

func SQLiteConn() (*sql.DB, error) {
	once.Do(func() {
		db, err = sql.Open("sqlite3", "./database.db")
		if err != nil {
			return
		}
		_, err = db.Exec(`
		create table if not exists products(
			id text primary key,
			name text not null,
			category text not null,
			price real,
			is_active integer not null default 0
		);
		delete from products;
		insert into products(id, name, category, price, is_active) values
			('1', 'Nike Running Shoes', 'Footwear', 120.00, 0),
			('2', 'Adidas Running Shoes', 'Footwear', 160.00, 1),
			('3', 'Adidas Shoes', 'Footwear',100.00, 1),
			('4', 'Adidas Sneakers', 'Footwear',110.00, 1),
			('5', 'Leather Boots', 'Footwear',150.00, 1),
			('6', 'Slim Fit T-Shirt', 'Clothing',25.00, 1),
			('7', 'Denim Jeans', 'Clothing',60.00, 1),
			('8', 'Hoodie Jacket', 'Clothing',45.00, 1),
			('9', 'Bluetooth Headphones', 'Electronics',89.99, 1),
			('10', 'Smartphone Case', 'Electronics',15.00, 1),
			('11', '4K Monitor', 'Electronics',320.00, 1),
			('12', 'Wireless Mouse', 'Electronics',29.99, 1),
			('13', 'Mechanical Keyboard', 'Electronics',99.99, 1),
			('14', 'Gaming Chair', 'Furniture',220.00, 1),
			('15', 'Office Desk', 'Furniture',180.00, 1),
			('16', 'Bookshelf', 'Furniture',85.00, 1),
			('17', 'LED Lamp', 'Home',35.00, 1),
			('18', 'Coffee Maker', 'Home',70.00, 1),
			('19', 'Blender', 'Home',55.00, 1),
			('20', 'Water Bottle', 'Accessories',12.00, 1),
			('21', 'Sunglasses', 'Accessories',40.00, 1),
			('22', 'Backpack', 'Accessories',50.00, 1),
			('23', 'Perfume', 'Beauty',65.00, 1),
			('24', 'Lipstick', 'Beauty',20.00, 1),
			('25', 'Foundation', 'Beauty',30.00, 1),
			('26', 'Yoga Mat', 'Sports',25.00, 1),
			('27', 'Dumbbell Set', 'Sports',90.00, 1),
			('28', 'Tennis Racket', 'Sports',75.00, 1),
			('29', 'Cookbook', 'Books',18.00, 1),
			('30', 'Science Fiction Novel', 'Books',22.00, 1),
			('31', 'Notebook', 'Stationery',5.00, 1),
			('32', 'Pen Pack', 'Stationery',3.50, 1)
	`)
	})
	return db, err
}

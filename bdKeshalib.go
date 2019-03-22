package kesha

import "database/sql"

var (
	db *sql.DB
	dbName string

)

//SetDB create reference to your database
func SetDB(userDb *sql.DB,name string)  {
	db=userDb
	dbName=name

}

//GetColumn return list of values by name of column
func GetColumn(arg string) []string {
	var s string
	mas:=make([]string,0)
	rows,_:= db.Query("SELECT "+arg+" FROM "+dbName)
	for rows.Next() {
		rows.Scan(&s)
		mas=append(mas,s)
	}
	return mas
}

/*
func First(col,val string)string  {
	rows,_:= db.Query("SELECT "+col+" FROM "+dbName+"where "+)
}
*/

package models

type Test struct {
	Name  string
	Email string
}

// func AllTests() ([]*Test, error) {
// 	rows, err := db.Query("SELECT name, email FROM users")
// 	if err != nil {

// 		return nil, err
// 	}
// 	defer rows.Close()
// 	bks := make([]*Test, 0)
// 	for rows.Next() {
// 		bk := new(Test)
// 		err := rows.Scan(&bk.Name, &bk.Email)
// 		if err != nil {

// 			return nil, err
// 		}
// 		bks = append(bks, bk)
// 	}
// 	if err = rows.Err(); err != nil {

// 		return nil, err
// 	}

// 	return bks, nil
// }

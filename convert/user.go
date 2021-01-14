package convert

import (
	"fmt"
	"github.com/abusizhishen/fileConvert/src"
	"regexp"
)

// row[0] name, row[1] id
type User struct {
	Phone string `json:"id"`
	Name  string `json:"name"`
}

type Users []User

func (us *Users) Rows() (rows []src.Row) {
	rows = append(rows, src.Row{"姓名", "手机号"})
	for _, u := range *us {
		rows = append(rows, src.Row{u.Name, u.Phone})
	}

	return rows
}

var phoneReg *regexp.Regexp

func init() {
	var err error
	phoneReg, err = regexp.Compile(`^1[0-9]{10}$`)
	if err != nil {
		panic(err)
	}
}

func ToUser(rows []src.Row) (users []User, err error) {
	for i, row := range rows {
		if len(row) < 2 {
			err = fmt.Errorf("i%v列不匹配", i+1)
			return nil, err
		}
		phone := row[1]
		if !phoneReg.MatchString(phone) {
			err = fmt.Errorf("第%d行手机号格式不正确, 手机号: %v", i+1, phone)
			return nil, err
		}
		users = append(users, User{
			Phone: phone,
			Name:  row[0],
		})
	}

	return users, err
}

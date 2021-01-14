package convert

import (
	"fmt"
	"github.com/abusizhishen/fileConvert/src"
	"strconv"
)

// row[0] name, row[1] id
type City struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func ToCity(rows []src.Row) (citys []City, err error) {
	for i, row := range rows {
		if len(row) < 2 {
			err = fmt.Errorf("i%v列不匹配 row:%+v", i+1, row)
			return nil, err
		}
		id, err := strconv.Atoi(row[0])
		if err != nil {
			err = fmt.Errorf("第%d行城市id格式不正确, id:%v", i, row[0])
			return nil, err
		}
		citys = append(citys, City{
			Id:   id,
			Name: row[1],
		})
	}

	return citys, err
}

type Citys []City

func (cs *Citys) Rows() (rows []src.Row) {
	rows = append(rows, src.Row{"城市id", "城市名"})
	for _, c := range *cs {
		rows = append(rows, src.Row{c.Name, fmt.Sprint("%d", c.Id)})
	}

	return rows
}

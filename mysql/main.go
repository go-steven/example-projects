package main

import (
	"fmt"
	"github.com/bububa/goconfig/config"
	"github.com/bububa/mymysql/autorc"
	_ "github.com/bububa/mymysql/thrsafe"
	"github.com/go-steven/cerr"
	"github.com/go-steven/json"
	log "github.com/go-steven/logger"
	"strings"
)

const (
	_CONFIG_FILE = "/var/code/go/config.cfg"
)

var (
	logger = log.NewLogger("")
)

/*
CREATE TABLE `steven_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
*/
type User struct {
	Id   uint64 `json:"id,omitempty" codec:"id,omitempty"`
	Name string `json:"name,omitempty" codec:"name,omitempty"`
}

func main() {
	cfg, _ := config.ReadDefault(_CONFIG_FILE)

	host, _ := cfg.String("masterdb", "host")
	user, _ := cfg.String("masterdb", "user")
	passwd, _ := cfg.String("masterdb", "passwd")
	dbname, _ := cfg.String("masterdb", "dbname")

	mdb := autorc.New("tcp", "", host, user, passwd, dbname)
	mdb.Register("set names utf8")

	// batch insert
	users := []*User{
		&User{
			Id:   1,
			Name: "test",
		},
		&User{
			Id:   2,
			Name: "test2",
		},
	}
	logger.Infof("users: %v", json.Json(users))
	if err := batch_insert(mdb, users); err != nil {
		logger.Error(err.Error())
		return
	}

	// fetch all rows
	ret_users, err := fetch_all(mdb)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Infof("ret_users: %v", json.Json(ret_users))

	// fetch row by id
	ret_user, err := fetch_by_id(mdb, 1)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Infof("ret_user: %v", json.Json(ret_user))

	_, res, err := mdb.Query(`UPDATE steven_users SET name = '%s' WHERE id = %d`, mdb.Escape("test3"), 1)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Infof("update affected rows: %v", res.AffectedRows())

	_, res, err = mdb.Query(`DELETE FROM steven_users WHERE id = %d`, 2)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Infof("delete affected rows: %v", res.AffectedRows())
}

func batch_insert(db *autorc.Conn, users []*User) error {
	if len(users) == 0 {
		return nil
	}
	query := `INSERT INTO steven_users (id, name) VALUES %s ON DUPLICATE KEY UPDATE name=VALUES(name)`
	var val []string
	for _, v := range users {
		val = append(val, fmt.Sprintf("(%d, '%s')", v.Id, db.Escape(v.Name)))
		_, _, err := db.Query(query, strings.Join(val, ","))
		if err != nil {
			return cerr.New(err.Error())
		}
		val = []string{}
	}
	if len(val) > 0 {
		_, _, err := db.Query(query, strings.Join(val, ","))
		if err != nil {
			return cerr.New(err.Error())
		}
	}
	return nil
}

func fetch_all(db *autorc.Conn) ([]*User, error) {
	query := `SELECT id, name FROM steven_users ORDER BY id DESC`
	rows, res, err := db.Query(query)
	if err != nil {
		return nil, cerr.New(err.Error())
	}
	users := []*User{}
	for _, row := range rows {
		users = append(users, &User{
			Id:   row.Uint64(res.Map("id")),
			Name: row.Str(res.Map("name")),
		})
	}

	return users, nil
}

func fetch_by_id(db *autorc.Conn, id uint64) (*User, error) {
	query := `SELECT id, name FROM steven_users WHERE id = %d LIMIT 1`
	rows, res, err := db.Query(query, id)
	if err != nil {
		return nil, cerr.New(err.Error())
	}
	if len(rows) == 0 {
		return nil, nil
	}

	return &User{
		Id:   rows[0].Uint64(res.Map("id")),
		Name: rows[0].Str(res.Map("name")),
	}, nil
}

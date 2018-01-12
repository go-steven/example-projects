package main

import (
	"github.com/bububa/mymysql/autorc"
	_ "github.com/bububa/mymysql/thrsafe"
	"sync"
	"time"
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

type UsersService struct {
	db *autorc.Conn

	data map[uint64]*User
	m    *sync.RWMutex

	exitChan chan struct{}
}

func NewUsersService(db *autorc.Conn) *UsersService {
	return &UsersService{
		db:   db,
		data: make(map[uint64]*User),
		m:    new(sync.RWMutex),

		exitChan: make(chan struct{}),
	}
}

func (s *UsersService) Start() {
	s.LoadData()

	queueTicker := time.NewTicker(1 * time.Minute) // 定时
	for {
		select {
		case <-queueTicker.C:
			s.LoadData()
		case <-s.exitChan:
			queueTicker.Stop()
			return
		}
	}
}

func (s *UsersService) Stop() {
	select {
	case <-s.exitChan:
	default:
		close(s.exitChan)
	}
}

func (s *UsersService) Get(id uint64) (ret *User) {
	s.m.RLock()
	if v, ok := s.data[id]; ok {
		ret = v
	}
	s.m.RUnlock()
	return
}

func (s *UsersService) LoadData() error {
	logger.Infof("UsersService: loading data.")
	task_start_time := time.Now()
	defer func() {
		logger.Infof("UsersService: total duration for load data: %v", time.Since(task_start_time))
	}()

	data := make(map[uint64]*User)
	query := `SELECT id, name FROM steven_users WHERE id > %d ORDER BY id LIMIT 1000`
	var startId uint64
	for {
		rows, res, err := s.db.Query(query, startId)
		if err != nil {
			logger.Error(err)
			return err
		}
		logger.Infof("rows:%d", len(rows))
		var id uint64
		for _, row := range rows {
			id = row.Uint64(res.Map("id"))
			data[id] = &User{
				Id:   id,
				Name: row.Str(res.Map("name")),
			}
		}

		if len(rows) < 1000 {
			break
		}

		startId = id
	}

	s.m.Lock()
	s.data = data
	s.m.Unlock()
	return nil
}

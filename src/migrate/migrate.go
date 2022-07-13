package migrate

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/mensylisir/kmpp-middleware/src/constant"
	"github.com/mensylisir/kmpp-middleware/src/db"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/model"
	"github.com/mensylisir/kmpp-middleware/src/util/encrypt"
	"github.com/mensylisir/kmpp-middleware/src/util/file"
)

const (
	phaseName = "migrate"
)

const (
	releaseMigrationDir = "/usr/local/lib/middleware/migration"
	localMigrationDir   = "./migration"
)

var migrationDirs = []string{
	localMigrationDir,
	releaseMigrationDir,
}

type InitMigrateDBPhase struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

func (i *InitMigrateDBPhase) Init() error {
	aesPasswd, er1 := encrypt.StringEncrypt(i.Password)
	if er1 != nil {
		return er1
	}
	p, err := encrypt.StringDecrypt(aesPasswd)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Asia%%2FShanghai",
		i.User,
		p,
		i.Host,
		i.Port,
		i.Name)
	var path string
	for _, d := range migrationDirs {
		if file.Exists(d) {
			path = d
		}
	}
	if path == "" {
		return fmt.Errorf("can not find migration in [%s,%s]", localMigrationDir, releaseMigrationDir)
	}
	filePath := fmt.Sprintf("file://%s", path)
	m, err := migrate.New(
		filePath, url)
	if err != nil {
		return err
	}
	// 初始化默认用户
	v, _, _ := m.Version()
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logger.Log.Info("no databases change,skip migrate")
			return nil
		}
		return err
	}
	dp, err := encrypt.StringEncrypt(constant.DefaultPassword)
	if err != nil {
		return fmt.Errorf("can not init default user")
	}
	if !(v > 0) {
		if err := db.DB.Model(&model.User{}).Where("name = ?", "admin").Updates(map[string]interface{}{"Password": dp}).Error; err != nil {
			return fmt.Errorf("can not update default user")
		}
	}
	return nil
}

func (i *InitMigrateDBPhase) PhaseName() string {
	return phaseName
}

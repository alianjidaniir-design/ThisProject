package task

import (
	"log"
	"sync"

	"ThisProject/models/repositories"
	taskDataSources "ThisProject/models/task/dataSources"
	memoryDataSource "ThisProject/models/task/dataSources/memoryDS"
	mysqlDataSource "ThisProject/models/task/dataSources/mysqlDS"
)

type Repository struct {
	cacheDS taskDataSources.TaskCacheDS
	dbDS    taskDataSources.TaskDBDS
	initErr error
}

var (
	once    sync.Once
	repoIns *Repository
)

func GetRepo() *Repository {
	once.Do(func() {
		repoIns = &Repository{
			cacheDS: memoryDataSource.NewTaskCacheDS(),
		}
		repoIns.initializeDataSources()
	})

	return repoIns
}

func init() {
	repositories.TaskRepo = GetRepo()
}

func (repo *Repository) initializeDataSources() {
	mysqlDS, enabled, err := mysqlDataSource.NewTaskDBDSFromEnv()
	if err != nil {
		repo.initErr = err
		return
	}

	if enabled {
		repo.dbDS = mysqlDS
		log.Printf("[task-repository] mysql datasource enabled table=%s", mysqlDS.TableName())
		return
	}

	repo.dbDS = memoryDataSource.NewTaskDBDS(100)
	log.Println("[task-repository] MYSQL_DSN is empty, using memory datasource")
}

func (repo *Repository) db() taskDataSources.TaskDBDS {
	return repo.dbDS
}

func (repo *Repository) cache() taskDataSources.TaskCacheDS {
	return repo.cacheDS
}

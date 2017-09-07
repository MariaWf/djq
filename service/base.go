package service

import (
	"mimi/djq/dao"
	"mimi/djq/db/mysql"
	"github.com/pkg/errors"
	"database/sql"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"mimi/djq/util"
	"log"
)

var ErrIdEmpty = errors.New("ID为空")

var ErrObjectNotFound = errors.New("找不到对象")

var ErrObjectEmpty = errors.New("对象为空")

var ErrUnknown = errors.New("操作失败，请稍后重试")

func checkErr(err error) error {
	if err == nil {
		return nil
	}
	if err == dao.ErrIdEmpty {
		return ErrIdEmpty
	}
	if err == dao.ErrObjectEmpty {
		return ErrObjectEmpty
	}
	if err == dao.ErrObjectNotFound {
		return ErrObjectNotFound
	}
	return errors.Wrap(err,"操作失败，请稍后重试")
	//return ErrUnknown
}

type BaseServiceInterface interface {
	GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface
}

func Page(service BaseServiceInterface, argObj arg.BaseArgInterface) (*util.PageVO, error) {
	conn, err := mysql.Get()
	if err != nil {
		return nil, errors.Wrap(err, "db")
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn)
	list, err := dao.Find(daoObj, argObj)
	if err != nil {
		rollback = true
		return nil, errors.Wrap(err, "page:find")
	}
	total, err := dao.Count(daoObj, argObj)
	if err != nil {
		rollback = true
		return nil, errors.Wrap(err, "page:count")
	}
	return util.BuildPageVO(argObj.GetTargetPage(), argObj.GetPageSize(), total, list), nil
}

func ResultList(service BaseServiceInterface, argObj arg.BaseArgInterface) (*util.ResultVO) {
	return Result(Page(service, argObj))
}

func ResultGet(service BaseServiceInterface, id string) (*util.ResultVO) {
	return Result(Get(service, id))
}

func ResultGount(service BaseServiceInterface, argObj arg.BaseArgInterface) (*util.ResultVO) {
	return Result(Count(service, argObj))
}

func ResultAdd(service BaseServiceInterface, aobj model.BaseModelInterface) (*util.ResultVO) {
	return Result(Add(service, aobj))
}

func ResultUpdate(service BaseServiceInterface, obj model.BaseModelInterface, args ... string) (*util.ResultVO) {
	return Result(Update(service, obj, args...))
}

func ResultBatchUpdate(service BaseServiceInterface, argObj arg.BaseArgInterface) (*util.ResultVO) {
	return Result(BatchUpdate(service, argObj))
}

func ResultDelete(service BaseServiceInterface, id string) (*util.ResultVO) {
	return Result(Delete(service, id))
}

func ResultBatchDelete(service BaseServiceInterface, argObj arg.BaseArgInterface) (*util.ResultVO) {
	return Result(BatchDelete(service, argObj))
}

func Result(obj interface{}, err error) *util.ResultVO {
	if err != nil {
		log.Println(err)
		return util.BuildFailResult(checkErr(err).Error())
	} else {
		return util.BuildSuccessResult(obj)
	}
}

func Find(service BaseServiceInterface, argObj arg.BaseArgInterface) ([]interface{}, error) {
	conn, err := mysql.Get()
	if err != nil {
		return nil, checkErr(err)
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn)
	list, err := dao.Find(daoObj, argObj)
	if err != nil {
		rollback = true
	}
	return list, checkErr(err)
}

func Count(service BaseServiceInterface, argObj arg.BaseArgInterface) (int, error) {
	conn, err := mysql.Get()
	if err != nil {
		return 0, checkErr(err)
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn)
	count, err := dao.Count(daoObj, argObj)
	if err != nil {
		rollback = true
	}
	return count, checkErr(err)
}

func Get(service BaseServiceInterface, id string) (interface{}, error) {
	conn, err := mysql.Get()
	if err != nil {
		return nil, checkErr(err)
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn)
	obj, err := dao.Get(daoObj, id)
	if err != nil {
		rollback = true
	}
	return obj, checkErr(err)
}

func Add(service BaseServiceInterface, obj model.BaseModelInterface) (interface{}, error) {
	conn, err := mysql.Get()
	if err != nil {
		return nil, checkErr(err)
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn)
	newObj, err := dao.Add(daoObj, obj)
	if err != nil {
		rollback = true
	}
	return newObj, checkErr(err)
}

func Update(service BaseServiceInterface, obj model.BaseModelInterface, args ... string) (interface{}, error) {
	conn, err := mysql.Get()
	if err != nil {
		return nil, checkErr(err)
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn)
	newObj, err := dao.Update(daoObj, obj, args...)
	if err != nil {
		rollback = true
	}
	return newObj, checkErr(err)
}

func BatchUpdate(service BaseServiceInterface, argObj arg.BaseArgInterface) (int64, error) {
	conn, err := mysql.Get()
	if err != nil {
		return 0, checkErr(err)
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn)
	count, err := dao.BatchUpdate(daoObj, argObj)
	if err != nil {
		rollback = true
	}
	return count, checkErr(err)
}

func Delete(service BaseServiceInterface, id string) (int64, error) {
	conn, err := mysql.Get()
	if err != nil {
		return 0, checkErr(err)
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn)
	argObj := daoObj.GetArgInstance()
	argObj.SetIdEqual(id)
	count, err := dao.LogicalDelete(daoObj, argObj)
	if err != nil {
		rollback = true
	}
	return count, checkErr(err)
}

func BatchDelete(service BaseServiceInterface, argObj arg.BaseArgInterface) (int64, error) {
	conn, err := mysql.Get()
	if err != nil {
		return 0, checkErr(err)
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn)
	count, err := dao.LogicalDelete(daoObj, argObj)
	if err != nil {
		rollback = true
	}
	return count, checkErr(err)
}

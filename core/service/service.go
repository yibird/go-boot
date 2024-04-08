package service

import (
	"github.com/gin-gonic/gin"
)

type BaseService[M any, D any] struct{}

func (s *BaseService[M, D]) Save(c *gin.Context, m M, d D) error {
	//if err := convert.Convert(c, &d, &m); err != nil {
	//	return err
	//}
	//result := gormDB.Create(&m)
	//if result.RowsAffected > 0 {
	//	return nil
	//}
	//return result.Error
	return nil
}

func (s *BaseService[M, T]) DelByIds(model *M, ids []int) error {
	//if ids == nil || len(ids) == 0 {
	//	return nil
	//}
	//result := gormDB.Delete(model, ids)
	//if result.RowsAffected > 0 {
	//	return nil
	//}
	//return result.Error
	return nil
}

func (s *BaseService[M, D]) Update(c *gin.Context, m M, d D) error {
	//if err := convert.Convert(c, &d, &m); err != nil {
	//	return err
	//}
	//fmt.Println("m:", m, d)
	//result := gormDB.Model(&m).Omit("id").Updates(m)
	//if result.RowsAffected > 0 {
	//	return nil
	//}
	//return result.Error
	return nil
}

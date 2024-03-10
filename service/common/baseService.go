package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-boot/global"
	modelCommon "go-boot/model/common"
)

type BaseService[M any, D any] struct{}

func (s *BaseService[M, D]) Save(c *gin.Context, m M, d D) error {
	if err := modelCommon.Convert(c, &d, &m); err != nil {
		return err
	}
	result := global.DB.Create(&m)
	if result.RowsAffected > 0 {
		return nil
	}
	return result.Error
}

func (s *BaseService[M, T]) DelByIds(model *M, ids []int) error {
	if ids == nil || len(ids) == 0 {
		return nil
	}
	result := global.DB.Delete(model, ids)
	if result.RowsAffected > 0 {
		return nil
	}
	return result.Error
}

func (s *BaseService[M, D]) Update(c *gin.Context, m M, d D) error {
	if err := modelCommon.Convert(c, &d, &m); err != nil {
		return err
	}
	fmt.Println("m:", m, d)
	result := global.DB.Model(&m).Omit("id").Updates(m)
	if result.RowsAffected > 0 {
		return nil
	}
	return result.Error
}

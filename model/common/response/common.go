package response

import (
	"gorm.io/gorm"
)

type PageResult[T any] struct {
	List     []T   `json:"list"`     // 列表数据
	Total    int64 `json:"total"`    // 总条数
	Current  int64 `json:"current"`  // 当前页码
	PageSize int64 `json:"pageSize"` // 每页显示条数
	Pages    int64 `json:"pages"`    // 总页码
}

// Paginate 分页器
func (page *PageResult[T]) Paginate(db *gorm.DB) (e error) {
	var model T
	db.Model(&model).Count(&page.Total)
	if page.Total == 0 {
		page.List = []T{}
		return
	}
	err := db.Model(&model).Scopes(CalcPaging(page)).Find(&page.List).Error
	return err
}

// CalcPaging 计算分页
func CalcPaging[T any](page *PageResult[T]) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page.Current < 1 {
			page.Current = 1
		}
		if page.PageSize <= 0 {
			page.PageSize = 10
		} else if page.PageSize > 500 {
			page.PageSize = 500
		}
		// 计算总页码
		page.Pages = page.Total / page.PageSize
		if page.Total%page.PageSize != 0 {
			page.Pages++
		}
		p := page.Current
		// 如果当前页码大于总页码
		if page.Current > page.Pages {
			p = page.Pages
		}
		size := page.PageSize
		offset := int((p - 1) * size)
		return db.Offset(offset).Limit(int(size))
	}
}

// SelectPage 查询分页
func SelectPage[T any](db *gorm.DB, current int64, pageSize int64) (PageResult[T], error) {
	page := PageResult[T]{Current: current, PageSize: pageSize}
	err := page.Paginate(db)
	return page, err
}

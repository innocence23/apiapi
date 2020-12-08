package dto

// Pagination 分页查询条件
type Pagination struct {
	Pagination bool `form:"-"`                                     // 是否使用分页查询
	OnlyCount  bool `form:"-"`                                     // 是否仅查询count
	Current    uint `form:"current,default=1"`                     // 当前页
	PageSize   uint `form:"pageSize,default=10" binding:"max=100"` // 页大小
}

// GetCurrent 获取当前页
func (a Pagination) GetCurrent() uint {
	return a.Current
}

// GetPageSize 获取页大小
func (a Pagination) GetPageSize() uint {
	pageSize := a.PageSize
	if a.PageSize == 0 {
		pageSize = 100
	}
	return pageSize
}

// OrderField 排序字段
type OrderField struct {
	Key       string // 字段名(字段名约束为小写蛇形)
	Direction int    // 排序方向
}

const (
	// OrderASC 升序排序
	OrderASC = 1
	// OrderDESC 降序排序
	OrderDESC = 2
)

package http

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"net/http"
	"strings"
	"wilder.cn/gogo/comm"
	"wilder.cn/gogo/database"
	"wilder.cn/gogo/log"
)

const (
	DefPageSize = 20
)
const (
	WK_IS_NOT_NULL = "wk_is_not_null"
	WK_IS_NULL     = "wk_is_null"
)

// WhereBuild sql build where
func WhereBuild(where map[string]interface{}) (whereSQL string, vals []interface{}, err error) {
	for k, v := range where {
		ks := strings.Split(k, " ")
		if len(ks) > 3 {
			return "", nil, fmt.Errorf("Error in query condition: %s. ", k)
		}

		if whereSQL != "" {
			whereSQL += " AND "
		}
		strings.Join(ks, ",")
		switch len(ks) {
		case 1:
			//fmt.Println(reflect.TypeOf(v))
			switch v := v.(type) {
			case string:
				if v == WK_IS_NOT_NULL {
					whereSQL += fmt.Sprint(k, " IS NOT NULL")
				} else if v == WK_IS_NULL {
					whereSQL += fmt.Sprint(k, " IS NULL")
				} else {
					whereSQL += fmt.Sprint(k, "=?")
					vals = append(vals, v)
				}
			default:
				whereSQL += fmt.Sprint(k, "=?")
				vals = append(vals, v)
			}
			break
		case 2:
			k = ks[0]
			switch ks[1] {
			case "=":
				whereSQL += fmt.Sprint(k, "=?")
				vals = append(vals, v)
				break
			case ">":
				whereSQL += fmt.Sprint(k, ">?")
				vals = append(vals, v)
				break
			case ">=":
				whereSQL += fmt.Sprint(k, ">=?")
				vals = append(vals, v)
				break
			case "<":
				whereSQL += fmt.Sprint(k, "<?")
				vals = append(vals, v)
				break
			case "<=":
				whereSQL += fmt.Sprint(k, "<=?")
				vals = append(vals, v)
				break
			case "!=":
				whereSQL += fmt.Sprint(k, "!=?")
				vals = append(vals, v)
				break
			case "<>":
				whereSQL += fmt.Sprint(k, "!=?")
				vals = append(vals, v)
				break
			case "in":
				whereSQL += fmt.Sprint(k, " in (?)")
				vals = append(vals, v)
				break
			case "like":
				whereSQL += fmt.Sprint(k, " like ?")
				vals = append(vals, v)
			}
			break
		case 3:
			whereSQL += k
			vals = append(vals, v)
			break
		}
	}
	return
}

type Page[T any] struct {
	PageSize   int    `json:"pageSize,omitempty" form:"pageSize"`
	PageIndex  int    `json:"pageIndex,omitempty" form:"pageIndex"`
	Sort       string `json:"sort,omitempty" form:"sort"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
	Rows       []*T   `json:"rows"`
}

func NewPage[T any]() *Page[T] {
	p := &Page[T]{
		PageSize:  DefPageSize,
		PageIndex: 1,
	}
	return p
}

func (p *Page[T]) Offset() int {
	return (p.CurrentPage() - 1) * p.Limit()
}
func (p *Page[T]) InitPage(totalRows int64) {
	p.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(p.PageSize)))
	p.TotalPages = totalPages

	log.Logger.InfoF("total rows%d", p.TotalRows)
	log.Logger.InfoF("total page%d", p.TotalPages)
}

func (p *Page[T]) Limit() int {
	if p.PageSize == 0 {
		p.PageSize = DefPageSize
	}
	return p.PageSize
}

func (p *Page[T]) CurrentPage() int {
	if p.PageIndex == 0 {
		p.PageIndex = 1
	}
	return p.PageIndex
}

func (p *Page[T]) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}

type PairKey struct {
	RequestKey string
	RawSqlKey  string
}

func Where(db *gorm.DB, where map[string]interface{}) (*gorm.DB, error) {
	if len(where) > 20 {
		return nil, errors.New("查询条件最长不能超过20个条件")
	}
	dbHolder := db
	//for k, v := range where {
	//	dbHolder = dbHolder.Where(k, v)
	//}
	whereSql, valList, err := WhereBuild(where)
	if err != nil {
		return dbHolder, err
	} else {
		dbHolder = dbHolder.Where(whereSql, valList...)
	}
	return dbHolder, nil
}

func Paginate[T any](sql *string, condition map[string]interface{}, model *T, pagination *Page[T], db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	if sql != nil {
		db = db.Raw(*sql)
	}
	for k, v := range condition {
		db = db.Where(k, v)
	}
	db.Model(model).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.PageSize)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.Offset()).Limit(pagination.Limit()).Order(pagination.GetSort())
	}
}

// RawPaginateParas 根据多个条件查询，
func RawPaginateParas[T any](req *QueryRequest, c *gin.Context, rawSql string, keys ...string) {
	var sqlNamed = make([]interface{}, len(keys))
	for i, key := range keys {
		kvalue, ok := req.Where[key]
		if !ok {
			c.JSON(http.StatusOK, FailedResponse("Not found '"+key+"' in where map", ""))
			return
		} else {
			log.Logger.InfoF("string[] para. %T", kvalue)
			switch kvalue.(type) {
			case []string:
			case []interface{}:
				log.Logger.InfoF("string[] para.")
				sqlNamed[i] = sql.Named(key, kvalue)
			case bool:
				log.Logger.InfoF("bool para.")
				sqlNamed[i] = sql.Named(key, kvalue)
			default:
				sqlNamed[i] = sql.Named(key, comm.IToString(kvalue))
			}
		}
	}
	RawPaginate[T](req, c, sqlNamed, rawSql)
}

// RawPaginate 根据原始SQL进行分页查询，paras 为sql.Named构造的数组
func RawPaginate[T any](req *QueryRequest, c *gin.Context, paras []interface{}, rawSql string) {
	var rowData []*T
	page := RequestToPage[T](*req)
	countSQL := database.ToDefCount(rawSql)
	var totalRows int64
	db := database.DBHander.Raw(countSQL, paras...).Find(&totalRows)
	log.Logger.InfoF("count sql:%s", countSQL)
	log.Logger.InfoF("count totalRows=%d", totalRows)
	if db.Error != nil {
		log.Logger.InfoF("count error=%s", db.Error.Error())
	}

	page.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(page.PageSize)))
	page.TotalPages = totalPages

	log.Logger.InfoF("limit:%d", page.Limit())
	log.Logger.InfoF("Offset:%d", page.Offset())
	rawSql += " limit @limit offset @offset"
	paras = append(paras, sql.Named("limit", page.Limit()))
	paras = append(paras, sql.Named("offset", page.Offset()))
	db2 := database.DBHander.Raw(rawSql, paras...)

	db2.Debug().Find(&rowData)
	page.Rows = rowData
	c.JSON(http.StatusOK, SuccessResponse(PageToResponse(page)))
}

//func RawPaginate2[T any](req *QueryRequest, c *gin.Context, rawSql string, keys ...PairKey) {
//	var sqlNamed = make([]interface{}, len(keys))
//	for i, key := range keys {
//		kvalue, ok := req.Where[key.RequestKey]
//		if !ok {
//			c.JSON(http.StatusOK, FailedResponse("Not found '"+key.RequestKey+"' in where map", ""))
//			return
//		} else {
//			sqlNamed[i] = sql.Named(key.RawSqlKey, dmath.IToString(kvalue))
//		}
//	}
//	RawPaginate[T](req, c, sqlNamed, rawSql)
//
//}

func RawPaginateNoParas[T any](req *QueryRequest, c *gin.Context, rawSql string) {
	var paras []interface{}
	RawPaginate[T](req, c, paras, rawSql)
}

func SelectPage[T any](condition map[string]interface{}, page *Page[T]) (*Page[T], error) {
	var rowData []*T
	db := database.DBHander.Scopes(Paginate[T](nil, condition, new(T), page, database.DBHander))
	if len(condition) > 20 {
		return nil, errors.New("查询条件最长不能超过20个条件")
	}
	for k, v := range condition {
		db = db.Where(k, v)
	}
	db.Find(&rowData)

	page.Rows = rowData

	return page, nil
}

func RawSelect[T any](sql string, req *QueryRequest, c *gin.Context) {
	page := RequestToPage[T](*req)
	var rowData []*T
	db := database.DBHander.Scopes(Paginate[T](&sql, req.Where, new(T), page, database.DBHander))
	if len(req.Where) > 20 {
		c.JSON(http.StatusOK, FailedResponse("查询条件最长不能超过20个条件", len(req.Where)))
		return
	}
	db = db.Debug().Raw(sql)
	for k, v := range req.Where {
		db = db.Where(k, v)
	}
	db.Find(&rowData)
	page.Rows = rowData

	if db.Error != nil {
		c.JSON(http.StatusOK, FailedResponse("query failed", db.Error))
	} else {
		qres := &QResponse{}
		qres.PageSize = page.PageSize
		qres.PageIndex = page.PageIndex
		qres.TotalPages = page.TotalPages
		qres.Data = page.Rows
		c.JSON(http.StatusOK, SuccessResponse(qres))
	}
}
func RawSQL[T any](sql string, req *QueryRequest, c *gin.Context, countSQL string) {
	page := RequestToPage[T](*req)
	var rowData []*T
	var totalRows int64
	where, sqlBuild := database.SQLExpression(req.Where)
	log.Logger.InfoF("where%s", where)
	sqlFull := sql + " " + where
	database.DBHander.Debug().Raw(database.ToCount(sqlFull, countSQL)).Find(&totalRows)
	log.Logger.InfoF("total %d", totalRows)
	page.InitPage(totalRows)
	sqlPage, _, _ := sqlBuild.Offset(uint(page.Offset())).Limit(uint(page.Limit())).ToSQL()
	log.Logger.InfoF("limit sql:%s", sqlPage)

	db := database.DBHander.Debug().Raw(sql + " " + database.LookWhere(sqlPage)).Find(&rowData)
	page.Rows = rowData
	if db.Error != nil {
		c.JSON(http.StatusOK, FailedResponse("query failed", db.Error))
	} else {
		qres := &QResponse{}
		qres.PageSize = page.PageSize
		qres.PageIndex = page.PageIndex
		qres.TotalPages = page.TotalPages
		qres.Data = page.Rows
		c.JSON(http.StatusOK, SuccessResponse(qres))
	}
}

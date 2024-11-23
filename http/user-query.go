package http

import (
	"wilder.cn/gogo/database"
	"wilder.cn/gogo/log"
)

// SelectPage 分页查询
//func SelectPage[T any](condition map[string]interface{}, page *comm.Page[T]) (*comm.Page[T], error) {
//	var rowData []*T
//	db := database.DBHander.Scopes(comm.Paginate[T](condition, new(T), page, database.DBHander))
//	if len(condition) > 20 {
//		return nil, errors.New("查询条件最长不能超过20个条件")
//	}
//	for k, v := range condition {
//		db = db.Where(k, v)
//	}
//	db.Find(&rowData)
//
//	page.Rows = rowData
//
//	return page, nil
//}

const (
	OPERS_SQL = "select distinct d.operatorid from s_users a left join s_groupuser b on a.id=b.userid  left join s_rolegroup c on b.groupid=c.groupid left join s_roleoperator d on d.roleid=c.roleid    where a.email=?;"
	URL_SQL   = "select s.operatorid,s.url from s_urlmappings s where s.operatorid in (select distinct d.operatorid from s_users a left join s_groupuser b on a.id=b.userid  left join s_rolegroup c on b.groupid=c.groupid left join s_roleoperator d on d.roleid=c.roleid    where a.id=?)"
)

//根据用户查询可以操作的Operatorid

func UserOperators(user string) map[int32]interface{} {
	operMap := make(map[int32]interface{})
	var operList []UserOperator
	err := database.DBHander.Raw(OPERS_SQL, user).Scan(&operList).Error
	if err != nil {
		log.Logger.ErrorF("Load user operator failed. %s", err.Error())
	} else {
		for _, oper := range operList {
			operMap[oper.OperatorID] = oper
		}
	}
	return operMap
}

//根据用户查询可以访问的URL

func UserAllUrlList(user string) map[string]interface{} {
	operMap := make(map[string]interface{})
	var urlList []UserAllowAccess
	err := database.DBHander.Raw(URL_SQL, user).Scan(&urlList).Error
	if err != nil {
		log.Logger.ErrorF("Load user operator failed. %s", err.Error())
	} else {
		log.Logger.InfoF("found url count %d", len(urlList))
		for _, urlObj := range urlList {
			operMap[urlObj.Url] = urlObj
			log.Logger.InfoF("add user url %s", urlObj.Url)
		}
	}
	return operMap
}

type UserOperator struct {
	//Owner     string    `gorm:"column:owner" json:"owner"`
	OperatorID int32 `gorm:"column:operatorid" json:"operator"`
}
type UserAllowAccess struct {
	Url string `gorm:"column:url" json:"url"`
}

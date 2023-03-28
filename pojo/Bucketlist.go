package pojo

import DB "GolangAPI/database"

type Bucketlist struct {
	ListId       int    `json:"ListId"`
	UserId       string `json:"UserId"`
	ListTitle    string `json:"ListTitle"`
	Content      string `json:"Content"`
	ListStatus   string `json:"ListStatus "`
	CreationDate string `json:"CreationDate"`
}

func FindAllListByUserId(userId string) []Bucketlist {
	var bucketList []Bucketlist
	DB.DBconnect.Raw("select * from BUCKETLISTS where USER_ID = ?", userId).Scan(&bucketList)
	return bucketList
}

func FindListByListId(listId string) Bucketlist {
	var bucketList Bucketlist
	// database.DBconnect.Find(&user)
	DB.DBconnect.Raw("select * from BUCKETLISTS where LIST_ID = ?", listId).Scan(&bucketList)
	return bucketList
}

func FindLastListByUserId(userId string) Bucketlist {
	list := Bucketlist{}
	DB.DBconnect.Where("USER_ID = ?", userId).Last(&list)
	return list
}

func GreateList(userId string, listTitle string, content string) bool {
	DB.DBconnect.Exec("insert into BUCKETLISTS (USER_ID, LIST_TITLE, CONTENT, LIST_STATUS, CREATION_DATE) values (?, ?, ?, 0, now())",
		userId, listTitle, content)
	return true
}

func UpdateList(listId string, bucketList Bucketlist) Bucketlist {
	DB.DBconnect.Model(&bucketList).Where("List_ID = ?", listId).Updates(bucketList)
	return bucketList
}

func UpdateListStatus(listId string) bool {
	DB.DBconnect.Exec("update BUCKETLISTS set LIST_STATUS = 1 where LIST_ID = ?", listId)
	return true
}

func DeleteList(listId string) bool {
	list := Bucketlist{}
	result := DB.DBconnect.Where("LIST_ID = ?", listId).Delete(&list)
	return result.RowsAffected > 0
}

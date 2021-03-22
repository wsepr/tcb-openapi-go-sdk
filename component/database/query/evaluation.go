package query

import "gopkg.in/mgo.v2/bson"

//取模运算
func (q *Query) Mod(key string, m, result int64) *Query {
	q.query[key] = bson.M{"$mod": []int64{m, result}}
	return q
}

//模糊匹配
func (q *Query) Regex(key string, value interface{}) *Query {
	q.query[key] = bson.M{"$regex": value}
	return q
}

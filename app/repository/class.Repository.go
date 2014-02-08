package repository

import (
	"reflect"

	"github.com/coocood/qbs"
	"github.com/pa001024/MoeWorker/util"
)

type Model interface{}

// 抽象化
type Repository interface {
	Put(Model) Repository
	Delete(Model) Repository
	Get(v Model, key string, value interface{}) Repository
	GetN(structPtr Model, kvs ...interface{}) Repository
	Find(v Model, key string, value interface{}) Repository
	FindN(structPtr Model, n int, kvs ...interface{}) Repository
	Count(structPtr Model, v *int64, key string, value interface{}) Repository
	Close() error
}

// 解耦
func QbsRepo() *QbsRepository {
	q, err := qbs.GetQbs()
	util.Try(err)
	return &QbsRepository{Qbs: q}
}

// 加强代码复用
type QbsRepository struct {
	Repository
	*qbs.Qbs
}

// 自动初始化指针为空值
func newAuto(structPtrPtr interface{}) interface{} {
	val := reflect.ValueOf(structPtrPtr).Elem() // **Model ->  *Model
	if val.IsNil() {
		typ := val.Type().Elem()  // *Model ->  Model
		val.Set(reflect.New(typ)) // *Model = new(Model)
	}
	return val.Interface()
}

// 自动将空值设为nil
func nilAuto(structPtrPtr interface{}) {
	val := reflect.ValueOf(structPtrPtr).Elem() // **Model ->  *Model
	if !val.IsNil() {
		val.Set(reflect.Zero(val.Type())) // *Model = *Model(nil)
	}
}

// 置入
func (this *QbsRepository) Put(structPtr Model) Repository { this.Save(structPtr); return this }

// 删除
func (this *QbsRepository) Delete(structPtr Model) Repository { this.Qbs.Delete(structPtr); return this }

// 获取
func (this *QbsRepository) Get(structPtr Model, key string, value interface{}) Repository {
	this.WhereEqual(key, value).Find(structPtr)
	return this
}

// 获取
func (this *QbsRepository) GetRef(structPtrPtr Model, key string, value interface{}) Repository {
	err := this.WhereEqual(key, value).Find(newAuto(structPtrPtr))
	if err != nil {
		nilAuto(structPtrPtr)
	}
	return this
}

// 多条件获取
func (this *QbsRepository) GetN(structPtr Model, kvs ...interface{}) Repository {
	c := qbs.NewCondition(kvs[0].(string), kvs[1])
	for l, i := len(kvs)-1, 2; i < l; i += 2 {
		c.And(kvs[i].(string), kvs[i+1])
	}
	err := this.Condition(c).Find(structPtr)
	if err != nil {
		structPtr = nil
	}
	return this
}

// 多条件获取
func (this *QbsRepository) GetNRef(structPtrPtr Model, kvs ...interface{}) Repository {
	c := qbs.NewCondition(kvs[0].(string), kvs[1])
	for l, i := len(kvs)-1, 2; i < l; i += 2 {
		c.And(kvs[i].(string), kvs[i+1])
	}
	err := this.Condition(c).Find(newAuto(structPtrPtr))
	if err != nil {
		nilAuto(structPtrPtr)
	}
	return this
}

// 获取
func (this *QbsRepository) Find(structPtr Model, key string, value interface{}) Repository {
	err := this.WhereEqual(key, value).FindAll(structPtr)
	if err != nil {
		structPtr = nil
	}
	return this
}

// 多条件获取
func (this *QbsRepository) FindN(structPtr Model, n int, kvs ...interface{}) Repository {
	c := qbs.NewCondition(kvs[0].(string), kvs[1])
	for l, i := len(kvs)-1, 2; i < l; i += 2 {
		c.And(kvs[i].(string), kvs[i+1])
	}
	if n > 0 {
		this.Limit(n)
	}
	err := this.Condition(c).FindAll(structPtr)
	if err != nil {
		structPtr = nil
	}
	return this
}

// 获取
func (this *QbsRepository) Count(structPtr Model, v *int64, key string, value interface{}) Repository {
	*v = this.WhereEqual(key, value).Count(structPtr)
	return this
}

func (this *QbsRepository) Close() error {
	return this.Qbs.Close()
}

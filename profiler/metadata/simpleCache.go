package metadata

import (
	"sync"
	"sync/atomic"
)

type SimpleCache struct {
	idGen         int32
	ID2MetaInfo   map[int32]string
	MetaInfo2ID   map[string]int32
	ID2MetaObject sync.Map
}

func NewSimpleCache() *SimpleCache {
	return &SimpleCache{
		idGen:       0,
		ID2MetaInfo: make(map[int32]string),
		MetaInfo2ID: make(map[string]int32),
	}
}

func (p *SimpleCache) SetMetaObject(id int32, metaObject interface{}) {
	p.ID2MetaObject.Store(id, metaObject)
}

//返回apiID，和 是否是新的id
func (p *SimpleCache) GetOrCreateID(metaInfo string) (int32, bool) {
	newValue := false
	ID, ok := p.MetaInfo2ID[metaInfo]
	if !ok {
		ID = atomic.AddInt32(&p.idGen, 1)
		p.ID2MetaInfo[ID] = metaInfo
		p.MetaInfo2ID[metaInfo] = ID
		newValue = true
	}

	return ID, newValue
}

func (p *SimpleCache) GetID(metaInfo string) int32 {
	if id, ok := p.MetaInfo2ID[metaInfo]; ok {
		return id
	}

	return 0
}

//根据id获取api
func (p *SimpleCache) GetMetaInfo(id int32) string {
	if api, ok := p.ID2MetaInfo[id]; ok {
		return api
	}
	return ""
}

package Set

import (
	"sync"
)

type set struct {
	m   map[interface{}]bool
	syn sync.Mutex
}

// 创建set
func New() *set {
	return &set{
		m: map[interface{}]bool{},
	}
}

// 判断是否为空
func (s *set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

// 长度
func (s *set) Len() int {
	return len(s.m)
}

// 添加元素
func (s *set) Add(item ...interface{}) {
	s.syn.Lock()
	for _, v := range item {
		s.m[v] = true
	}
	s.syn.Unlock()
}

// 获取所有元素
func (s *set) List() []interface{} {
	s.syn.Lock()
	defer s.syn.Unlock()
	var sl []interface{}
	for k := range s.m {
		sl = append(sl, k)
	}
	return sl
}

// 查找元素是否存在
func (s *set) Has(item interface{}) bool {
	s.syn.Lock()
	defer s.syn.Unlock()
	if _, ok := s.m[item]; ok {
		return true
	}
	return false
}

// 清空
func (s *set) Clear() {
	s.syn.Lock()
	s.m = map[interface{}]bool{}
	s.syn.Unlock()
}

// 删除某个元素
func (s *set) Remove(item interface{}) {
	s.syn.Lock()
	delete(s.m, item)
	s.syn.Unlock()
}

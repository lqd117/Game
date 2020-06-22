package session

import (
	"container/list"
	"sync"
	"time"
)

var provider = &FromMemory{list: list.New()}

func init() {
	provider.sessions = make(map[string]*list.Element, 0)
	//注册  memory 调用的时候一定有一致
	RegisterProvider("memory", provider)
}

//session实现
type SessionStore struct {
	sid              string                      //session id 唯一标示
	LastAccessedTime time.Time                   //最后访问时间
	value            map[interface{}]interface{} //session 里面存储的值
}

//设置
func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	err := provider.SessionUpdate(st.sid)
	if err != nil {
		panic(err)
	}
	return nil
}

//获取session
func (st *SessionStore) Get(key interface{}) interface{} {
	err := provider.SessionUpdate(st.sid)
	if err != nil {
		panic(err)
	}
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
}

//删除
func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	err := provider.SessionUpdate(st.sid)
	if err != nil {
		panic(err)
	}
	return nil
}
func (st *SessionStore) SessionID() string {
	return st.sid
}

//session来自内存 实现
type FromMemory struct {
	lock     sync.Mutex               //用来锁
	sessions map[string]*list.Element //用来存储在内存
	list     *list.List               //用来做 gc
}

func (fromMemory *FromMemory) SessionInit(sid string) (Session, error) {
	fromMemory.lock.Lock()
	defer fromMemory.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newSess := &SessionStore{sid: sid, LastAccessedTime: time.Now(), value: v}
	element := fromMemory.list.PushBack(newSess)
	fromMemory.sessions[sid] = element
	return newSess, nil
}

func (fromMemory *FromMemory) SessionRead(sid string) (Session, error) {

	if element, ok := fromMemory.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		return nil, nil
	}
}

func (fromMemory *FromMemory) SessionDestroy(sid string) error {
	if element, ok := fromMemory.sessions[sid]; ok {
		delete(fromMemory.sessions, sid)
		fromMemory.list.Remove(element)
		return nil
	}
	return nil
}

func (fromMemory *FromMemory) SessionGC(maxLifeTime int64) {
	fromMemory.lock.Lock()
	defer fromMemory.lock.Unlock()
	for {
		element := fromMemory.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).LastAccessedTime.Unix() + maxLifeTime) <
			time.Now().Unix() {
			fromMemory.list.Remove(element)
			delete(fromMemory.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

func (fromMemory *FromMemory) SessionUpdate(sid string) error {
	fromMemory.lock.Lock()
	defer fromMemory.lock.Unlock()
	if element, ok := fromMemory.sessions[sid]; ok {
		element.Value.(*SessionStore).LastAccessedTime = time.Now()
		fromMemory.list.MoveToFront(element)
		return nil
	}
	return nil
}

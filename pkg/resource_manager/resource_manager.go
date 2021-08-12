package resource_manager

type ResourceManager struct {
	resources map[int]interface{}
	ids       chan int
	idx       int
}

func NewResourceManager() *ResourceManager {
	return &ResourceManager{
		resources: map[int]interface{}{},
		ids:       make(chan int),
		idx:       0,
	}
}

func (rm *ResourceManager) AddResource(resource interface{}) {
	resourceId := rm.idx
	rm.resources[resourceId] = resource
	rm.idx++
	go func() {
		rm.ids <- resourceId
	}()
}

func (rm *ResourceManager) Use(task func(resource interface{})) {
	resourceId := <-rm.ids

	task(rm.resources[resourceId])

	go func() {
		rm.ids <- resourceId
	}()
}

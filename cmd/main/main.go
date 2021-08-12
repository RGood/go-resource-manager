package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/RGood/resource_manager/pkg/resource_manager"
)

type TestResource struct {
	id string
}

func NewTestResource(id string) *TestResource {
	return &TestResource{
		id: id,
	}
}

func (tr *TestResource) Announce(message string) {
	fmt.Printf("ID %s: %s\n", tr.id, message)
}

func handleIteration(rm *resource_manager.ResourceManager, id int) {
	rm.Use(func(resource interface{}) {
		// Type-cast our resource to the applicable value
		typedResource, ok := resource.(*TestResource)
		if ok {
			typedResource.Announce(fmt.Sprintf("Iteration %d", id))
		}

		time.Sleep(5 * time.Second)
	})
}

func main() {
	tr1, tr2 := NewTestResource("Foo"), NewTestResource("Bar")

	rm := resource_manager.NewResourceManager()
	rm.AddResource(tr1)
	rm.AddResource(tr2)

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			handleIteration(rm, i)
		}(i)

	}
	wg.Wait()
}

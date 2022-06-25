package assets

import (
	"fmt"
	"log"
)

type EntityManager[T comparable] struct {
	entityAmount   int
	Entities       []T
	deadEntityList []T
	entityToIdx    map[T]int
}

func NewEntityManager[T comparable]() *EntityManager[T] {
	e := EntityManager[T]{}
	e.Entities = make([]T, 0)
	e.deadEntityList = make([]T, 0)
	e.entityToIdx = make(map[T]int)
	e.entityAmount = 0
	return &e
}

func (e *EntityManager[T]) removeEntity(entity T) {
	if _, ok := e.entityToIdx[entity]; !ok {
		log.Printf("Got: %+v", e.entityToIdx)
		panic(fmt.Sprintf("trying to remove non existent entity: %T %+v", entity, entity))
	} else {
		idx := e.entityToIdx[entity]
		if e.entityAmount > 1 {
			swapEntity := e.Entities[e.entityAmount-1]
			delete(e.entityToIdx, e.Entities[e.entityAmount-1])
			e.Entities[idx] = swapEntity
			e.entityToIdx[swapEntity] = idx
		}
		e.entityAmount--
		e.Entities = e.Entities[:e.entityAmount]
	}
}

func (e *EntityManager[T]) AddEntity(entity T) {
	e.Entities = append(e.Entities, entity)
	e.entityToIdx[entity] = e.entityAmount
	e.entityAmount++
}

func (e *EntityManager[T]) SetForRemoval(entity T) {
	e.deadEntityList = append(e.deadEntityList, entity)
}

func (e *EntityManager[T]) CleanDeadEntities() {
	if len(e.deadEntityList) > 0 {
		for i := range e.deadEntityList {
			e.removeEntity(e.deadEntityList[i])
		}
		e.deadEntityList = e.deadEntityList[:0]
	}
}

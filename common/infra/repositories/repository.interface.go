package repository

type IRepository[TDomainEntity any, TDalEntity any] interface {
	FindOneById(id string) (*TDomainEntity, error)
	Save(entity *TDomainEntity) error
	Delete(entity *TDomainEntity) error
	IsExisted(id string) (bool, error)
}

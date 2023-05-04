package repository

type IRepository[TDomainEntity any, TDalEntity any] interface {
	FindOneById(id string) (*TDomainEntity, error)
	Save(entity *TDomainEntity) error
	Update(entity *TDomainEntity) error
	Delete(id string) error
	IsExisted(id string) (bool, error)
}

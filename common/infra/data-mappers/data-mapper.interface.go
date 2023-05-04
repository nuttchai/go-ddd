package data_mapper

type IDataMapper[TDomainEntity any, TDalEntity any] interface {
	ToDomainEntity(dalEntity *TDalEntity) (*TDomainEntity, error)
	ToDalEntity(domainEntity *TDomainEntity) (*TDalEntity, error)
}

package data_mapper

type IDataMapper[TDomainEntity any, TDalEntity any] interface {
	ToDomainEntity(dalEntity *TDalEntity) *TDomainEntity
	ToDalEntity(domainEntity *TDomainEntity) *TDalEntity
}

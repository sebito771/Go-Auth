package mariadb

import (
	"Auth/internal/adapters/repository"
	"Auth/internal/domain/user"

	"gorm.io/gorm"
)

type MariaQueries struct {
	db *gorm.DB
	
}

func NewMariaQueries(db *gorm.DB)*MariaQueries{
 return &MariaQueries{db: db}
}


func (ma *MariaQueries) Save(user *user.User)error{
  model:= FromDomain(user)

  result:= ma.db.Create(model)

  if result.Error != nil{
	return result.Error
  }

  user.SetId(int64(model.Id))
  return nil
}

func (ma *MariaQueries) FindByEmail(email string)(*user.User,error){
	var model MariaModel
	result:=ma.db.Where("email=?",email).First(&model)
	if result.Error!=nil&& result.Error== gorm.ErrRecordNotFound{
		return nil , repository.ErrorNotFound
	}
	if result.Error!=nil{
		return nil , result.Error
	}
	user:= ToDomain(&model)
	return user , nil
}


func (ma *MariaQueries) FindById(id int64)(*user.User,error){
	var model MariaModel
	result:= ma.db.First(&model,id)
	if result.Error!=nil{
		return nil , result.Error
	}
	user:= ToDomain(&model)
	return user,nil
}









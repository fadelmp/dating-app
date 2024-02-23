package mapper

import (
	"dating-app/shared/domain"
	"dating-app/shared/dto"
)

func ToBaseDto(domain domain.Base) dto.Base {

	return dto.Base{
		IsActived: domain.IsActived,
		IsDeleted: domain.IsDeleted,
		CreatedBy: domain.CreatedBy,
		UpdatedBy: domain.UpdatedBy,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

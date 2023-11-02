package repository

import "polygon/internal/entity"

type MemberDTO struct {
	ID        uint64
	FirstName string
	LastName  string
	Age       uint8
}

func (dto *MemberDTO) ToEntity() *entity.Member {
	return &entity.Member{
		ID:        dto.ID,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Age:       dto.Age,
	}
}

package service

import (
	"log"
	"math"
	"myapp/dto"
	"myapp/pkg/repository"
)

type PlaylistService interface {
	GetPlaylist(page, limit int) (currentPage, lastPage, totalRow int, res []dto.MsPlaylist, err error)
	CreatePlaylist(req *dto.MsPlaylist, createdBy string) error
}

type PlaylistServiceImpl struct {
	repo repository.PlaylistRepository
}

func (r *PlaylistServiceImpl) CreatePlaylist(req *dto.MsPlaylist, createdBy string) error {
	if _, err := r.repo.CreatePlaylist(req, createdBy); err != nil {
		return err
	}
	return nil
}

func (p PlaylistServiceImpl) GetPlaylist(page, limit int) (currentPage, lastPage, totalRow int, res []dto.MsPlaylist, err error) {
	result, totalRow, err := p.repo.GetAllPlaylist(page, limit)
	log.Default().Println(result, "di services")
	if err != nil {
		return 0, 0, 0, nil, err
	}
	lastPage = int(math.Ceil(float64(totalRow) / float64(limit)))

	return page, lastPage, totalRow, result, nil
}

func NewPlaylistService(repo repository.PlaylistRepository) *PlaylistServiceImpl {
	return &PlaylistServiceImpl{
		repo: repo,
	}
}

//
//func (r *RoleServiceImpl) GetRoles(page, limit int) (currentPage, lastPage, totalRow int, res []dto.CrRole, err error) {
//	result, totalRow, err := r.repo.GetAllRoles(page, limit)
//	if err != nil {
//		return 0, 0, 0, nil, err
//	}
//
//	lastPage = int(math.Ceil(float64(totalRow) / float64(limit)))
//
//	return page, lastPage, totalRow, result, nil
//}
//
//func (r *RoleServiceImpl) GetRoleById(id int) (*dto.CrRole, error) {
//	return r.repo.GetRoleByID(id)
//}
//
//func (r *RoleServiceImpl) CreateRole(req *dto.CrRole, createdBy string) error {
//	id, err := r.repo.CreateRole(req, createdBy)
//	if err != nil {
//		return err
//	}
//
//	if len(req.RolePermissions) > 0 {
//		for _, i := range req.RolePermissions {
//			if err := r.repo.AssignPermissionToRole(id, i.PermissionID); err != nil {
//				return err
//			}
//		}
//	}
//
//	return nil
//}
//
//func (r *RoleServiceImpl) UpdateRole(req *dto.CrRole, updatedBy string) error {
//	if err := r.repo.UpdateRole(req, updatedBy); err != nil {
//		return err
//	}
//
//	if len(req.RolePermissions) > 0 {
//		if err := r.repo.RemovePermissionsFromRole(req.ID); err != nil {
//			return err
//		}
//
//		for _, i := range req.RolePermissions {
//			if err := r.repo.AssignPermissionToRole(req.ID, i.PermissionID); err != nil {
//				return err
//			}
//		}
//	}
//
//	return nil
//}
//
//func (r *RoleServiceImpl) DeleteRole(id int, deletedBy string) error {
//	return r.repo.DeleteRole(id, deletedBy)
//}

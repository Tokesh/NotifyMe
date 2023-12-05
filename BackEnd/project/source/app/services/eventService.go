package services

import "project/source/domain/entity"

func (s Service) SelectEventBySubId(subs []int) ([]string, error) {
	return s.Repository.SelectEventBySubIdsRepo(subs)
}
func (s Service) SelectEventByEventIds(eventIds []string) ([]entity.Event, error) {
	return s.Repository.SelectEventsByIdRepo(eventIds)
}
func (s Service) SelectEventsByUserId(userId int) ([]entity.Event, error) {
	return s.Repository.SelectEventsByUserIdRepo(userId)
}

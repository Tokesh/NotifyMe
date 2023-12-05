package services

func (s Service) SelectUserSubscription(userId int) ([]int, error) {
	return s.Repository.SelectUserSubscriptionRepo(userId)
}

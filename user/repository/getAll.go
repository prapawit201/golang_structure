package repository

import "log"

// GetAll implements UserRepository.
func (u *userRepository) GetAll() error {
	log.Println("Get all data from repository success")
	return nil
}

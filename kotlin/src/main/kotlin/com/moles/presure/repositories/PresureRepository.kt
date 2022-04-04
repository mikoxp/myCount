package com.moles.presure.repositories

import com.moles.presure.entity.Presure
import org.springframework.data.repository.CrudRepository
import org.springframework.stereotype.Repository

@Repository
interface PresureRepository: CrudRepository<Presure, Integer> {
}
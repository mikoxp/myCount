package com.moles.app.repositories

import com.moles.app.entity.Presure
import org.springframework.data.repository.CrudRepository
import org.springframework.stereotype.Repository

@Repository
interface PresureRepository: CrudRepository<Presure, Integer> {
}
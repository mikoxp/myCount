package com.moles.presure.controllers

import com.moles.presure.entity.Presure
import com.moles.presure.repositories.PresureRepository
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RestController

@RestController
class PresureController(private val repository:PresureRepository) {



    @GetMapping("/logs")
    fun logsAll(): List<Presure> {
        val findAll = repository.findAll()
        val toList = findAll.iterator().asSequence().toList()
        return toList
    }
}
package com.moles.app.controllers

import com.moles.app.entity.Presure
import com.moles.app.repositories.PresureRepository
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RestController

@RestController
class PresureController(private val repository: PresureRepository) {



    @GetMapping("/logs")
    fun logsAll(): List<Presure> {
        val findAll = repository.findAll()
        val toList = findAll.iterator().asSequence().toList()
        return toList
    }
}
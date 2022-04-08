package com.moles.app.entity

import java.time.LocalDateTime
import javax.persistence.*
import javax.print.attribute.IntegerSyntax

@Entity
@Table(name = "PRESURE")
data class Presure(

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    var id: Int = -1,

    @Column(name = "p_date", columnDefinition = "TIMESTAMP")
    var date: LocalDateTime,

    var systolic: Int,

    var diastolic: Int,

    var pulse: Int

) {}
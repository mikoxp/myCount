import moles.proccessing.Presure
import org.apache.poi.ss.usermodel.Cell
import org.apache.poi.ss.usermodel.WorkbookFactory
import java.io.FileInputStream


fun readToObjects(filepath: String): List<Presure> {
    val result: MutableList<Presure> = mutableListOf()
    val inputStream = FileInputStream(filepath)
    var xlWb = WorkbookFactory.create(inputStream)
    var rowNumber = 1
    val xlWs = xlWb.getSheetAt(0)
    while (true) {
        val row = xlWs.getRow(rowNumber) ?: break
        val cell = row.getCell(0)
        val tmp = cell.stringCellValue
        result.add(Presure(tmp, intValue(row.getCell(1)), intValue(row.getCell(2)), intValue(row.getCell(3))))
        rowNumber += 1
    }
    return result
}

fun intValue(cell: Cell): Int {
    return cell.numericCellValue.toInt()
}

fun main() {
    val filepath = "./data.xlsx"
    val readToObjects = readToObjects(filepath)
    println(readToObjects)
}
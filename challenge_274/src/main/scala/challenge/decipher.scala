package challenge

import scala.io.Source
import java.io.InputStream

object Challenge274 {
  def decipher(cipher: List[Int], words: List[String], decipheredWord: List[String]): String = cipher match {
    case Nil => decipheredWord.mkString
    case x::xs => decipher(xs, words, decipheredWord :+ words((x-1) % 1322)(0).toString)
  }
}

object Main extends App {
  val wordsStream : InputStream = getClass.getResourceAsStream("/declaration_of_independence.txt")
  val words = Source.fromInputStream(wordsStream).getLines.mkString.split(" ").toList
  val cipherStream : InputStream = getClass.getResourceAsStream("/cipher.txt")
  val cipherText = Source.fromInputStream(cipherStream).getLines.mkString.replaceAll(" ", "").split(",").toList.map(_.toInt)

  val decipheredText = Challenge274.decipher(cipherText, words, List())
  println(decipheredText)
}

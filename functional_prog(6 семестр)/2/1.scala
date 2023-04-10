import scala.math.{min, max}

case class Point(x: Double, y: Double)

class PointsSet(val points: List[Point]) {
  def size: Int = points.length
  def add(point: Point): PointsSet = new PointsSet(point :: points)
  def contains(point: Point): Boolean = points.contains(point)
  def +(that: PointsSet): PointsSet = new PointsSet(this.points ++ that.points)
  def -(that: PointsSet): PointsSet = new PointsSet(this.points.filterNot(that.points.contains))
}
object PointsSet {
  def rectangle(bottomLeft: Point, topRight: Point): PointsSet = {
    val xs = bottomLeft.x.toInt to topRight.x.toInt
    val ys = bottomLeft.y.toInt to topRight.y.toInt
    val points = for (x <- xs; y <- ys) yield Point(x.toDouble, y.toDouble)
    new PointsSet(points.toList)
  }
  
  def circle(center: Point, radius: Double): PointsSet = {
    val numPoints = (radius * 2 * math.Pi).ceil.toInt
    val points = for {
      i <- 0 until numPoints
      angle = i.toDouble / numPoints * 2 * scala.math.Pi
      r <- 1 to radius.toInt
      x = center.x + r * scala.math.cos(angle)
      y = center.y + r * scala.math.sin(angle)
    } yield Point(x, y)
    new PointsSet(points.toList)
  }
}

//20. Функция powers: (List[Int], Int ^> Boolean) ^> List[Int], удаляющая
//из списка те числа, которые не являются заданными предикатом степенями числа
//2.
val ispowers: Int => Int = {
  case 1 => 0
  case x if (x < 1) => -65
  case x if (x%2 == 1) => -65
  case x => ispowers(x / 2) + 1
}
val powers: (List[Int], Int => Boolean) => List[Int] = {
  case (Nil, _) => Nil
  case (x::xs, p) if ((ispowers(x) >=0) && p(ispower(x)))  => x::powers(xs, p)
  case (x::xs, p) => powers(xs, p)
}
println(powers(List(0,1,2,3,4,5,6,8,10), _ % 4 == 0))

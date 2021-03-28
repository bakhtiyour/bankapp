package transfer

//Total рассчитывает параметры вклада для Корти Милли
func Total(amount int) (totalSum int) {


	totalSum=amount+amount*5/1000

	return totalSum
}
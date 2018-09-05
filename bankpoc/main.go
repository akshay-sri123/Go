package main

func main()  {
  var customer *Customer
  customer = NewCustomer("Akshay", "Srivastava", "123", 500.00)
  customer.printDetails()
  customer.deposit(1002.23)
  customer.printDetails()
  customer.withdraw(500)
  customer.printDetails()
}

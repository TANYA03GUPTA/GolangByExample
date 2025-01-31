package main
type iPlayer interface{
	getSymbol() Symbol
	getNextMove() (int,int,error)
	getid()int
}
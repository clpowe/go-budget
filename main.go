package main

import (
	"math/rand"

	"github.com/gin-gonic/gin"
)

type Expense struct {
		ID int `json:"id"`
		Name string `json:"name"`
		Amount int `json:"amount"`
		BudgetID int `json:"budget_id"`
	}

type Budget struct {
		ID int `json:"id"`
		Name string `json:"name"`
		Amount int `json:"amount"`
}

var budgets []Budget
var expenses []Expense

func pong(c *gin.Context){
		c.JSON(200, gin.H{
			"message":"pong",
	})
}

func getBudgets(c *gin.Context){
	c.JSON(200,budgets)
}
func getExpenses( c *gin.Context){
	c.JSON(200,expenses )
}

func createBudget(c *gin.Context){
	var budget Budget
	if c.ShouldBind(&budget) == nil {
		budget.ID = rand.Intn(10000000)
		budgets = append(budgets, budget)
	}
}
func createExpense(c *gin.Context){
	var expense Expense
	if c.ShouldBind(&expense) == nil {
		expense.ID = rand.Intn(10000000)
		expenses = append(expenses, expense)
	}
}

func deleteExpense(c *gin.Context){
	var expense Expense
	if c.ShouldBind(&expense) == nil {
		for index, item := range expenses {
			if item.ID == expense.ID {
				expenses = append(expenses[:index], expenses[index+1:]...)
				break
			}
	}
	}
}

func main()  {
	router := gin.Default()


	budgets = append(budgets, Budget{ID: 1,Name: "Food", Amount: 300})
	budgets = append(budgets, Budget{ID: 2,Name: "Car", Amount: 200})
	
	expenses = append(expenses, Expense{ID: 1, Name: "McDonalds", Amount:30, BudgetID: 1})
	expenses = append(expenses, Expense{ID: 2, Name: "Gas", Amount: 25, BudgetID: 2})

	router.GET("/ping", pong)
	router.GET("/budgets", getBudgets)
	router.GET("/expenses", getExpenses)
	router.POST("/createBudget", createBudget)
	router.POST("/createExpense", createExpense)
	router.DELETE("/deleteExpense", deleteExpense)

	router.Run()
}
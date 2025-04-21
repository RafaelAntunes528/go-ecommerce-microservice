package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "ecomcli",
		Short: "E-Commerce CLI Tool",
	}

	var listProductsCmd = &cobra.Command{
		Use:   "list-products",
		Short: "List all products",
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := http.Get("http://localhost:8080/products")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer resp.Body.Close()

			var products []struct {
				ID    string  `json:"id"`
				Name  string  `json:"name"`
				Price float64 `json:"price"`
			}

			if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
				fmt.Println("Error decoding response:", err)
				return
			}

			fmt.Println("Available Products:")
			fmt.Println("------------------")
			for _, p := range products {
				fmt.Printf("%s: %s ($%.2f)\n", p.ID, p.Name, p.Price)
			}
		},
	}

	var createOrderCmd = &cobra.Command{
		Use:   "create-order",
		Short: "Create a new order",
		Run: func(cmd *cobra.Command, args []string) {
			productID, _ := cmd.Flags().GetString("product-id")

			order := map[string]string{
				"product_id": productID,
				"status":     "pending",
			}

			jsonData, _ := json.Marshal(order)
			resp, err := http.Post("http://localhost:8080/orders", "application/json", bytes.NewBuffer(jsonData))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusCreated {
				fmt.Println("Order created successfully!")
			} else {
				fmt.Println("Failed to create order. Status:", resp.Status)
			}
		},
	}

	createOrderCmd.Flags().String("product-id", "", "Product ID to order")
	createOrderCmd.MarkFlagRequired("product-id")

	rootCmd.AddCommand(listProductsCmd, createOrderCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

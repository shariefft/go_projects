# Revenue and Expense Calculator

This Go program calculates the **Earnings Before Tax (EBT)**, **Profit**, and **Profit-to-EBT Ratio** based on the user's input for revenue, expenses, and tax rate. The result is displayed in the console and saved to a file named `result.txt`.

## Features

- Prompts the user to input **revenue**, **expenses**, and **tax rate**.
- Validates that the inputs are positive numbers.
- Calculates:
  - **Earnings Before Tax (EBT)**: Revenue minus expenses.
  - **Profit**: EBT after applying the tax rate.
  - **Profit-to-EBT Ratio**: Ratio of EBT to profit.
- Writes the calculation results to a file named `result.txt`.

## How to Use

1. Run the program.
2. Provide the required inputs when prompted:
   - **Revenue**: The total revenue (should be a positive number).
   - **Expenses**: The total expenses (should be a positive number).
   - **Tax Rate**: The tax rate as a percentage (should be a positive number).
3. The program will display the following results:
   - **EBT**: Earnings before tax.
   - **Profit**: The calculated profit after applying the tax rate.
   - **Profit-to-EBT Ratio**: Ratio between EBT and profit.
4. The results are also saved in the file `result.txt`.

### Example Input:

```plaintext
Revenue: 5000
Expenses: 2000
Tax Rate: 20

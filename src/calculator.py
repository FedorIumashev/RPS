# This is a simple calculator module that provides basic arithmetic operations.
# It includes functions for addition, subtraction, multiplication, and division.
# Each function takes two arguments and returns the result of the operation.
def add(a, b):
    return a + b

def subtract(a, b):
    return a - b

def multiply(a, b):
    return a * b

def divide(a, b):
    if b == 0:
        raise ValueError("Division by zero is not allowed")
    return a / b

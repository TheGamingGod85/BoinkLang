def verify_pin(user_pin, actual_pin):
    return user_pin == actual_pin

def atm_simulator():
    # Default PIN and balance
    pin = "1234"
    balance = 1000.0

    print("===== Welcome to the ATM Simulator =====")

    try:
        # PIN verification loop
        attempts = 3
        while attempts > 0:
            entered_pin = input("Enter your 4-digit PIN: ")
            if verify_pin(entered_pin, pin):
                break
            else:
                attempts -= 1
                print(f"Incorrect PIN. {attempts} attempt(s) left.")
        else:
            print("Too many incorrect attempts. Exiting...")
            return

        # Main operation loop
        while True:
            print("\nChoose an operation:")
            print("1. Check Balance")
            print("2. Withdraw Money")
            print("3. Deposit Money")
            print("4. Change PIN")
            print("5. Exit")

            choice = input("Enter your choice (1-5): ")

            if choice == '1':
                print(f"Your current balance is ₹{balance:.2f}")

            elif choice == '2':
                try:
                    amount = float(input("Enter amount to withdraw: ₹"))
                    if amount <= 0:
                        print("Enter a positive amount.")
                    elif amount > balance:
                        print("Insufficient balance.")
                    else:
                        balance -= amount
                        print(f"₹{amount:.2f} withdrawn. New balance: ₹{balance:.2f}")
                except ValueError:
                    print("Invalid amount entered.")

            elif choice == '3':
                try:
                    amount = float(input("Enter amount to deposit: ₹"))
                    if amount <= 0:
                        print("Enter a positive amount.")
                    else:
                        balance += amount
                        print(f"₹{amount:.2f} deposited. New balance: ₹{balance:.2f}")
                except ValueError:
                    print("Invalid amount entered.")

            elif choice == '4':
                old_pin = input("Enter your current PIN: ")
                if verify_pin(old_pin, pin):
                    new_pin = input("Enter new 4-digit PIN: ")
                    if len(new_pin) == 4 and new_pin.isdigit():
                        pin = new_pin
                        print("PIN successfully changed.")
                    else:
                        print("PIN must be 4 digits.")
                else:
                    print("Incorrect current PIN.")

            elif choice == '5':
                print("Thank you for using the ATM. Goodbye!")
                break

            else:
                print("Invalid choice. Please enter a number from 1 to 5.")

    except KeyboardInterrupt:
        print("\nSession terminated by user.")
    except Exception as e:
        print(f"An unexpected error occurred: {e}")

# Run the simulator
atm_simulator()
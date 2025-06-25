def calculate_grade(avg):
    if avg >= 90:
        return 'A'
    elif avg >= 80:
        return 'B'
    elif avg >= 70:
        return 'C'
    elif avg >= 60:
        return 'D'
    else:
        return 'F'

def student_marks_system():
    students = []

    try:
        n = int(input("Enter the number of students: "))
        for i in range(n):
            print(f"\nEnter details for Student {i + 1}:")
            name = input("Name: ")
            marks = []
            for j in range(1, 4):
                while True:
                    try:
                        mark = float(input(f"Enter marks for Subject {j}: "))
                        if 0 <= mark <= 100:
                            marks.append(mark)
                            break
                        else:
                            print("Marks must be between 0 and 100.")
                    except ValueError:
                        print("Please enter a valid number.")

            total = sum(marks)
            average = total / 3
            grade = calculate_grade(average)

            students.append({
                "name": name,
                "marks": marks,
                "total": total,
                "average": average,
                "grade": grade
            })

        # Display results
        print("\n==== Student Report ====")
        for s in students:
            print(f"\nName: {s['name']}")
            print(f"Marks: {s['marks']}")
            print(f"Total: {s['total']}")
            print(f"Average: {s['average']:.2f}")
            print(f"Grade: {s['grade']}")

        # Highest and lowest scorer
        highest = max(students, key=lambda x: x['total'])
        lowest = min(students, key=lambda x: x['total'])

        print("\n==== Top Performer ====")
        print(f"{highest['name']} with {highest['total']} marks.")

        print("\n==== Lowest Performer ====")
        print(f"{lowest['name']} with {lowest['total']} marks.")

    except ValueError:
        print("Invalid input! Please enter numbers where required.")
    except KeyboardInterrupt:
        print("\nProgram interrupted by user.")
    except Exception as e:
        print(f"An unexpected error occurred: {e}")

# Run the system
student_marks_system()
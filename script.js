/*

    TO DO:
    A Company has many Departments, and each Department has many Employees.
    Add Filter to the search function (Day, Time of Clock in/out, Employee Name, Department Name)
    The user (employee) can clock in/out.
    The user (employee) can see their clock in/out history.
    If the user is already login and tries to login again, show a message "You are already logged in".
    The user's current session should be saved in local storage.

*/
let isClockedIn = false;
const btnClockIn = document.getElementById("btnClockIn").addEventListener("click", () => {
    if (isClockedIn) {
        alert("You are already clocked in.");
        return;
    } else {
        isClockedIn = true;
        alert("You have clocked in successfully.");
    }
});
const btnClockOut = document.getElementById("btnClockOut").addEventListener("click", () => {
    if (!isClockedIn) {
        alert("You are already clocked out.");
        return;
    }
    isClockedIn = false;
    alert("You have clocked out successfully.");
});
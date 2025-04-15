/*

    TO DO:
    A Company has many Departments, and each Department has many Employees.
    Add Filter to the search function (Day, Time of Clock in/out, Employee Name, Department Name)
    The user (employee) can clock in/out.
    The user (employee) can see their clock in/out history.
    If the user is already login and tries to login again, show a message "You are already logged in".
    The user's current session should be saved in local storage.

    ! user login credentials are in scriptlogin.js !

*/
const logTable = document.getElementById("logTable");
const newRow = logTable.insertRow();
let isClockedIn = false;
const btnClockIn = document.getElementById("btnClockIn").addEventListener("click", () => {
    if (isClockedIn) {
        alert("You are already clocked in.");
        return;
    } else {
        isClockedIn = true;
        const currentTime = new Date();
        alert("You have clocked in successfully.\n" + currentTime.toLocaleString());
        // Append new clock in here
        
        const cellDepartment = newRow.insertCell(0);
        cellDepartment.innerHTML = "Department Name";
        const cellName = newRow.insertCell(1);
        cellName.innerHTML = "Employee Name";
        const cellClockIn = newRow.insertCell(2);
        cellClockIn.innerHTML = currentTime.toLocaleString();
    }
});
const btnClockOut = document.getElementById("btnClockOut").addEventListener("click", () => {
    if (!isClockedIn) {
        alert("You are already clocked out.");
        return;
    }
    isClockedIn = false;
    const currentTime = new Date();
    alert("You have clocked out successfully.\n" + new Date().toLocaleString());
    // Append new clock out here
        const cellClockOut = newRow.insertCell(3);
        cellClockOut.innerHTML = currentTime.toLocaleString();
});
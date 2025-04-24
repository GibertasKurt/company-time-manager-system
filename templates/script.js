/*
    TO DO:
    A Company has many Departments, and each Department has many Employees.
    Add Filter to the search function (Day, Time of Clock in/out, Employee Name, Department Name)
    The user (employee) can clock in/out.
    The user (employee) can see their clock in/out history.
    If the user is already login and tries to login again, show a message "You are already logged in".
    The user's current session should be saved in local storage.
*/
const logTable = document.getElementById("logTable");
let isClockedIn = false;
let currentRow = null;

const btnClockIn = document.getElementById("btnClockIn").addEventListener("click", () => {
    if (isClockedIn) {
        alert("You are already clocked in.");
        return;
    }
    isClockedIn = true;
    const currentTime = new Date();
    alert("You have clocked in successfully at " + currentTime.toLocaleString());

    currentRow = logTable.insertRow();
    currentRow.insertCell(0).innerHTML = "Department Name"; // Replace with actual department
    currentRow.insertCell(1).innerHTML = "Employee Name"; // Replace with actual employee name
    currentRow.insertCell(2).innerHTML = currentTime.toLocaleString();
});

const btnClockOut = document.getElementById("btnClockOut").addEventListener("click", () => {
    if (!isClockedIn) {
        alert("You are already clocked out.");
        return;
    }
    isClockedIn = false;
    const currentTime = new Date();
    alert("You have clocked out successfully at " + currentTime.toLocaleString());

    if (currentRow) {
        currentRow.insertCell(3).innerHTML = currentTime.toLocaleString();
    }
});

document.getElementById("filterBtn").addEventListener("click", () => {
    const filterValue = document.getElementById("search").value.toLowerCase();
    const rows = logTable.getElementsByTagName("tr");
    for (let i = 1; i < rows.length; i++) {
        const cells = rows[i].getElementsByTagName("td");
        let rowVisible = false;
        for (let j = 0; j < cells.length; j++) {
            if (cells[j].innerHTML.toLowerCase().includes(filterValue)) {
                rowVisible = true;
                break;
            }
        }
        rows[i].style.display = rowVisible ? "" : "none";
    }
});

let url = "/uadmin/api/d/clockhistory/"
fetch(url, {
    method: "POST",
    headers: {
        "X-CSRF-TOKEN": getCookie('session'),
    }, body: logTable
}).then(response => {
    return response.json();
}).then(response => {
    if(response.status == "ok"){
        alert(
            "Clock history fetched successfully"
        )
    }else{
        alert(
            "Error fetching clock history"
        )
    }
})
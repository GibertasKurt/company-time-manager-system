/* Following is set in home.html
    console.log("EmpID: ", `{{.EmployeeID}}`);
    let empid = '{{.EmployeeID}}';
    console.log("currentID: ", `{{.Current.ID}}`);
    let current_id = '{{.Current.ID}}'; */
const logTable = document.getElementById("logTable");
let isClockedIn = localStorage.getItem("recentlogin") == 0 ? false : true;
let isBreakStarted;
let clockInTime, breakStartTime, breakEndTime, timestamp;
//// Important backend stuff
function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
};
localStorage.setItem("recentlogin", current_id);
//// Buttons
const btnClockIn = document.getElementById("btnClockIn").addEventListener("click", () => {

    if (isClockedIn) {
        alert("You are already clocked in.");
        return;
    }
    isClockedIn = true;
    clockInTime = new Date();
    alert("You have clocked in successfully at " + clockInTime.toLocaleString());
    currentRow = logTable.insertRow();

    const logData = {
        "_employee_id": empid,
    };
    const url = `/admin/api/d/clockhistory/add/?_employee_id=${empid}&x-csrf-token=${getCookie("session")}`;
    fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "X-CSRF-TOKEN": getCookie("session"),
        },
        body: JSON.stringify(logData),
    })
    .then(response => response.json())
    .then(response => {
        console.log("logData: ", JSON.stringify(logData))
        console.log("response: ", response)

        if (response.status === "ok") {
            console.log("Data successfully sent to uadmin:", response);
        } else {
            alert("Error sending data to uadmin.");
        }
    })
    .catch(error => {
        console.error("Error:", error);
        alert("An error occurred while sending data to uadmin.");
    });
});

//break start
const btnbrkstrt = document.getElementById("btnbrkstrt").addEventListener("click", () => {
    if (!isClockedIn) {
        alert("You need to clock in first.");
        return;
    }
    breakStartTime = Date.now();
    console.log("Break started at " + breakStartTime);
    isBreakStarted = true;
    console.log(isBreakStarted);

    const logData = {
        "_break_start": breakStartTime,
    };
    const url = `/admin/api/d/clockhistory/edit/${current_id}/?_break_start=${logData}&x-csrf-token=${getCookie("session")}`; // Update this, king.
    fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(logData),
    })
    .then(response => response.json())
    .then(response => {
        if (response.status === "ok") {
            console.log("Break start time successfully sent to uadmin:", response);
        } else {
            alert("Error sending break start time to uadmin.");
        }
    })
    .catch(error => {
        console.error("Error:", error);
        alert("An error occurred while sending break start time to uadmin.");
    });
});
const btnbrkend = document.getElementById("btnbrkend").addEventListener("click", () => {
    if (!isBreakStarted) {
        alert("You need to start a break first.");
        return;
    }
    breakEndTime = new Date();
    alert("Break ended at " + breakEndTime.toLocaleString());

    const logData = {
        "_break_end": breakEndTime.toLocaleString(),
    };
    const url = `/admin/api/d/clockhistory/edit/${current_id}/?_break_end=${logData}&x-csrf-token=${getCookie("session")}`; // Update this, king.
    fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(logData),
    })
    .then(response => response.json())
    .then(response => {
        if (response.status === "ok") {
            console.log("Break end time successfully sent to uadmin:", response);
        } else {
            alert("Error sending break end time to uadmin.");
        }
    })
    .catch(error => {
        console.error("Error:", error);
        alert("An error occurred while sending break start time to uadmin.");
    });
});
const btnClockOut = document.getElementById("btnClockOut").addEventListener("click", () => {
    if (!isClockedIn) {
        alert("You are already clocked out.");
        return;
    }
    if (!isBreakStarted) {
        alert("You need to start a break first.");
        return;
    }
    if (!breakEndTime) {
        alert("You need to end the break first.");
        return;
    }
    isClockedIn = false;
    const clockOutTime = new Date();
    alert("You have clocked out successfully at " + clockOutTime.toLocaleString());
    const totalHours = ((clockOutTime - clockInTime) - (breakEndTime - breakStartTime)) / (1000 * 60 * 60);
    const logData = {
        "_clock_out": clockOutTime.toLocaleString(),
    };
    const url = `/admin/api/d/clockhistory/edit/${current_id}/?_clock_out=${logData}&x-csrf-token=${getCookie("session")}`; // Update this, king.
    fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(logData),
    })
    .then(response => response.json())
    .then(response => {
        if (response.status === "ok") {
            console.log("Clock out time successfully sent to uadmin:", response);
        } else {
            alert("Error sending clock out time to uadmin.");
        }
    })
    .catch(error => {
        console.error("Error:", error);
        alert("An error occurred while sending clock out time to uadmin.");
    });
});
// Convert go lang date & time in table to PREFERRED date and time
document.addEventListener("DOMContentLoaded", function () {

    const dateCells = document.querySelectorAll('.date');
    dateCells.forEach(cell => {
        const originalDate = cell.textContent.trim();
        if (originalDate) {
            const convertedDate = convertDate(originalDate);
            cell.textContent = convertedDate;
        }
    });
});
function convertDate(dateStr) {
    const date = new Date(dateStr);
    if (isNaN(date.getTime())) {
        return ""; // <nil> STRING BECOMES EMPTY, DAZ RITE MFER.
    }
    const options = {
        year: 'numeric',
        month: 'numeric',
        day: 'numeric',
        hour: 'numeric',
        minute: 'numeric',
        second: 'numeric',
        hour12: true,
    };
    const formattedDate = date.toLocaleString('en-US', options);
    return formattedDate.replace(',', '');
}
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
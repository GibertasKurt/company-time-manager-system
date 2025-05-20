const logTable = document.getElementById("logTable");
let currentRow = logTable.rows.length > 0 ? logTable.rows[logTable.rows.length - 1] : null;
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
//// Get CurrentRow
function getCurrentNullCell(currentRow) {
    if (!currentRow) {
        console.error("currentRow is undefined.");
        return null;
    }
    const cells = currentRow.cells;
    for (let i = 0; i < cells.length; i++) {
        if (!cells[i].innerHTML.trim()) {
            return cells[i];
        }
    }
    return null;
}
//// Buttons
const btnClockIn = document.getElementById("btnClockIn").addEventListener("click", () => {

    if (isClockedIn) {
        alert("You are already clocked in.");
        return;
    }
    isClockedIn = true;
    clockInTime = new Date();
    const nullCell = getCurrentNullCell(currentRow);
    nullCell ? console.log("Found an empty cell:", nullCell) : console.log("No empty cells found in the current row.");
    currentRow = logTable.insertRow();
    currentRow.insertCell(0).innerHTML = departmentName.innerHTML;
    currentRow.insertCell(1).innerHTML = Username;
    currentRow.insertCell(2).innerHTML = clockInTime.toLocaleString();
    alert("You have clocked in successfully at " + clockInTime.toLocaleString());

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

const btnbrkstrt = document.getElementById("btnbrkstrt");
btnbrkstrt.addEventListener("click", () => {
    if (!isClockedIn) {
        alert("You need to clock in first.");
        return;
    }
    const nullCell = getCurrentNullCell(currentRow);
    nullCell ? console.log("Found an empty cell:", nullCell) : console.log("No empty cells found in the current row.");
    isBreakStarted = true;

    let formData = new FormData()
    formData.append("data_value", btnbrkstrt.getAttribute("data-value"))

    const url = "/api/update_break";
    fetch(url, {
        method: "PATCH",
        body: formData,
    })
        .then(response => {
            // console.log(response)
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.text();
        })
        .then(text => {
            try {
                // console.log(text);
            } catch (e) {
                console.error('Failed to parse JSON:', e);
            }
        })
        .catch(error => {
            console.error("Error:", error);
            alert("An error occurred while sending break start time to uadmin.");
        });
});
const btnbrkend = document.getElementById("btnbrkend");
btnbrkend.addEventListener("click", () => {
    if (!isBreakStarted) {
        alert("You need to start a break first.");
        return;
    }
    const nullCell = getCurrentNullCell(currentRow);
    nullCell ? console.log("Found an empty cell:", nullCell) : console.log("No empty cells found in the current row.");
    const logData = {
        "data_value": btnbrkstrt.getAttribute("data-value")
    };

    let formData = new FormData()
    formData.append("data_value", btnbrkend.getAttribute("data-value"))

    const url = "/api/update_break";
    fetch(url, {
        method: "PATCH",
        body: formData,
    })
        .then(response => {
            // console.log(response)
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.text();
        })
        .then(text => {
            try {
                // console.log(text);
            } catch (e) {
                console.error('Failed to parse JSON:', e);
            }
        })
        .catch(error => {
            console.error("Error:", error);
            alert("An error occurred while sending break end time to uadmin.");
        });
});
const btnClockOut = document.getElementById("btnClockOut").addEventListener("click", () => {
    if (!isClockedIn) {
        alert("You are already clocked out.");
        return;
    }
    if (!breakEndTime) {
        alert("You need to end the break first.");
        return;
    }
    isClockedIn = false;
    const clockOutTime = new Date();
    const nullCell = getCurrentNullCell(currentRow);
    nullCell ? console.log("Found an empty cell:", nullCell) : console.log("No empty cells found in the current row.");
    currentRow.insertCell(5).innerHTML = clockOutTime.toLocaleString();
    alert("You have clocked out successfully at " + clockOutTime.toLocaleString());
    const totalHours = ((clockOutTime - clockInTime) - (breakEndTime - breakStartTime)) / (1000 * 60 * 60);
    const logData = {
        "_clock_out": clockOutTime.toLocaleString(),
    };
    const url = "/api/update_break";
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
        return ""; // <nil> BECOMES EMPTY STRING, DAZ RITE MFER.
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
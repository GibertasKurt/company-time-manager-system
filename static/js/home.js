const logTable = document.getElementById("logTable");
let isClockedIn = localStorage.getItem("recentlogin") == 0 || localStorage.getItem("recentlogin") == null ? false : true;
let isBreakStarted = localStorage.getItem("breakstarted") == 0 || localStorage.getItem("breakstarted") == null ? false : true;
let clockInTime, breakStartTime, breakEndTime, clockOutTime;
let currentRow = null;
var popupDialog = document.getElementById("popupDialog");
var popupDialogclose = document.getElementsByClassName("close")[0];
const popupDialogText = document.getElementById("popupDialog-text");
function getCookie(name) { //// Important backend stuff
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
};
localStorage.setItem("recentlogin", current_id);
localStorage.setItem("breakstarted", breakstarted_id);
const btnClockIn = document.getElementById("btnClockIn"); //// Buttons
btnClockIn.addEventListener("click", () => {
    if (isAdmin) {
        popupDialog.style.display = "block";
        popupDialogText.innerHTML = "You cant clock in as an admin.";
        return;
    }
    if (isClockedIn) {
        popupDialog.style.display = "block";
        popupDialogText.innerHTML = "You are already clocked in.";
        return;
    }
    isClockedIn = true;
    clockInTime = new Date();
    currentRow = logTable.insertRow();
    departmentName = document.getElementById("departmentName");
    console.log("Department Name: ", departmentName.getAttribute("data-value"));
    currentRow.insertCell(0).innerHTML = departmentName.innerHTML;
    currentRow.insertCell(1).innerHTML = `{{.Username}}`;
    currentRow.insertCell(2).innerHTML = clockInTime.toLocaleString();
    let formData = new FormData()
    formData.append("data_value", btnClockIn.getAttribute("data-value"))
    const url = "/api/updateclock";
    fetch(url, {
        method: "POST",
        body: formData,
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.text();
        })
        .then(text => {
            try {
                console.log('Successfully sent clock in time to uadmin!');
                location.reload();
            } catch (e) {
                console.error('Failed to parse JSON:', e);
            }
        })
        .catch(error => {
            console.error("Error:", error);
            alert("An error occurred while sending clock in time to uadmin.");
        });
});
const btnbrkstrt = document.getElementById("btnbrkstrt");
btnbrkstrt.addEventListener("click", () => {
    if (!isClockedIn) {
        popupDialog.style.display = "block";
        popupDialogText.innerHTML = "You need to clock in first.";
        return;
    }
    if (isBreakStarted) {
        popupDialog.style.display = "block";
        popupDialogText.innerHTML = "You are already taking a break.";
        return;
    }
    if (tookabreak) {
        popupDialog.style.display = "block";
        popupDialogText.innerHTML = "You have already taken a break today.";
        return;
    }
    isBreakStarted = true;
    breakStartTime = new Date();
    let formData = new FormData()
    formData.append("data_value", btnbrkstrt.getAttribute("data-value"))
    const url = "/api/updateclock";
    fetch(url, {
        method: "PATCH",
        body: formData,
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.text();
        })
        .then(text => {
            try {
                console.log('Successfully sent break start to uadmin!');
                location.reload();
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
    if (!isClockedIn) {
        popupDialog.style.display = "block";
        popupDialogText.innerHTML = "You need to clock in first.";
        return;
    }
    if (!isBreakStarted) {
        popupDialog.style.display = "block";
        popupDialogText.innerHTML = "You need to start a break first.";
        return;
    }
    isBreakStarted = false;
    tookabreak = true
    breakEndTime = new Date();
    let formData = new FormData()
    formData.append("data_value", btnbrkend.getAttribute("data-value"))
    const url = "/api/updateclock";
    fetch(url, {
        method: "PATCH",
        body: formData,
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.text();
        })
        .then(text => {
            try {
                console.log('Successfully sent break end to uadmin!');
                location.reload();
            } catch (e) {
                console.error('Failed to parse JSON:', e);
            }
        })
        .catch(error => {
            console.error("Error:", error);
            alert("An error occurred while sending break end time to uadmin.");
        });
});
const btnClockOut = document.getElementById("btnClockOut");
btnClockOut.addEventListener("click", () => {
    if (!isClockedIn) {
        popupDialog.style.display = "block";
        popupDialogText.innerHTML = "You are already clocked out.";
        return;
    }
    if (isBreakStarted) {
        popupDialog.style.display = "block";
        popupDialogText.innerHTML = "You cannot clock out without ending your break first.";
        return;
    }
    isBreakStarted = false
    isClockedIn = false
    clockOutTime = new Date();
    let formData = new FormData()
    formData.append("data_value", btnClockOut.getAttribute("data-value"))
    const url = "/api/updateclock";
    fetch(url, {
        method: "PATCH",
        body: formData,
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.text();
        })
        .then(text => {
            try {
                console.log('Successfully sent clock out to uadmin!');
                location.reload();
            } catch (e) {
                console.error('Failed to parse JSON:', e);
            }
        })
        .catch(error => {
            console.error("Error:", error);
            alert("An error occurred while sending clock out time to uadmin.");
        });
});
document.addEventListener("DOMContentLoaded", function () { // Convert go lang date & time in table to PREFERRED date and time
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
    if (isNaN(date.getTime())) { return "" } // <nil> BECOMES EMPTY STRING, DAZ RITE MFER.
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
};
document.getElementById("filterBtn").addEventListener("click", () => { // Filter function
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
popupDialogclose.onclick = function() { // Popup Dialog
    popupDialog.style.display = "none";
    location.reload();
};
window.onclick = function(event) {
    if (event.target == popupDialog) {
        popupDialog.style.display = "none";
        location.reload();
    }
};
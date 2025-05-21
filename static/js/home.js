const logTable = document.getElementById("logTable");

let isClockedIn = localStorage.getItem("recentlogin") == 0 || localStorage.getItem("recentlogin") == null ? false : true;
let isBreakStarted;
let clockInTime, breakStartTime, breakEndTime, clockOutTime;
var popupDialog = document.getElementById("popupDialog");
var popupDialogclose = document.getElementsByClassName("close")[0];
const popupDialogText = document.getElementById("popupDialog-text");
//// Important backend stuff
function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
};
localStorage.setItem("recentlogin", current_id);
//// Buttons
const btnClockIn = document.getElementById("btnClockIn");
btnClockIn.addEventListener("click", () => {

    if (isClockedIn) {
        popupDialog.style.display = "block";
        popupDialogText.innerHTML = "You are already clocked in.";
        return;
    }
    isClockedIn = true;
    clockInTime = new Date();

    const logData = {
        "_employee_id": empid,
        "_clock_in": clockInTime.toISOString(),
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
        popupDialog.style.display = "block";
        popupDialogText.innerHTML = "You need to clock in first.";
        return;
    }
    isBreakStarted = true;
    breakStartTime = new Date();

    let formData = new FormData()
    formData.append("data_value", btnbrkstrt.getAttribute("data-value"))

    const url = "/api/update_break";
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
        popupDialog.style.display = "block";
        popupDialogText.innerHTML = "You need to start a break first.";
        return;
    }
    isBreakStarted = false;
    breakEndTime = new Date();

    let formData = new FormData()
    formData.append("data_value", btnbrkend.getAttribute("data-value"))

    const url = "/api/update_break";
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
        popupDialogText.innerHTML = "You need to end the break first.";
        return;
    }
    isClockedIn = false;
    clockOutTime = new Date();
    
    let formData = new FormData()
    formData.append("data_value", btnClockOut.getAttribute("data-value"))

    const url = "/api/clockout";
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
            } catch (e) {
                console.error('Failed to parse JSON:', e);
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
// Filter function
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
// Popup Dialog
popupDialogclose.onclick = function() {
    popupDialog.style.display = "none";
}
window.onclick = function(event) {
    if (event.target == popupDialog) {
        popupDialog.style.display = "none";
    }
}
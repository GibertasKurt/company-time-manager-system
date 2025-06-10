var popupDialog = document.getElementById("popupDialog");
var popupDialogclose = document.getElementById("closePopup");
const popupDialogText = document.getElementById("popupDialog-text");
function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
}

document.getElementById("btnRegister").addEventListener("click", () => {
    const logData = {
        first_name: document.getElementById("first_name").value,
        last_name: document.getElementById("last_name").value,
        username: document.getElementById("username").value,
        password: document.getElementById("password").value,
        email: document.getElementById("email").value,
        department_id: document.getElementById("department_id").value,
    };
    for (const key in logData) {
        if (!logData[key]) {
            popupDialog.style.display = "block";
            popupDialogText.innerHTML = "Please fill in all fields.";
            return;
        }
    }

    console.log("Log data:", JSON.stringify(logData));

    fetch("/api/registration", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            // "X-CSRFToken": csrfToken
        },
        body: JSON.stringify(logData),
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`Server responded with status ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            console.log("Server response:", data);
        })
        .catch(error => {
            console.error("Fetch error:", error);
            alert("Failed to register: " + error.message + ". Please check your inputs.");
        });


});
popupDialogclose.onclick = function() { // Popup Dialog
    popupDialog.style.display = "none"
};
window.onclick = function(event) {
    if (event.target == popupDialog) { popupDialog.style.display = "none" }
};
function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
}

document.getElementById('currentYear').textContent = new Date().getFullYear();

document.getElementById("btnRegister").addEventListener("click", () => {
    const logData = {
        first_name: document.querySelector("#first_name").value,
        last_name: document.querySelector("#last_name").value,
        username: document.querySelector("#username").value,
        password: document.querySelector("#password").value,
        email: document.querySelector("#email").value,
        department_id: document.querySelector("#department_id").value,
    };

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
            // Handle success (e.g., redirect or show message)
        })
        .catch(error => {
            console.error("Fetch error:", error);
            alert("Failed to register: " + error.message + ". Please check your inputs.");
        });
});
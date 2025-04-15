const inputUsername = document.getElementById("username")
const inputPassword = document.getElementById("password")
let username;
let password;
const users = [
    {
        department: "Developer",
        username: "kurtgibertas",
        password: "admin123"
    },
    {
        department: "Developer",
        username: "josetapic",
        password: "admin123"
    },
    {
        department: "Developer",
        username: "arianecerezo",
        password: "admin123"
    }
]
const login = () => {
    username = inputUsername.value
    password = inputPassword.value

    const user = users.find(user => user.username === username && user.password === password)
    if (user) {
        alert("Login successful")
        localStorage.setItem("loggedInUser", JSON.stringify(user))
        window.location.href = "index.html"
    } else {
        alert("Invalid username or password")
    }
}
const btnLogin = document.getElementById("btnLogin").addEventListener("click", login)
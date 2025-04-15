const inputUsername = document.getElementById("username")
const inputPassword = document.getElementById("password")
const btnLogin = document.getElementById("btnLogin")
const users = [
    {
        username: "kurtgibertas",
        password: "admin123"
    },
    {
        username: "josetapic",
        password: "admin123"
    },
    {
        username: "arianecerezo",
        password: "admin123"
    }
]
const login = () => {
    const username = inputUsername.value
    const password = inputPassword.value

    const user = users.find(user => user.username === username && user.password === password)
    if (user) {
        alert("Login successful")
        localStorage.setItem("loggedInUser", JSON.stringify(user))
        window.location.href = "index.html"
    } else {
        alert("Invalid username or password")
    }
}
btnLogin.addEventListener("click", login)
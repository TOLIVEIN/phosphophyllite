"use strict";

let registerForm = document.getElementById("register-form");
let loginForm = document.getElementById("login-form");
let writeArticleForm = document.getElementById("write-article-form");

let username = document.getElementById("username");
let password = document.getElementById("password");
let repassword = document.getElementById("repassword");
let registerButton = document.getElementById("register");
let loginButton = document.getElementById("login");

function debounce(fn) {
    let timeout = null;
    return function () {
        clearTimeout(timeout);
        timeout = setTimeout(() => {
            fn.apply(this, arguments);
        }, 1000);
    };
}

function checkRePassword() {
    // let password = document.getElementById("password");
    // let repassword = document.getElementById("repassword");
    console.log(password.value, repassword.value);
    if (password.value != repassword.value) {
        // alert("密码不一致！");
        repassword.setCustomValidity("两次密码不一致！");
        registerButton.disabled = true;
    } else {
        repassword.setCustomValidity("");
        registerButton.disabled = false;
    }
}

function regiser() {
    let data = {
        username: username.value,
        password: password.value,
        repassword: repassword.value,
    };
    console.log(data);
    ajax({
        type: "post",
        url: "/register",
        data: data,
        success: (message) => {
            // console.log(response);
            window.location.href="/login"
            alert(message.message);
        },
        error: (error) => {
            alert(`error: ${error.message}`);
        },
    });
}

function login() {
    let data = {
        username: username.value,
        password: password.value,
    };
    console.log(data);
    ajax({
        type: "post",
        url: "/login",
        data: data,
        success: (message) => {
            window.location.href="/"
            alert(message.message);
            // console.log(response);
        },
        error: (error) => {
            alert(`error: ${error.message}`);
        },
    });
}

function writeArticle() {
    let title = document.getElementById("title").value;
    let tags = document.getElementById("tags").value;
    let brief = document.getElementById("brief").value;
    let content = document.getElementById("content").value;

    let data = {
        title: title,
        tags: tags,
        brief: brief,
        content: content,
    };
    console.log(data);
    ajax({
        type: "post",
        url: "/article/add",
        data: data,
        success: (message) => {
            window.location.href="/"
            alert(message.message);
            // console.log(response);
        },
        error: (error) => {
            alert(`error: ${error.message}`);
        },
    });
}

function ajax(options) {
    let defaults = {
        type: "get",
        url: "",
        data: {},
        header: {
            "Content-Type": "application/x-www-form-urlencoded",
        },
        success: function () {},
        error: function () {},
    };

    Object.assign(defaults, options);
    let xhr = new XMLHttpRequest();
    let params = "";
    for (let key in defaults.data) {
        params += key + "=" + defaults.data[key] + "&";
    }
    params = params.substr(0, params.length - 1);

    if (defaults.type == "get") {
        defaults.url = defaults.url + "?" + params;
    }

    xhr.open(defaults.type, defaults.url);

    if (defaults.type == "post") {
        let contentType = defaults.header["Content-Type"];
        xhr.setRequestHeader("Content-Type", contentType);
        if (contentType == "application/json") {
            xhr.send(JSON.stringify(defaults.data));
        } else {
            xhr.send(params);
        }
    } else {
        xhr.send();
    }

    xhr.onload = () => {
        // alert('onload.......')
        let contentType = xhr.getResponseHeader("Content-Type");
        let responseText = xhr.responseText;
        if (contentType.includes("application/json")) {
            responseText = JSON.parse(responseText);
        }
        if (xhr.status == 200) {
            defaults.success(responseText, xhr);
        } else {
            defaults.error(responseText, xhr);
        }
    };
}

if (repassword) {
    repassword.addEventListener("keyup", () => {
        checkRePassword();
    });
}

if (registerForm) {
    registerForm.addEventListener("submit", () => {
        regiser();
    });
}

if (loginForm) {
    loginForm.addEventListener("submit", () => {
        login();
    });
}


if (writeArticleForm) {
    writeArticleForm.addEventListener("submit", () => {
        writeArticle();
    })
}
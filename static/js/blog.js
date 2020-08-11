"use strict";

let form = document.getElementsByTagName("form")[0];

let username = document.getElementById("username");
let password = document.getElementById("password");
let repassword = document.getElementById("repassword");
let submit = document.getElementById("login-button");

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
    console.log(password.value, repassword.value);
    if (password.value != repassword.value) {
        // alert("密码不一致！");
        repassword.setCustomValidity("两次密码不一致！");
    }
}

function regiser() {
    // let username = document.getElementById("username").value;
    // let password = document.getElementById("password").value;
    // let repassword = document.getElementById("repassword").value;
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
            // window.location.href="/login"
            alert(message.message);
            // console.log(response);
        },
        error: (error) => {
            alert(`error: ${error.message}`)
        }
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

repassword.addEventListener("keyup", debounce(checkRePassword));

form.addEventListener("submit", (event) => {
    regiser();
});

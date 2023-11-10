// /*
// 		Designed by: SELECTO
// 		Original image: https://dribbble.com/shots/5311359-Diprella-Login
// */

let switchCtn = document.querySelector("#switch-cnt");
let switchC1 = document.querySelector("#switch-c1");
let switchC2 = document.querySelector("#switch-c2");
let switchCircle = document.querySelectorAll(".switch__circle");
let switchBtn = document.querySelectorAll(".switch-btn");
let aContainer = document.querySelector("#a-container");
let bContainer = document.querySelector("#b-container");

let getButtons = (e) => e.preventDefault()

let changeForm = (e) => {

    switchCtn.classList.add("is-gx");
    setTimeout(function(){
        switchCtn.classList.remove("is-gx");
    }, 1500)

    switchCtn.classList.toggle("is-txr");
    switchCircle[0].classList.toggle("is-txr");
    switchCircle[1].classList.toggle("is-txr");

    switchC1.classList.toggle("is-hidden");
    switchC2.classList.toggle("is-hidden");
    aContainer.classList.toggle("is-txl");
    bContainer.classList.toggle("is-txl");
    bContainer.classList.toggle("is-z200");
}

let mainF = (e) => {
    for (var i = 0; i < switchBtn.length; i++)
        switchBtn[i].addEventListener("click", changeForm)
}

window.addEventListener("load", mainF);



// // parse form
// document.getElementById('updatePassword').addEventListener('change', function() {
//     var passwordDiv = document.getElementById('passwordDiv');
//     if (this.checked) {
//         passwordDiv.style.display = 'block';
//     } else {
//         passwordDiv.style.display = 'none';
//     }
// });
// document.getElementById('imgInput').addEventListener('change', function() {
//     var file = this.files[0];
//     if (file.size > 800000) {
//         alert("The image size should not exceed 800KB.");
//         this.value = "";
//     } else {
//         var reader = new FileReader();
//         reader.onload = function(e) {
//             document.getElementById('preview').src = e.target.result;
//         }
//         reader.readAsDataURL(file);
//     }
// });
document.getElementsByClassName('myForm').addEventListener('submit', function(e) {
    e.preventDefault();
    
    var password = document.getElementsByClassName('password').value;
    var email = document.getElementsByClassName('email').value;
    
    if (password.length < 4 || email.trim() === "") {
        alert("Username and Email should not be empty.");
        return false;
    }
    
    if (document.getElementsByClassName('updatePassword').checked) {
        var password = document.getElementsByClassName('password').value;
        if (password.length < 4) {
            alert("Password should be at least 4 characters and should not contain spaces.");
            return false;
        }
    }
    
    var formData = new FormData(this);
    
    fetch('/login', {
        method: 'POST',
        body: formData
    })
    .then(response => response.json())
    .then(result => {
        console.log('Success:', result);
    })
    .catch(error => {
        console.error('Error:', error);
    });
});
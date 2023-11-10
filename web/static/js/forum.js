// profil dropdown
function myFunction(postID) {
    document.getElementById(postID).classList.toggle("show");
}
// Function to set the active link based on URL parameter
// function setActiveLink() {
//     const urlParams = new URLSearchParams(window.location.search);
//     const nameParam = urlParams.get('name');
//     const links = document.querySelectorAll('.profile-menu-link');
//     // Remove the "active" class from all links
//     links.forEach(link => link.classList.remove('active'));
//     // Add the "active" class to the link with the corresponding name parameter
//     links.forEach(link => {
//         if (link.getAttribute('href').includes(`name=${nameParam}`)) {
//             link.classList.add('active');
//         }
//     });
// }
// // Check if no link has the "active" class, then add it to the "All Posts" link
// const activeLinks = document.querySelectorAll('.profile-menu-link.active');
// if (activeLinks.length === 0) {
//     document.querySelector('.profile-menu-link[href^="/forum"]').classList.add('active');
// }
// // Call the setActiveLink function to set the initial active link based on URL parameter
// setActiveLink();


// Upload image
// var loadFile = function(event) {
//     var output = document.getElementById('output');
//     output.src = URL.createObjectURL(event.target.files[0]);
//     output.onload = function() {
//       URL.revokeObjectURL(output.src) // libère l'objet URL
//     }
// };

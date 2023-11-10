// Fonction pour envoyer la rétroaction (feedback) au serveur
function sendFeedback(postId, action, url) {
  const likeCountPost = document.querySelector(`.likeCount[data-post-id="${postId}"]`);
  const dislikeCountPost = document.querySelector(`.dislikeCount[data-post-id="${postId}"]`);

  // Exemple de requête Fetch API
  fetch("/like?name="+url, {
    method: 'POST',
    body: JSON.stringify({ postId: postId, action: action }),
    headers: {
        'Content-Type': 'application/json'
    }
  })
  .then(response => response.json())
  .then(data => {
    // Mettez à jour les compteurs avec les données de la réponse du serveur
    likeCountPost.textContent = data.likeCount;
    dislikeCountPost.textContent = data.dislikeCount;
  })
  .catch(error => {
    console.error(error);
  });
}


// function sendFeed(postId, action) {
//   console.log(postId, action);
//   const likeCountPost = document.querySelector(`.likeCount[data-post-id="${postId}"]`);
//   const dislikeCountPost = document.querySelector(`.dislikeCount[data-post-id="${postId}"]`);

//   // Exemple de requête Fetch API
//   fetch(`/like?name=comment_id`, {
//     method: 'POST',
//     body: JSON.stringify({ postId: postId, action: action }),
//     headers: {
//         'Content-Type': 'application/json'
//     }
//   })
//   .then(response => response.json())
//   .then(data => {
//     // Mettez à jour les compteurs avec les données de la réponse du serveur
//     likeCountPost.textContent = data.likeCount;
//     dislikeCountPost.textContent = data.dislikeCount;
//   })
//   .catch(error => {
//     console.error(error);
//   });
// }

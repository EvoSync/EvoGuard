// signout.js introduces the signout functionality within the
// login endpoint, this means once we recv the signout call
// within the html file/signout function is called it will call
// the api endpoint, whent he api endpoint is called it returns
// a status, if said status is equal to true, we destroy the
// cookie contained within the brower and redirect them directly
// to the */* portal, otherwise know as the login page.

function signout() {
      // destorys the current cookie which the browsers hold/stores
      document.cookie = ""
  
      // performs the signout request to the api endpoint
      // the api endpoint will absorb said request and
      // attempts to try retrieve it within the tokens map
      // and destroy it meaning any futher subsequent requests
      // have to relogin and make a new cookie & token
      fetch(currentTokenInsideURL() + "/api/signout").then(res => res.json()).then(res => {
          if (res["status"] == true) {
              location.href = "/"
          } else {
              console.log("Error occurred while performing the signout api request")
          }
      })
  }
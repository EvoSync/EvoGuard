/*
      cookie.js is the function which will inspect and show us
      information about the cookie it's self and it's owner
      like there username and information about there
      request.

      /mainframe.js
*/

// cookie returns my current account information
async function cookie() {
      const object = currentTokenInsideURL()

      /* saveReference is where we shall absorb the output into and handle accordingly */
      var saveReference = []

      // makes the fetch request towards the api endpoint
      await fetch(object + "/api/me").then(res => res.json()).then(res => {
            saveReference = res
      })

      return saveReference
}
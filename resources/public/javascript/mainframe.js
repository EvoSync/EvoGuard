// login.js will implement certain features/methods into
// the login endpoint. this means we can redirect to
// certain endpoints but keep all the querys parameters.

// redirect will implement the redirection of requests but keeping the query params
function redirect(newPath) {
        if (document.location.search.length == 0) {
                document.location.href = newPath
                return
        } 
            
        // If we know there is actual query params, we redirect with them included
        document.location.href = newPath + document.location.search
}

// redirectFromTokenOnwards will implement the required functionalit
function redirectFromTokenOnwards(newPath) {
      document.location.href = currentTokenInsideURL() + newPath
}

// token will return the required information and parameters
function currentTokenInsideURL() {
      return window.location.href.split('/').slice(0, 4).join("/")
}

function encodeData(data) {
      return Object.keys(data).map(function(key) {
          return [key, data[key]].map(encodeURIComponent).join("=");
      }).join("&");
}   
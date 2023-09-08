/*
      clients.js is the sdk which helps implementing fetching
      all of a clients children, a clients child is someone
      who was created under them.
*/

// clients will get a list of all the clients inside the database
async function clients() {
      const object = currentTokenInsideURL()

      /* saveReference is where we shall absorb the output into and handle accordingly */
      var saveReference = []

      // makes the fetch request towards the api endpoint
      await fetch(object + "/api/users/getusers").then(res => res.json()).then(res => {
            saveReference = res
      })

      return saveReference
}

// createClient will implement the api callback for making a client inside the database
async function createClient(username, password, email, account) {
      const object = currentTokenInsideURL() 

      // saveReference is where we shall absorb the output into and handle accordingly */
      var saveReference = []


      fetch(object + "/api/users/createuser?" + encodeData({"username": username, "password": password, "email": email, "accountlevel": account})).then(res => res.json()).then(res => {
            console.log(res)
      })

}
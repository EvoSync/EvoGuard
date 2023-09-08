/*
      clients.js is implemented for the /html/admin/clients.html
      hypertext file, which we will execute on startup and disable
      said loader
*/

// cookie will return information which is tailored towards the
// client who is currently logged in, some of this will contain
// there username to a hashed vairent of there password
cookie().then(res => {
      document.getElementById("Username").textContent = res["username"]

      document.getElementById("spinner").style.visibility = "hidden"
      document.getElementById("app").style.visibility = "visible"

      // clients will allow me to access all the users
      // found under a certain account/cookie
      clients().then(res => {
            const body = document.getElementById("component-users")

            // iterates over each account found within the array
            for (const element in res) {
                  insertRow(body.insertRow(), res[element]["username"], res[element]["email"], res[element]["account_level"])
            }
      })
})

function insertRow(row, username, email, accountlevel) {
      row.insertCell(0).innerHTML = "<a class=\"user\" href=\""+currentTokenInsideURL() + "/dashboard/clients?user=" + username +"\">"+ username +"</a>"
      row.insertCell(1).innerHTML = email
      switch (accountlevel) {

      case 0: // Admin
            row.insertCell(2).innerHTML = "Administrator"
            return

      case 1: // Client
            row.insertCell(2).innerHTML = "Client"
            return

      default: // Unknown
            row.insertCell(2).innerHTML = "Unconfirmed"
            return
      }
}

// createUser is the function event which triggers
// when the createUser button is fired.
function fireUserMenu() {
      document.getElementById("dashboard").style.opacity = .2;
      document.getElementById("modal").style.visibility = "visible";
}


// createUser is the event which is triggered when
// someone clicks the *add user* button to create
// said user via the api portal.
function createUser() {
      let username = document.getElementById("username").value;
      let password = document.getElementById("password").value;
      let email = document.getElementById("email").value;

      // offical is the account level we provide for the user
      let offical = 1

      // If they want to create an admin user they can do here
      if (document.getElementById("roles").value == "admin") {
            offical = 0
      }

      let res = createClient(username, password, email, offical);
      if (res["status"] == false) {
            // implement a toast for when the account already exists?
            return
      }

      insertRow(document.getElementById("component-users").insertRow(), username, email, offical)

      // Once the user has been created we will continue to hide the modal
      hideModal()
}